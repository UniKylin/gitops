# syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM golang:1.23 AS builder
ARG TARGETOS TARGETARCH

WORKDIR /app

# 先复制 mod 文件并拉依赖
COPY go.mod go.sum ./
RUN go mod download

# 再复制源码
COPY . .

# 构建二进制，启用缓存，加上 CGO_ENABLED=0 保证跨平台
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -o /app/server .

# 最小化镜像
FROM scratch
WORKDIR /app
COPY --from=builder /app/server .
CMD ["/app/server"]

