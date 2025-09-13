# =============================================================================
# Vue 3 应用的多阶段 Docker 构建
# 第一阶段：构建 Vue 应用
# =============================================================================
FROM node:20-alpine AS builder

# 设置构建阶段的工作目录
WORKDIR /app

# 先复制 package 文件以获得更好的 Docker 层缓存
COPY package*.json ./

# 使用 npm ci 安装依赖，更快更可靠的构建
RUN npm ci --only=production=false

# 复制源代码
COPY . .

# 构建 Vue 应用的生产版本
RUN npm run build

# =============================================================================
# 第二阶段：生产就绪的 nginx 服务器
# =============================================================================
FROM alpine:3.18

# 安装 nginx 并清理不必要的文件以最小化镜像大小
RUN apk add --no-cache nginx && \
    rm -rf /var/cache/apk/* /tmp/* /var/tmp/* /usr/share/nginx/html/* /etc/nginx/conf.d/default.conf /etc/nginx/nginx.conf

# 从构建阶段复制构建好的应用
COPY --from=builder /app/dist /usr/share/nginx/html

# 复制自定义的 nginx 配置
COPY nginx.conf /etc/nginx/nginx.conf

# 设置正确的文件所有权以便 nginx 提供服务
RUN chown -R nginx:nginx /usr/share/nginx/html

# 暴露 9000 端口用于 web 服务器
EXPOSE 9000

# 健康检查以确保容器正常运行
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:9000/health || exit 1

# 以前台模式启动 nginx
CMD ["nginx", "-g", "daemon off;"]
