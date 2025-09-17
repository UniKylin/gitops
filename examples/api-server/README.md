# Todo List API Server

基于Go Gin + PostgreSQL的简单Todo List API服务端项目。

## 功能特性

- ✅ 获取Todo列表
- ✅ 更新Todo完成状态
- ✅ 删除Todo
- ✅ CORS支持（前后端分离）
- ✅ 自动创建数据库表
- ✅ 示例数据

## 技术栈

- **后端框架**: Go Gin
- **数据库**: PostgreSQL
- **数据库驱动**: lib/pq

## 快速开始

### 1. 环境要求

- Go 1.19+
- PostgreSQL 12+

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 数据库配置

#### 使用Docker Compose启动PostgreSQL（推荐）

1. 确保已安装Docker和Docker Compose
2. 启动PostgreSQL服务：
```bash
docker-compose up -d
```

3. 等待数据库启动完成（约30秒），然后运行API服务：
```bash
go run main.go
```

#### 手动安装PostgreSQL

1. 安装并启动PostgreSQL
2. 创建数据库：
```sql
CREATE DATABASE todo_db;
```

3. 修改`main.go`中的数据库连接字符串（如果需要）：
```go
connStr := "user=your_username password=your_password dbname=todo_db sslmode=disable host=localhost port=5432"
```

**注意**：使用Docker Compose时，连接参数已预配置：
- 用户名：`postgres`
- 密码：`password`
- 数据库：`todo_db`
- 主机：`localhost`
- 端口：`5432`

### 4. 运行项目

```bash
go run main.go
```

服务器将在 `http://localhost:9000` 启动。

## API接口

### 响应格式说明

所有API响应都遵循统一格式：
- **成功响应**: `code: 0`
- **错误响应**: `code != 0`，包含`error`字段说明错误原因

**completed字段说明**：
- `0` - 未完成
- `1` - 已完成

### 1. 获取Todo列表
```
GET /api/todos
```

**响应示例:**
```json
{
  "code": 0,
  "data": [
    {
      "id": 1,
      "title": "学习Go语言",
      "description": "这是一个示例todo项目",
      "completed": 0,
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ]
}
```

### 2. 创建新的Todo
```
POST /api/todos
```

**请求体:**
```json
{
  "title": "新的Todo项目",
  "description": "这是描述信息"
}
```

**响应示例:**
```json
{
  "code": 0,
  "data": {
    "id": 5,
    "title": "新的Todo项目",
    "description": "这是描述信息",
    "completed": 0,
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  },
  "message": "创建成功"
}
```

### 3. 更新Todo完成状态
```
PUT /api/todos/:id
```

**请求体:**
```json
{
  "completed": 1
}
```

**响应示例:**
```json
{
  "code": 0,
  "data": {
    "id": 1,
    "title": "学习Go语言",
    "description": "这是一个示例todo项目",
    "completed": 1,
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T11:00:00Z"
  },
  "message": "更新成功"
}
```

### 4. 删除Todo
```
DELETE /api/todos/:id
```

**响应示例:**
```json
{
  "code": 0,
  "message": "删除成功"
}
```

## 项目结构

```
api-server/
├── main.go              # 主程序文件
├── init.sql             # 数据库初始化脚本
├── docker-compose.yml   # Docker Compose配置文件
├── go.mod               # Go模块文件
├── go.sum               # Go依赖锁定文件
└── README.md            # 项目说明文档
```

## 注意事项

1. **使用Docker Compose时**：确保Docker和Docker Compose已安装
2. **手动安装PostgreSQL时**：确保PostgreSQL服务正在运行
3. 项目会自动创建数据库表并插入示例数据
4. 所有API都支持CORS，适合前后端分离开发
5. Docker Compose会自动挂载`init.sql`文件进行数据库初始化

## 开发说明

- 项目结构简单，所有功能都在`main.go`中实现
- 使用Gin框架的中间件处理CORS
- 数据库操作使用原生SQL，简单直接
- 包含错误处理和状态码规范
