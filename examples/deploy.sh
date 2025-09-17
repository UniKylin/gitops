#!/bin/bash

# Todo List 项目构建和部署脚本

set -e

echo "🚀 开始构建Todo List项目..."

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 函数：打印彩色消息
print_message() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查Docker是否安装
if ! command -v docker &> /dev/null; then
    print_error "Docker未安装，请先安装Docker"
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    print_error "Docker Compose未安装，请先安装Docker Compose"
    exit 1
fi

# 停止现有容器
print_message "停止现有容器..."
docker-compose down

# 清理旧镜像（可选）
if [ "$1" = "--clean" ]; then
    print_warning "清理旧镜像..."
    docker-compose down --rmi all --volumes --remove-orphans
fi

# 构建并启动服务
print_message "构建并启动所有服务..."
docker-compose up --build -d

# 等待服务启动
print_message "等待服务启动..."
sleep 10

# 检查服务状态
print_message "检查服务状态..."
docker-compose ps

# 检查健康状态
print_message "检查服务健康状态..."
for i in {1..30}; do
    if curl -f http://localhost/api/todos > /dev/null 2>&1; then
        print_message "✅ 所有服务启动成功！"
        echo ""
        echo "🌐 访问地址："
        echo "   前端: http://localhost"
        echo "   后端API: http://localhost:9000/api/todos"
        echo "   数据库: localhost:5432"
        echo ""
        echo "📋 管理命令："
        echo "   查看日志: docker-compose logs -f"
        echo "   停止服务: docker-compose down"
        echo "   重启服务: docker-compose restart"
        echo "   清理所有: docker-compose down --rmi all --volumes"
        exit 0
    fi
    echo -n "."
    sleep 2
done

print_error "服务启动超时，请检查日志：docker-compose logs"
exit 1
