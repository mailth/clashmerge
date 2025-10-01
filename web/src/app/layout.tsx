import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: 'ClashMerge 管理页面',
  description: 'ClashMerge Admin Frontend',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="zh-CN">
      <body>{children}</body>
    </html>
  )
}