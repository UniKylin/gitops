# Xvue - Vue 3 生产级应用

一个基于 Vue 3 + Vite + Docker 的现代化前端应用，专为生产环境设计。

## ✨ 特性

- 🚀 **Vue 3** 使用 Composition API
- ⚡ **Vite** 现代化构建工具
- 🎨 **响应式 UI** 美观的用户界面
- 🔄 **交互功能** 计数器和消息输入
- 🔥 **热重载** 开发时实时更新
- 🐳 **Docker 容器化** 生产就绪
- 📦 **极简镜像** 仅 14.4MB
- 🛡️ **安全配置** 生产级安全设置

## 🚀 快速开始

### 环境要求

- Node.js 16+ 
- npm 或 yarn
- Docker (可选)

### 本地开发

1. **克隆项目**
   ```bash
   git clone <repository-url>
   cd xvue
   ```

2. **安装依赖**
   ```bash
   npm install
   ```

3. **启动开发服务器**
   ```bash
   npm run dev
   ```

4. **访问应用**
   ```
   http://localhost:3000
   ```

### 可用脚本

| 命令 | 说明 |
|------|------|
| `npm run dev` | 启动开发服务器 |
| `npm run build` | 构建生产版本 |
| `npm run preview` | 预览构建结果 |
| `npm run serve` | 使用 http-server 服务构建文件 |

## 🐳 Docker 部署

### 构建镜像

```bash
# 构建 Docker 镜像
docker build -t xvue .

# 查看镜像大小
docker images xvue
```

### 运行容器

```bash
# 运行容器
docker run -p 9000:9000 xvue

# 后台运行
docker run -d -p 9000:9000 --name xvue-app xvue

# 查看容器状态
docker ps

# 查看健康状态
docker inspect --format='{{.State.Health.Status}}' xvue-app
```

### 访问应用

```
http://localhost:9000
```

### 健康检查

```bash
# 检查应用健康状态
curl http://localhost:9000/health
```

## 📁 项目结构

```
xvue/
├── src/                    # 源代码目录
│   ├── App.vue            # 主 Vue 组件
│   └── main.js            # 应用入口文件
├── index.html             # HTML 模板
├── package.json           # 项目配置和依赖
├── vite.config.js         # Vite 构建配置
├── Dockerfile             # 多阶段 Docker 构建文件
├── nginx.conf             # Nginx 生产环境配置
├── .dockerignore          # Docker 构建忽略文件
├── .env.example           # 环境变量示例
└── README.md              # 项目文档
```

## 🛠️ 技术栈

| 技术 | 版本 | 用途 |
|------|------|------|
| **Vue 3** | ^3.4.0 | 前端框架 |
| **Vite** | ^5.0.0 | 构建工具 |
| **Nginx** | 1.24+ | Web 服务器 |
| **Alpine Linux** | 3.18 | 基础镜像 |
| **Docker** | - | 容器化 |

## 🏭 生产特性

### 🚀 性能优化
- **多阶段构建** 最小化镜像大小 (14.4MB)
- **Gzip 压缩** 减少传输大小
- **静态资源缓存** 长期缓存策略
- **SPA 路由支持** Vue Router 兼容

### 🛡️ 安全配置
- **安全头部** XSS 保护、内容类型选项
- **隐藏文件保护** 防止访问敏感文件
- **HTTPS 就绪** 支持 SSL 配置

### 📊 监控运维
- **健康检查端点** `/health`
- **Docker 健康检查** 自动监控
- **日志配置** 访问和错误日志
- **优雅关闭** 容器生命周期管理

## 🔧 配置说明

### Nginx 配置

项目使用独立的 `nginx.conf` 配置文件，包含：

- 端口配置 (9000)
- SPA 路由支持
- 静态资源缓存
- Gzip 压缩
- 安全头部
- 健康检查端点

### Docker 配置

- **多阶段构建** 分离构建和运行环境
- **Alpine 基础镜像** 最小化镜像大小
- **非 root 用户** 安全运行
- **健康检查** 自动监控

## 🚨 故障排除

### 常见问题

1. **端口冲突**
   ```bash
   # 检查端口占用
   lsof -i :9000
   
   # 使用其他端口
   docker run -p 8080:9000 xvue
   ```

2. **构建失败**
   ```bash
   # 清理 Docker 缓存
   docker system prune -a
   
   # 重新构建
   docker build --no-cache -t xvue .
   ```

3. **权限问题**
   ```bash
   # 检查文件权限
   ls -la /usr/share/nginx/html
   ```

### 日志查看

```bash
# 查看容器日志
docker logs xvue-app

# 实时查看日志
docker logs -f xvue-app
```

## 📈 性能指标

| 指标 | 值 |
|------|-----|
| **镜像大小** | 14.4MB |
| **构建时间** | ~30s |
| **启动时间** | <2s |
| **内存占用** | ~10MB |
| **CPU 占用** | <1% |

## 🤝 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 📞 支持

如有问题或建议，请：

- 提交 [Issue](../../issues)
- 发送邮件至 [your-email@example.com]
- 查看 [Wiki](../../wiki) 获取更多文档
