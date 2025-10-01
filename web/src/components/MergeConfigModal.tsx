"use client";

import { apiService } from "@/services/api";
import { MergeConfig } from "@/types";
import { useEffect, useState } from "react";

interface MergeConfigModalProps {
  config: MergeConfig | null;
  onClose: () => void;
  onSubmit: () => void;
}

export default function MergeConfigModal({
  config,
  onClose,
  onSubmit,
}: MergeConfigModalProps) {
  const [formData, setFormData] = useState({
    id: 0,
    name: "",
    rules: "",
    proxies: "",
    proxy_groups: "",
    description: "",
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (config) {
      setFormData({
        id: config.id!,
        name: config.name,
        rules: config.rules || "",
        proxies: config.proxies || "",
        proxy_groups: config.proxy_groups || "",
        description: config.description || "",
      });
    } else {
      setFormData({
        id: 0,
        name: "",
        rules: "",
        proxies: "",
        proxy_groups: "",
        description: "",
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
        id: formData.id,
        name: formData.name,
        rules: formData.rules,
        proxies: formData.proxies,
        proxy_groups: formData.proxy_groups,
        description: formData.description,
      };

      if (config) {
        await apiService.updateMergeConfig(config.id!, data);
      } else {
        await apiService.createMergeConfig(data);
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
        <h3 id='merge-modal-title'>
          {config ? "编辑 Merge 配置" : "添加 Merge 配置"}
        </h3>

        <form id='merge-form' onSubmit={handleSubmit}>
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
            <label>描述:</label>
            <textarea
              name='description'
              value={formData.description}
              onChange={handleChange}
            />
          </div>

          <div className='form-group' id='rules-group'>
            <label>前置规则 (YAML 格式):</label>
            <textarea
              name='rules'
              value={formData.rules}
              onChange={handleChange}
              rows={5}
              placeholder='- DOMAIN-SUFFIX,example.com,DIRECT'
            />
          </div>

          <div className='form-group' id='proxies-group'>
            <label>前置代理 (YAML 格式):</label>
            <textarea
              name='proxies'
              value={formData.proxies}
              onChange={handleChange}
              rows={5}
              placeholder='- name: proxy1\n  type: ss\n  server: 1.2.3.4'
            />
          </div>

          <div className='form-group' id='proxies-group'>
            <label>前置代理组 (YAML 格式):</label>
            <textarea
              name='proxy_groups'
              value={formData.proxy_groups}
              onChange={handleChange}
              rows={5}
              placeholder='- name: proxy1\n  type: ss\n  server: 1.2.3.4'
            />
          </div>

          <button type='submit' className='btn' disabled={loading}>
            {loading ? "保存中..." : "保存"}
          </button>
          {error && <div className='error'>{error}</div>}
        </form>
      </div>
    </div>
  );
}
