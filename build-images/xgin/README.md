# XGin 容器构建优化

## 构建镜像
### 第一次
```bash
$ docker images | grep xgin
xgin    1    d835e33f60c8   10 minutes ago   1.73GB
```
镜像体积惊人！

### 第二次
更换 `alpine` 镜像构建后体积缩小了一半
```bash
$ docker images | grep xgin               
xgin    2   077cd450cefd   20 seconds ago   863MB
xgin    1   d835e33f60c8   15 minutes ago   1.73GB
```

### 第三次
打包构建工具全部放在了镜像中太大了，直接拷贝编译后的二进制文件到镜像中，另外把基础镜像换成`ubuntu:latest`

```bash
# 先在本地编译
go build -o xgin main.go 

docker build -t xgin:3 -f Dockerfile-3 .

$ docker images | grep xgin               
xgin    3   7444f8e3277e   5 seconds ago    156MB
xgin    2   077cd450cefd   11 minutes ago   863MB
xgin    1   d835e33f60c8   26 minutes ago   1.73GB
```
体积减小效果很明显了！
但是这种方式也存在一个问题，`Dockerfile` 本身就是用来构建打包和编译，编译放到本地了这不是一个很好的实践！

### 第四次构建：多阶段构建
也就是第一个步骤和第三个步骤放到一起，第一步编译、第二步拷贝执行命令!
到了这一步几乎可以用来做生产环境使用了！但是还是可以进一步压缩镜像体积！！！

### 第五步
从这个步骤开始几乎就是很疯狂了，一下几点都可以继续降低体积
* `alpine` 专门为容器定制化的 `Linux` 发行版很干净，当然更小
* 构建过程关闭各种依赖，比如 `CGO_ENABLED=0`
* 针对跨平台编译，比如 `GOOS=linux GOARCH=amd64`等等

```bash
$ docker images | grep xgin               
xgin    5     1e8c7f97d4d1   52 seconds ago   29.9MB
xgin    4     edc92069b913   10 minutes ago   156MB
xgin    3     7444f8e3277e   19 minutes ago   156MB
xgin    2     077cd450cefd   31 minutes ago   863MB
xgin    1     d835e33f60c8   46 minutes ago   1.73GB
```

### 第六步

更换空镜像 `scratch` 更疯狂，对于生产环境安全要求比较高的情况下可以使用！甚至日志都无法查看...
```bash
$ docker images | grep xgin               
xgin    6   cc89891f99a6   6 seconds ago    16.6MB
xgin    5   1e8c7f97d4d1   8 minutes ago    29.9MB
xgin    4   edc92069b913   17 minutes ago   156MB
xgin    3   7444f8e3277e   27 minutes ago   156MB
xgin    2   077cd450cefd   38 minutes ago   863MB
xgin    1   d835e33f60c8   54 minutes ago   1.73GB
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
