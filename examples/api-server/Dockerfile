# 使用官方Go镜像作为构建环境
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go mod文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o api-server .

# 使用alpine作为运行环境
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=builder /app/api-server .

# 复制初始化SQL文件
COPY --from=builder /app/init.sql .

# 暴露端口
EXPOSE 9000

# 运行应用
CMD ["./api-server"]
