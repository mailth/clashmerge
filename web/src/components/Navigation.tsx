'use client'

import { SectionType } from '@/types'

interface NavigationProps {
  activeSection: SectionType
  onSectionChange: (section: SectionType) => void
}

export default function Navigation({ activeSection, onSectionChange }: NavigationProps) {
  return (
    <div className="nav flex space-x-1 bg-gray-100 p-1 rounded-lg">
      <button
        className={`nav-item px-4 py-2 rounded-md text-sm font-medium transition-all duration-200 flex items-center ${
          activeSection === 'link-config'
            ? 'bg-white text-blue-600 shadow-sm'
            : 'text-gray-600 hover:text-gray-900 hover:bg-gray-50'
        }`}
        onClick={() => onSectionChange('link-config')}
      >
        <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-2" viewBox="0 0 20 20" fill="currentColor">
          <path fillRule="evenodd" d="M12.586 4.586a2 2 0 112.828 2.828l-3 3a2 2 0 01-2.828 0 1 1 0 00-1.414 1.414 4 4 0 005.656 0l3-3a4 4 0 00-5.656-5.656l-1.5 1.5a1 1 0 101.414 1.414l1.5-1.5zm-5 5a2 2 0 012.828 0 1 1 0 101.414-1.414 4 4 0 00-5.656 0l-3 3a4 4 0 105.656 5.656l1.5-1.5a1 1 0 10-1.414-1.414l-1.5 1.5a2 2 0 11-2.828-2.828l3-3z" clipRule="evenodd" />
        </svg>
        链接配置
      </button>
      <button
        className={`nav-item px-4 py-2 rounded-md text-sm font-medium transition-all duration-200 flex items-center ${
          activeSection === 'merge-config'
            ? 'bg-white text-blue-600 shadow-sm'
            : 'text-gray-600 hover:text-gray-900 hover:bg-gray-50'
        }`}
        onClick={() => onSectionChange('merge-config')}
      >
        <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4 mr-2" viewBox="0 0 20 20" fill="currentColor">
          <path fillRule="evenodd" d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm3.293-7.707a1 1 0 011.414 0L9 10.586V3a1 1 0 112 0v7.586l1.293-1.293a1 1 0 111.414 1.414l-3 3a1 1 0 01-1.414 0l-3-3a1 1 0 010-1.414z" clipRule="evenodd" />
        </svg>
        Merge 配置
      </button>
    </div>
  )
}