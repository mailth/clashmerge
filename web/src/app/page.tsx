'use client'

import { useState } from 'react'
import Navigation from '@/components/Navigation'
import LinkConfigSection from '@/components/LinkConfigSection'
import MergeConfigSection from '@/components/MergeConfigSection'

export default function Home() {
  const [activeSection, setActiveSection] = useState<'link-config' | 'merge-config'>('link-config')

  return (
    <div id="app" className="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100">
      <div className="container mx-auto px-4 py-8 max-w-7xl">
        <div className="header bg-white rounded-xl shadow-lg p-6 mb-8 border border-gray-200">
          <div className="flex flex-col md:flex-row md:items-center md:justify-between">
            <div className="mb-4 md:mb-0">
              <h1 className="text-3xl font-bold text-gray-800 flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8 mr-3 text-blue-600" viewBox="0 0 20 20" fill="currentColor">
                  <path fillRule="evenodd" d="M12.316 3.051a1 1 0 01.633 1.265l-4 12a1 1 0 11-1.898-.632l4-12a1 1 0 011.265-.633zM5.707 6.293a1 1 0 010 1.414L3.414 10l2.293 2.293a1 1 0 11-1.414 1.414l-3-3a1 1 0 010-1.414l3-3a1 1 0 011.414 0zm8.586 0a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 11-1.414-1.414L16.586 10l-2.293-2.293a1 1 0 010-1.414z" clipRule="evenodd" />
                </svg>
                ClashMerge 管理页面
              </h1>
              <p className="text-gray-600 mt-2">管理您的 Clash 配置和 Merge 规则</p>
            </div>
            <Navigation activeSection={activeSection} onSectionChange={setActiveSection} />
          </div>
        </div>

        <div className="content bg-white rounded-xl shadow-lg p-6 border border-gray-200">
          <div className="mb-6">
            <div className="flex items-center justify-between">
              <h2 className="text-xl font-semibold text-gray-800">
                {activeSection === 'link-config' ? '链接配置管理' : 'Merge 配置管理'}
              </h2>
              <div className="text-sm text-gray-500 bg-gray-100 px-3 py-1 rounded-full">
                {activeSection === 'link-config' ? '管理 Clash 链接配置' : '管理 Merge 规则配置'}
              </div>
            </div>
            <div className="h-1 w-20 bg-blue-500 rounded mt-2"></div>
          </div>
          
          <LinkConfigSection isActive={activeSection === 'link-config'} />
          <MergeConfigSection isActive={activeSection === 'merge-config'} />
        </div>
        
        <footer className="mt-8 text-center text-gray-500 text-sm">
          <p>© 2023 ClashMerge Admin. All rights reserved.</p>
        </footer>
      </div>
    </div>
  )
}