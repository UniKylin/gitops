# Todo List 项目

一个完整的前后端分离的Todo List应用，使用Docker容器化部署。

## 项目架构

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   前端 (Vue3)   │    │   后端 (Go)     │    │   数据库 (PG)   │
│   Port: 3000    │───▶│   Port: 9000    │───▶│   Port: 5432    │
│   Nginx         │    │   Gin           │    │   PostgreSQL   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 技术栈

### 前端
- **Vue 3** - 渐进式JavaScript框架
- **Vite** - 快速构建工具
- **Axios** - HTTP客户端
- **Nginx** - Web服务器和反向代理

### 后端
- **Go** - 编程语言
- **Gin** - Web框架
- **PostgreSQL** - 关系型数据库
- **lib/pq** - PostgreSQL驱动

### 部署
- **Docker** - 容器化
- **Docker Compose** - 多容器编排

## 功能特性

- ✅ 查看Todo列表
- ✅ 创建新Todo
- ✅ 标记Todo完成状态
- ✅ 删除Todo
- ✅ PC端优化布局
- ✅ 响应式设计
- ✅ Docker容器化部署

## 快速开始

### 1. 环境要求

- Docker 20.10+
- Docker Compose 2.0+

### 2. 一键部署

```bash
# 进入项目根目录
cd /Users/test/gitops/examples

# 运行部署脚本
./deploy.sh
```

### 3. 手动部署

```bash
# 构建并启动所有服务
docker-compose up --build -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f
```

## 服务说明

### 服务详情

| 服务 | 容器名 | 端口 | 描述 |
|------|--------|------|------|
| fe | todo_frontend | 3000 | Vue3前端应用 |
| api-server | todo_api_server | 9000 | Go后端API服务 |
| postgres | todo_postgres | 5432 | PostgreSQL数据库 |

## 访问地址

- **前端应用**: http://localhost:3000
- **后端API**: http://localhost:9000/api
- **数据库**: localhost:5432

## API接口

### 获取Todo列表
```bash
GET /api/todos
```

### 创建Todo
```bash
POST /api/todos
Content-Type: application/json

{
  "title": "任务标题",
  "description": "任务描述"
}
```

### 更新Todo状态
```bash
PUT /api/todos/{id}
Content-Type: application/json

{
  "completed": 1
}
```

### 删除Todo
```bash
DELETE /api/todos/{id}
```

## 测试

运行API测试脚本：

```bash
./test_api.sh
```

## 管理命令

### 基本操作

```bash
# 启动服务
docker-compose up -d

# 停止服务
docker-compose down

# 重启服务
docker-compose restart

# 查看日志
docker-compose logs -f [service_name]

# 查看服务状态
docker-compose ps
```

### 开发调试

```bash
# 进入容器
docker exec -it todo_api_server sh
docker exec -it todo_frontend sh
docker exec -it todo_postgres psql -U postgres -d todo_db

# 查看实时日志
docker-compose logs -f api-server
docker-compose logs -f fe
docker-compose logs -f postgres

# 重新构建特定服务
docker-compose up --build -d api-server
docker-compose up --build -d fe
```

### 数据管理

```bash
# 备份数据库
docker exec todo_postgres pg_dump -U postgres todo_db > backup.sql

# 恢复数据库
docker exec -i todo_postgres psql -U postgres todo_db < backup.sql

# 清理所有数据（谨慎使用）
docker-compose down --volumes
```

## 环境变量

### API Server 环境变量

| 变量名 | 默认值 | 描述 |
|--------|--------|------|
| DB_HOST | localhost | 数据库主机 |
| DB_PORT | 5432 | 数据库端口 |
| DB_USER | postgres | 数据库用户名 |
| DB_PASSWORD | password | 数据库密码 |
| DB_NAME | todo_db | 数据库名称 |

### 自定义配置

创建 `.env` 文件来覆盖默认配置：

```bash
# .env
DB_HOST=your_db_host
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database
```

## 生产环境部署

### 1. 安全配置

```bash
# 修改默认密码
export DB_PASSWORD=your_secure_password

# 使用HTTPS
# 在nginx.conf中配置SSL证书
```

### 2. 性能优化

```bash
# 调整资源限制
docker-compose up -d --scale api-server=2

# 使用外部数据库
# 修改docker-compose.yml中的postgres服务配置
```

### 3. 监控和日志

```bash
# 配置日志轮转
# 添加监控工具（如Prometheus + Grafana）
# 设置健康检查告警
```

## 故障排除

### 常见问题

1. **端口冲突**
   ```bash
   # 检查端口占用
   lsof -i :3000
   lsof -i :9000
   lsof -i :5432
   ```

2. **数据库连接失败**
   ```bash
   # 检查数据库状态
   docker-compose logs postgres
   docker exec todo_postgres pg_isready -U postgres
   ```

3. **前端无法访问后端**
   ```bash
   # 检查网络连接
   docker network ls
   docker network inspect examples_todo_network
   ```

4. **服务启动失败**
   ```bash
   # 查看详细日志
   docker-compose logs [service_name]
   
   # 检查镜像构建
   docker images
   docker build -t test ./api-server
   ```

### 清理和重置

```bash
# 完全清理（删除所有容器、镜像、数据）
docker-compose down --rmi all --volumes --remove-orphans

# 清理Docker系统
docker system prune -a

# 重新部署
./deploy.sh --clean
```

## 开发指南

### 本地开发

```bash
# 只启动数据库
docker-compose up -d postgres

# 本地运行后端
cd api-server
go run main.go

# 本地运行前端
cd fe
npm install
npm run dev
```

### 代码更新

```bash
# 更新代码后重新构建
docker-compose up --build -d

# 或者只更新特定服务
docker-compose up --build -d api-server
```

## 项目结构

```
examples/
├── api-server/           # 后端Go项目
│   ├── Dockerfile
│   ├── .dockerignore
│   ├── main.go
│   ├── go.mod
│   └── init.sql
├── fe/                   # 前端Vue3项目
│   ├── Dockerfile
│   ├── .dockerignore
│   ├── nginx.conf
│   ├── package.json
│   └── src/
├── docker-compose.yml    # Docker编排文件
├── deploy.sh            # 部署脚本
├── test_api.sh          # API测试脚本
└── README.md           # 项目说明
```

## 技术支持

如有问题，请检查：
1. Docker和Docker Compose版本
2. 端口占用情况
3. 服务日志输出
4. 网络连接状态

## 许可证

MIT License
