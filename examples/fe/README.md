# Todo List Frontend

基于Vue3的Todo List前端项目，对接api-server后端API。

## 功能特性

- ✅ 获取Todo列表
- ✅ 创建新的Todo
- ✅ 切换Todo完成状态
- ✅ 删除Todo
- ✅ 响应式设计
- ✅ 简洁清晰的UI

## 技术栈

- **前端框架**: Vue 3
- **构建工具**: Vite
- **HTTP客户端**: Axios
- **样式**: CSS3

## 快速开始

### 1. 环境要求

- Node.js 16+
- npm 或 yarn

### 2. 安装依赖

```bash
npm install
```

### 3. 启动开发服务器

```bash
npm run dev
```

前端将在 `http://localhost:3000` 启动。

### 4. 确保后端服务运行

确保api-server后端服务在 `http://localhost:9000` 运行：

```bash
# 在api-server目录下
docker-compose up -d
go run main.go
```

## 项目结构

```
fe/
├── src/
│   ├── components/
│   │   └── TodoList.vue      # Todo列表组件
│   ├── services/
│   │   └── api.js            # API服务
│   ├── App.vue               # 主应用组件
│   └── main.js               # 应用入口
├── index.html                # HTML模板
├── vite.config.js            # Vite配置
├── package.json              # 项目配置
└── README.md                 # 项目说明
```

## API对接

项目通过以下API与后端交互：

- `GET /api/todos` - 获取Todo列表
- `POST /api/todos` - 创建新的Todo
- `PUT /api/todos/:id` - 更新Todo状态
- `DELETE /api/todos/:id` - 删除Todo

## 开发说明

- 使用Vite作为构建工具，支持热重载
- 通过代理配置解决跨域问题
- 响应式设计，适配不同屏幕尺寸
- 简洁的UI设计，良好的用户体验

## 构建部署

```bash
# 构建生产版本
npm run build

# 预览构建结果
npm run preview
```

构建后的文件在 `dist` 目录中。
