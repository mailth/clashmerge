"use client";

import { apiService } from "@/services/api";
import { LinkConfig, MergeConfig } from "@/types";
import { useEffect, useState } from "react";

interface LinkConfigModalProps {
  config: LinkConfig | null;
  mergeConfigs: MergeConfig[];
  onClose: () => void;
  onSubmit: () => void;
}

export default function LinkConfigModal({
  config,
  mergeConfigs,
  onClose,
  onSubmit,
}: LinkConfigModalProps) {
  const [formData, setFormData] = useState({
    name: "",
    clash_url: "",
    description: "",
    merge_config_id: "",
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (config) {
      setFormData({
        name: config.name,
        clash_url: config.clash_url,
        description: config.description || "",
        merge_config_id: config.merge_config_id?.toString() || "",
      });
    } else {
      setFormData({
        name: "",
        clash_url: "",
        description: "",
        merge_config_id: "",
      });
    }
  }, [config]);

  const handleChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement
    >
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    try {
      const data = {
        name: formData.name,
        clash_url: formData.clash_url,
        description: formData.description,
        merge_config_id: formData.merge_config_id
          ? parseInt(formData.merge_config_id)
          : undefined,
      };

      if (config) {
        await apiService.updateLinkConfig(config.id!, data);
      } else {
        await apiService.createLinkConfig(data);
      }

      onSubmit();
    } catch (err) {
      setError(err instanceof Error ? err.message : "保存失败");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className='modal' onClick={onClose}>
      <div className='modal-content' onClick={(e) => e.stopPropagation()}>
        <span className='close' onClick={onClose}>
          &times;
        </span>
        <h3 id='link-modal-title'>
          {config ? "编辑链接配置" : "添加链接配置"}
        </h3>

        {error && <div className='error'>{error}</div>}

        <form id='link-form' onSubmit={handleSubmit}>
          <div className='form-group'>
            <label>名称:</label>
            <input
              type='text'
              name='name'
              value={formData.name}
              onChange={handleChange}
              required
            />
          </div>

          <div className='form-group'>
            <label>Clash URL:</label>
            <input
              type='url'
              name='clash_url'
              value={formData.clash_url}
              onChange={handleChange}
              required
            />
          </div>

          <div className='form-group'>
            <label>描述:</label>
            <textarea
              name='description'
              value={formData.description}
              onChange={handleChange}
            />
          </div>

          <div className='form-group'>
            <label>关联 Merge 配置:</label>
            <select
              name='merge_config_id'
              value={formData.merge_config_id}
              onChange={handleChange}
            >
              <option value=''>请选择</option>
              {mergeConfigs.map((mergeConfig) => (
                <option key={mergeConfig.id} value={mergeConfig.id}>
                  {mergeConfig.name}
                </option>
              ))}
            </select>
          </div>

          <button type='submit' className='btn' disabled={loading}>
            {loading ? "保存中..." : "保存"}
          </button>
        </form>
      </div>
    </div>
  );
}
