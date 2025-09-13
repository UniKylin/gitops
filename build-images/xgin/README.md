# XGin 容器构建优化

## 第一次构建
```bash
$ docker images | grep xgin
xgin    1    d835e33f60c8   10 minutes ago   1.73GB
```
镜像体积惊人！

## 第二次构建
更换 `alpine` 镜像构建后体积缩小了一半
```bash
$ docker images | grep xgin               
xgin    2   077cd450cefd   20 seconds ago   863MB
xgin    1   d835e33f60c8   15 minutes ago   1.73GB
```


使用 Gin 框架构建的简单 HTTP API。

## 功能特性

- Hi Gin 接口
- Docker 支持

## API 接口

### GET /
返回 Hi Gin 消息。

**响应:**
```json
{
  "message": "Hi Gin!"
}
```

## 运行应用

### 本地开发

1. 安装依赖:
```bash
go mod download
```

2. 运行应用:
```bash
go run main.go
```

服务器将在 9000 端口启动。

### 使用 Docker

1. 构建 Docker 镜像:
```bash
docker build -t xgin .
```

2. 运行容器:
```bash
# 前台运行
docker run -p 9000:9000 xgin

# 后台运行（推荐）
docker run -d -p 9000:9000 --name xgin-app xgin
```

### Docker 容器管理

**查看运行中的容器：**
```bash
docker ps | grep xgin
```

**查看容器日志：**
```bash
docker logs xgin-app
```

**停止容器：**
```bash
docker stop xgin-app
```

**重启容器：**
```bash
docker restart xgin-app
```

**删除容器：**
```bash
docker rm xgin-app

```

**清空所有运行容器记录:**
```bash
docker container prune
```

## 测试 API

你可以使用 curl 测试 API:

```bash
# Hi Gin
curl http://localhost:9000/
```

## 项目结构

```
xgin/
├── main.go          # 主应用文件
├── go.mod           # Go 模块文件
├── go.sum           # Go 模块校验和
├── Dockerfile       # Docker 配置
└── README.md        # 本文件
```
