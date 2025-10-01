# ClashMerge Admin Frontend (Next.js)

这是 ClashMerge 管理页面的 Next.js 实现，替代了原有的 webpack + TypeScript 实现。

## 功能特性

- 链接配置管理（增删改查）
- Merge 配置管理（增删改查）
- 响应式设计
- 模态框表单交互
- TypeScript 支持

## 技术栈

- Next.js 14
- React 18
- TypeScript
- Tailwind CSS
- 原生 CSS（迁移自原有样式）

## 项目结构

```
nextjs/
├── src/
│   ├── app/                 # Next.js App Router
│   │   ├── layout.tsx       # 根布局
│   │   ├── page.tsx         # 主页面
│   │   └── globals.css      # 全局样式
│   ├── components/          # React 组件
│   │   ├── Navigation.tsx   # 导航组件
│   │   ├── LinkConfigSection.tsx    # 链接配置区域
│   │   ├── MergeConfigSection.tsx    # Merge 配置区域
│   │   ├── LinkConfigModal.tsx      # 链接配置模态框
│   │   └── MergeConfigModal.tsx      # Merge 配置模态框
│   ├── services/           # API 服务
│   │   └── api.ts         # API 调用封装
│   └── types/              # TypeScript 类型定义
│       └── index.ts        # 类型定义
├── package.json
├── next.config.js
├── tsconfig.json
├── tailwind.config.js
└── postcss.config.js
```

## 安装和运行

1. 安装依赖：
```bash
cd nextjs
npm install
```

2. 启动开发服务器：
```bash
npm run dev
```

3. 构建生产版本：
```bash
npm run build
```

4. 启动生产服务器：
```bash
npm start
```

## 与原项目的对比

### 原项目 (web/)
- 使用 webpack 作为构建工具
- 原生 JavaScript 操作 DOM
- 手动管理组件状态和生命周期
- 单一的 app.ts 文件包含所有逻辑

### 新项目 (nextjs/)
- 使用 Next.js 框架
- React 组件化开发
- 使用 React Hooks 管理状态
- 组件拆分更清晰，职责单一
- 更好的 TypeScript 支持
- 更现代的开发体验

## API 接口

项目使用与原项目相同的 API 接口，基础 URL 为 `/admin`：

### 链接配置 API
- `GET /admin/link-configs` - 获取所有链接配置
- `POST /admin/link-configs` - 创建链接配置
- `PUT /admin/link-configs/:id` - 更新链接配置
- `DELETE /admin/link-configs/:id` - 删除链接配置

### Merge 配置 API
- `GET /admin/merge-configs` - 获取所有 Merge 配置
- `POST /admin/merge-configs` - 创建 Merge 配置
- `PUT /admin/merge-configs/:id` - 更新 Merge 配置
- `DELETE /admin/merge-configs/:id` - 删除 Merge 配置

## 注意事项

1. 确保后端 API 服务正常运行
2. 项目使用了 Tailwind CSS，但也保留了原有的 CSS 样式
3. 所有组件都使用了 TypeScript，提供了类型安全
4. 模态框组件支持点击背景关闭