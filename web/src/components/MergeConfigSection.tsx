"use client";

import MergeConfigModal from "@/components/MergeConfigModal";
import { apiService } from "@/services/api";
import { MergeConfig } from "@/types";
import { useEffect, useState } from "react";

interface MergeConfigSectionProps {
  isActive: boolean;
}

export default function MergeConfigSection({
  isActive,
}: MergeConfigSectionProps) {
  const [mergeConfigs, setMergeConfigs] = useState<MergeConfig[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [showModal, setShowModal] = useState(false);
  const [editingConfig, setEditingConfig] = useState<MergeConfig | null>(null);

  useEffect(() => {
    if (isActive) {
      loadData();
    }
  }, [isActive]);

  const loadData = async () => {
    try {
      setLoading(true);
      setError(null);

      const configs = await apiService.getMergeConfigs();
      setMergeConfigs(configs);
    } catch (err) {
      setError(err instanceof Error ? err.message : "加载数据失败");
    } finally {
      setLoading(false);
    }
  };

  const handleAdd = () => {
    setEditingConfig(null);
    setShowModal(true);
  };

  const handleEdit = (config: MergeConfig) => {
    setEditingConfig(config);
    setShowModal(true);
  };

  const handleDelete = async (id: number) => {
    if (!confirm("确定要删除这个 Merge 配置吗？")) return;

    try {
      await apiService.deleteMergeConfig(id);
      await loadData();
      alert("删除成功");
    } catch (err) {
      alert(err instanceof Error ? err.message : "删除失败");
    }
  };

  const handleModalClose = () => {
    setShowModal(false);
    setEditingConfig(null);
  };

  const handleModalSubmit = async () => {
    await loadData();
    handleModalClose();
  };

  if (!isActive) {
    return <div id='merge-config' className='section hidden'></div>;
  }

  return (
    <div id='merge-config' className='section'>
      <div className='flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6'>
        <div>
          <h2 className='text-xl font-semibold text-gray-800'>
            Merge 配置管理
          </h2>
          <p className='text-gray-600 text-sm mt-1'>管理您的 Merge 规则配置</p>
        </div>
        <button
          className='mt-4 sm:mt-0 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors duration-200'
          onClick={handleAdd}
        >
          <svg
            xmlns='http://www.w3.org/2000/svg'
            className='h-5 w-5 mr-2'
            viewBox='0 0 20 20'
            fill='currentColor'
          >
            <path
              fillRule='evenodd'
              d='M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z'
              clipRule='evenodd'
            />
          </svg>
          添加 Merge 配置
        </button>
      </div>

      {loading && (
        <div className='flex justify-center items-center py-12'>
          <div className='animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500'></div>
          <span className='ml-3 text-gray-600'>加载中...</span>
        </div>
      )}

      {error && (
        <div className='bg-red-50 border-l-4 border-red-500 p-4 mb-6 rounded'>
          <div className='flex'>
            <div className='flex-shrink-0'>
              <svg
                className='h-5 w-5 text-red-400'
                xmlns='http://www.w3.org/2000/svg'
                viewBox='0 0 20 20'
                fill='currentColor'
              >
                <path
                  fillRule='evenodd'
                  d='M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z'
                  clipRule='evenodd'
                />
              </svg>
            </div>
            <div className='ml-3'>
              <p className='text-sm text-red-700'>{error}</p>
            </div>
          </div>
        </div>
      )}

      {!loading && !error && mergeConfigs.length === 0 && (
        <div className='text-center py-12 bg-gray-50 rounded-lg border-2 border-dashed border-gray-300'>
          <svg
            xmlns='http://www.w3.org/2000/svg'
            className='mx-auto h-12 w-12 text-gray-400'
            fill='none'
            viewBox='0 0 24 24'
            stroke='currentColor'
          >
            <path
              strokeLinecap='round'
              strokeLinejoin='round'
              strokeWidth={2}
              d='M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z'
            />
          </svg>
          <h3 className='mt-2 text-sm font-medium text-gray-900'>
            暂无 Merge 配置
          </h3>
          <p className='mt-1 text-sm text-gray-500'>
            点击上方按钮添加您的第一个 Merge 配置
          </p>
        </div>
      )}

      {!loading && !error && mergeConfigs.length > 0 && (
        <div className='overflow-hidden shadow ring-1 ring-black ring-opacity-5 rounded-lg'>
          <table className='min-w-full divide-y divide-gray-300'>
            <thead className='bg-gray-50'>
              <tr>
                <th
                  scope='col'
                  className='py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6'
                >
                  名称
                </th>
                <th
                  scope='col'
                  className='px-3 py-3.5 text-left text-sm font-semibold text-gray-900'
                >
                  描述
                </th>
                <th scope='col' className='relative py-3.5 pl-3 pr-4 sm:pr-6'>
                  <span className='sr-only'>操作</span>
                </th>
              </tr>
            </thead>
            <tbody className='divide-y divide-gray-200 bg-white'>
              {mergeConfigs.map((config) => (
                <tr
                  key={config.id}
                  className='hover:bg-gray-50 transition-colors duration-150'
                >
                  <td className='whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-6'>
                    {config.name}
                  </td>
                  <td className='px-3 py-4 text-sm text-gray-500'>
                    {config.description || "-"}
                  </td>
                  <td className='whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-6'>
                    <button
                      className='text-blue-600 hover:text-blue-900 mr-3 inline-flex items-center'
                      onClick={() => handleEdit(config)}
                    >
                      <svg
                        xmlns='http://www.w3.org/2000/svg'
                        className='h-4 w-4 mr-1'
                        viewBox='0 0 20 20'
                        fill='currentColor'
                      >
                        <path d='M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z' />
                      </svg>
                      编辑
                    </button>
                    <button
                      className='text-red-600 hover:text-red-900 inline-flex items-center'
                      onClick={() => handleDelete(config.id!)}
                    >
                      <svg
                        xmlns='http://www.w3.org/2000/svg'
                        className='h-4 w-4 mr-1'
                        viewBox='0 0 20 20'
                        fill='currentColor'
                      >
                        <path
                          fillRule='evenodd'
                          d='M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z'
                          clipRule='evenodd'
                        />
                      </svg>
                      删除
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      {showModal && (
        <MergeConfigModal
          config={editingConfig}
          onClose={handleModalClose}
          onSubmit={handleModalSubmit}
        />
      )}
    </div>
  );
}
