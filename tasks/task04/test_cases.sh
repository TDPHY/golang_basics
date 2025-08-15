#!/bin/bash

# 博客系统API测试脚本
# 测试前请确保服务已在8080端口启动: go run main.go

echo "开始测试博客系统API..."

# 服务器地址
BASE_URL="http://localhost:8080"

# 1. 用户注册测试
echo "=== 测试1: 用户注册 ==="
curl -X POST $BASE_URL/auth/register \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser1","password":"password123","email":"test1@example.com"}'
echo -e "\n"

# 注册第二个用户用于测试权限
curl -X POST $BASE_URL/auth/register \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser2","password":"password456","email":"test2@example.com"}'
echo -e "\n"

# 2. 用户登录测试
echo "=== 测试2: 用户登录 ==="
LOGIN_RESPONSE=$(curl -s -X POST $BASE_URL/auth/login \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser1","password":"password123"}')
echo $LOGIN_RESPONSE

# 提取token用于后续测试
TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*"' | sed 's/"token":"//' | sed 's/"$//')
echo "Token: $TOKEN"
echo -e "\n"

# 3. 获取用户信息测试
echo "=== 测试3: 获取用户信息 ==="
# 先获取用户ID
echo "获取用户信息需要用户ID，该功能在当前实现中可能存在问题"
echo -e "\n"

# 4. 创建文章测试
echo "=== 测试4: 创建文章 ==="
CREATE_POST_RESPONSE=$(curl -s -X POST $BASE_URL/posts/ \
     -H "Content-Type: application/json; charset=utf-8" \
     -H "Authorization: Bearer $TOKEN" \
     --data-binary '{"title":"测试文章","content":"这是一篇测试文章的内容"}' \
     --trace-ascii trace.log)
echo $CREATE_POST_RESPONSE
# 查看 trace.log 文件内容
cat trace.log
echo ""


# 提取文章ID用于后续测试
POST_ID=$(echo $CREATE_POST_RESPONSE | grep -o '"ID":[0-9]*' | grep -o '[0-9]*')
echo "创建的文章ID: $POST_ID"
echo ""

# 5. 获取所有文章测试
echo "=== 测试5: 获取所有文章 ==="
curl -s -X GET $BASE_URL/posts/ \
     -H "Authorization: Bearer $TOKEN"
echo -e "\n"

# 6. 获取单个文章测试
echo "=== 测试6: 获取单个文章 ==="
curl -s -X GET $BASE_URL/posts/$POST_ID \
     -H "Authorization: Bearer $TOKEN"
echo -e "\n"

# 7. 更新文章测试
echo "=== 测试7: 更新文章 ==="
curl -s -X PUT $BASE_URL/posts/$POST_ID \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer $TOKEN" \
     -d '{"title":"更新后的测试文章","content":"这是更新后的文章内容"}'
echo -e "\n"

# 8. 创建评论测试
echo "=== 测试8: 创建评论 ==="
# 确保POST_ID不为空
if [ -z "$POST_ID" ]; then
    echo "错误: 无法提取文章ID"
    echo -e "\n"
else
    curl -s -X POST $BASE_URL/comments/ \
         -H "Content-Type: application/json" \
         -H "Authorization: Bearer $TOKEN" \
         -d "{\"content\":\"这是一条评论\",\"post_id\":$POST_ID}"
    echo -e "\n"
fi

# 9. 获取文章的所有评论测试
echo "=== 测试9: 获取文章的所有评论 ==="
curl -s -X GET $BASE_URL/comments/post/$POST_ID \
     -H "Authorization: Bearer $TOKEN"
echo -e "\n"

# 10. 权限测试 - 尝试更新不属于自己的文章
echo "=== 测试10: 权限验证（使用另一个用户尝试更新文章）==="

# 用第二个用户登录
LOGIN_RESPONSE2=$(curl -s -X POST $BASE_URL/auth/login \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser2","password":"password456"}')
TOKEN2=$(echo $LOGIN_RESPONSE2 | grep -o '"token":"[^"]*"' | sed 's/"token":"//' | sed 's/"$//')

# 尝试更新第一个用户的文章
curl -s -X PUT $BASE_URL/posts/$POST_ID \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer $TOKEN2" \
     -d '{"title":"恶意更新","content":"尝试更新别人的文章"}'
echo -e "\n"

# 11. 删除文章测试
echo "=== 测试11: 删除文章 ==="
# 使用正确的用户token删除文章
curl -s -X DELETE $BASE_URL/posts/$POST_ID \
     -H "Authorization: Bearer $TOKEN"
echo -e "\n"

echo "=== 测试完成 ==="