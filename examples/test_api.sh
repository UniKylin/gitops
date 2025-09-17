#!/bin/bash

# Todo List API 测试脚本

echo "=== Todo List API 测试 ==="
echo

# 测试获取所有todos
echo "1. 测试获取所有todos:"
curl -s http://localhost:9000/api/todos | jq '.'
echo
echo

# 测试创建新todo
echo "2. 测试创建新todo:"
curl -s -X POST http://localhost:9000/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Docker测试任务","description":"这是一个Docker部署测试任务"}' | jq '.'
echo
echo

# 再次获取todos查看新创建的
echo "3. 再次获取所有todos:"
curl -s http://localhost:9000/api/todos | jq '.'
echo
echo

# 测试更新todo状态
echo "4. 测试更新todo状态:"
curl -s -X PUT http://localhost:9000/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"completed":1}' | jq '.'
echo
echo

# 测试删除todo
echo "5. 测试删除todo:"
curl -s -X DELETE http://localhost:9000/api/todos/1 | jq '.'
echo
echo

echo "=== 测试完成 ==="
echo "前端访问地址: http://localhost:3000"
echo "API访问地址: http://localhost:9000/api"
