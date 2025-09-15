# Multi-Arch Server

一个基于 Go 和 Gin 框架的轻量级多架构 Web 服务器，支持跨平台构建和部署。

## 项目概述

这是一个专门为多平台构建设计的 HTTP 服务器，提供系统信息查询接口。该项目的主要目的是演示如何使用 Docker Buildx 构建支持多种 CPU 架构的容器镜像，包括 AMD64、ARM64、ARMv7 等平台。服务器会返回当前运行环境的 Go 版本、架构和操作系统信息，帮助验证多平台部署的效果。

## 功能特性

- 🌐 基于 Gin 框架的 HTTP 服务器
- 📊 提供系统信息查询接口 (`GET /`)
- 🏗️ **多平台 Docker 镜像构建** - 支持 AMD64、ARM64、ARMv7 等架构
- 🚀 使用 scratch 基础镜像，体积最小化 (10-15MB)
- ⚡ 启用 Go 构建缓存，提升多平台构建速度
- 🔧 专门优化的 Dockerfile，支持交叉编译
- 📱 完美适配 Apple Silicon (M1/M2) 和 ARM 云服务器

## 技术栈

- **语言**: Go 1.23.9 (支持交叉编译)
- **Web 框架**: Gin v1.10.1
- **容器化**: Docker Buildx (多平台构建)
- **构建工具**: Docker Buildx + Go 交叉编译
- **基础镜像**: golang:1.23 (构建阶段) + scratch (运行阶段)
- **支持平台**: linux/amd64, linux/arm64, linux/arm/v7, linux/386

## 项目结构

```
multi-arch-server/
├── Dockerfile          # 多架构 Docker 构建文件
├── go.mod             # Go 模块依赖管理
├── go.sum             # Go 模块校验和
├── main.go            # 主程序入口
└── README.md          # 项目文档
```

## API 接口

### GET /

返回当前运行环境的系统信息。

**响应示例**:
```json
{
  "version": "go1.23.9",
  "arch": "amd64",
  "os": "linux"
}
```

## 本地开发

### 环境要求

- Go 1.23.9 或更高版本
- Docker (用于容器化部署)

### 运行步骤

1. **克隆项目**
   ```bash
   git clone <repository-url>
   cd multi-arch-server
   ```

2. **安装依赖**
   ```bash
   go mod download
   ```

3. **运行服务器**
   ```bash
   go run main.go
   ```

4. **访问服务**
   ```bash
   curl http://localhost:9000/
   ```

## 多平台 Docker 构建

本项目专门设计用于多平台构建，支持在单一代码库中构建适用于不同 CPU 架构的 Docker 镜像。

### 快速开始

```bash
# 1. 设置 Docker Buildx 环境
docker buildx create --name builder
docker buildx use builder
docker buildx inspect --bootstrap
docker login  # 如需要推送到远程仓库

# 2. 构建多平台镜像
docker buildx build --platform linux/amd64,linux/arm64 -t multi-arch-server:latest --load .

# 3. 运行并测试
docker run -p 9000:9000 multi-arch-server:latest
curl http://localhost:9000/
```

### 支持的平台架构

- **linux/amd64**: Intel/AMD 64位处理器
- **linux/arm64**: ARM 64位处理器 (如 Apple M1/M2, AWS Graviton)
- **linux/arm/v7**: ARM 32位处理器
- **linux/386**: Intel 32位处理器

### 多平台构建方法

#### 方法一：使用 Docker Buildx (推荐)

```bash
# 1. 创建多架构构建器
docker buildx create --name builder

# 2. 切换到构建器
docker buildx use builder

# 3. 检查构建器状态和平台支持
docker buildx inspect --bootstrap

# 4. 登录到镜像仓库 (如需要推送)
docker login

# 5. 构建并推送到镜像仓库
docker buildx build \
  --platform linux/amd64,linux/arm64,linux/arm/v7 \
  -t your-registry/multi-arch-server:latest \
  --push .

# 6. 构建到本地 (仅支持单一平台)
docker buildx build \
  --platform linux/amd64 \
  -t multi-arch-server:latest \
  --load .
```

#### 方法二：使用 docker buildx build 直接构建

```bash
# 构建多平台镜像并推送到仓库
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t multi-arch-server:latest \
  --push .

# 构建多平台镜像并保存到本地
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t multi-arch-server:latest \
  --output type=docker .
```

#### 方法三：使用 GitHub Actions 自动构建

```yaml
name: Multi-Platform Build
on:
  push:
    tags: ['v*']

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2
      
      - name: Build and push
        uses: docker/build-push-action@v4
        with:
          context: .
          platforms: linux/amd64,linux/arm64,linux/arm/v7
          push: true
          tags: |
            your-registry/multi-arch-server:latest
            your-registry/multi-arch-server:${{ github.ref_name }}
```

### 验证多平台镜像

```bash
# 查看构建器支持的平台
docker buildx ls

# 查看镜像支持的平台
docker buildx imagetools inspect multi-arch-server:latest

# 在特定平台运行容器
docker run --platform linux/arm64 -p 9000:9000 multi-arch-server:latest
docker run --platform linux/amd64 -p 9000:9000 multi-arch-server:latest

# 测试不同平台的系统信息
curl http://localhost:9000/
```

### Dockerfile 多平台构建特性

该 Dockerfile 专门为多平台构建优化：

1. **平台感知构建**: 使用 `FROM --platform=$BUILDPLATFORM` 确保构建环境匹配
2. **目标平台参数**: 通过 `ARG TARGETOS TARGETARCH` 获取目标平台信息
3. **交叉编译**: 使用 `GOOS=$TARGETOS GOARCH=$TARGETARCH` 进行交叉编译
4. **静态链接**: `CGO_ENABLED=0` 确保生成静态二进制，无需依赖系统库
5. **最小化镜像**: `scratch` 基础镜像确保跨平台兼容性
6. **构建缓存**: `--mount=type=cache` 提升多平台构建速度

## 快速部署

### 单平台部署

```bash
# 构建并运行当前平台镜像
docker build -t multi-arch-server .
docker run -p 9000:9000 multi-arch-server

# 访问服务
curl http://localhost:9000/
```

### 多平台部署

```bash
# 构建多平台镜像
docker buildx build --platform linux/amd64,linux/arm64 -t multi-arch-server:latest --load .

# 运行特定平台容器
docker run --platform linux/amd64 -p 9000:9000 multi-arch-server:latest
```

## 性能特点

- **镜像大小**: 约 10-15MB (scratch 基础镜像)
- **内存占用**: 约 8-16MB 运行时内存
- **启动时间**: < 1 秒
- **并发处理**: 支持高并发请求

## 监控和健康检查

服务器提供简单的健康检查端点，可以用于容器编排和监控：

```bash
# 健康检查
curl http://localhost:9000/

# 检查响应时间
curl -w "@curl-format.txt" -o /dev/null -s http://localhost:9000/
```

## 开发说明

### 添加新功能

1. 在 `main.go` 中添加新的路由处理函数
2. 更新依赖（如需要）: `go mod tidy`
3. 重新构建镜像: `docker build -t multi-arch-server .`

### 调试模式

开发时可以使用 Gin 的调试模式：

```go
// 在 main.go 中修改
gin.SetMode(gin.DebugMode)
```

## 许可证

本项目采用 MIT 许可证。

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目。

## 更新日志

- **v1.0.0**: 初始版本，支持多架构构建和系统信息查询
