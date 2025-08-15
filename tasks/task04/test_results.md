# 博客系统测试用例和预期结果

## 测试环境准备

1. 确保MySQL数据库正在运行
2. 设置环境变量或使用默认数据库配置:
   - DB_USER: root (默认)
   - DB_PASS: 空 (默认)
   - DB_HOST: localhost (默认)
   - DB_PORT: 3306 (默认)
   - DB_NAME: blog (默认)
3. 启动服务: `go run main.go`

## 测试用例及预期结果

### 1. 用户注册

**请求:**
```bash
curl -X POST http://localhost:8080/auth/register \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser","password":"password123","email":"test@example.com"}'
```

**预期响应:**
```json
{"message": "User registered successfully"}
```

**状态码:** 201 Created

### 2. 用户登录

**请求:**
```bash
curl -X POST http://localhost:8080/auth/login \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser","password":"password123"}'
```

**预期响应:**
```json
{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."}
```

**状态码:** 200 OK

### 3. 创建文章

**请求:**
```bash
curl -X POST http://localhost:8080/posts/ \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer <token>" \
     -d '{"title":"测试文章","content":"这是一篇测试文章的内容"}'
```

**预期响应:**
```json
{
  "ID": 1,
  "CreatedAt": "2023-01-01T00:00:00Z",
  "UpdatedAt": "2023-01-01T00:00:00Z",
  "DeletedAt": null,
  "title": "测试文章",
  "content": "这是一篇测试文章的内容",
  "user_id": 1
}
```

**状态码:** 201 Created

### 4. 获取所有文章

**请求:**
```bash
curl -X GET http://localhost:8080/posts/ \
     -H "Authorization: Bearer <token>"
```

**预期响应:**
```json
[
  {
    "ID": 1,
    "CreatedAt": "2023-01-01T00:00:00Z",
    "UpdatedAt": "2023-01-01T00:00:00Z",
    "DeletedAt": null,
    "title": "测试文章",
    "content": "这是一篇测试文章的内容",
    "user_id": 1,
    "User": {
      "ID": 1,
      "CreatedAt": "2023-01-01T00:00:00Z",
      "UpdatedAt": "2023-01-01T00:00:00Z",
      "DeletedAt": null,
      "Username": "testuser",
      "Password": "$2a$10$...",
      "Email": "test@example.com"
    }
  }
]
```

**状态码:** 200 OK

### 5. 获取单个文章

**请求:**
```bash
curl -X GET http://localhost:8080/posts/1 \
     -H "Authorization: Bearer <token>"
```

**预期响应:**
```json
{
  "ID": 1,
  "CreatedAt": "2023-01-01T00:00:00Z",
  "UpdatedAt": "2023-01-01T00:00:00Z",
  "DeletedAt": null,
  "title": "测试文章",
  "content": "这是一篇测试文章的内容",
  "user_id": 1,
  "User": {
    "ID": 1,
    "CreatedAt": "2023-01-01T00:00:00Z",
    "UpdatedAt": "2023-01-01T00:00:00Z",
    "DeletedAt": null,
    "Username": "testuser",
    "Password": "$2a$10$...",
    "Email": "test@example.com"
  }
}
```

**状态码:** 200 OK

### 6. 更新文章

**请求:**
```bash
curl -X PUT http://localhost:8080/posts/1 \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer <token>" \
     -d '{"title":"更新后的测试文章","content":"这是更新后的文章内容"}'
```

**预期响应:**
```json
{
  "ID": 1,
  "CreatedAt": "2023-01-01T00:00:00Z",
  "UpdatedAt": "2023-01-01T00:05:00Z",
  "DeletedAt": null,
  "title": "更新后的测试文章",
  "content": "这是更新后的文章内容",
  "user_id": 1
}
```

**状态码:** 200 OK

### 7. 创建评论

**请求:**
```bash
curl -X POST http://localhost:8080/comments/ \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer <token>" \
     -d '{"content":"这是一条评论","post_id":"1"}'
```

**预期响应:**
```json
{
  "ID": 1,
  "CreatedAt": "2023-01-01T00:00:00Z",
  "UpdatedAt": "2023-01-01T00:00:00Z",
  "DeletedAt": null,
  "content": "这是一条评论",
  "user_id": 1,
  "post_id": 1
}
```

**状态码:** 201 Created

### 8. 获取文章的所有评论

**请求:**
```bash
curl -X GET http://localhost:8080/comments/post/1 \
     -H "Authorization: Bearer <token>"
```

**预期响应:**
```json
[
  {
    "ID": 1,
    "CreatedAt": "2023-01-01T00:00:00Z",
    "UpdatedAt": "2023-01-01T00:00:00Z",
    "DeletedAt": null,
    "content": "这是一条评论",
    "user_id": 1,
    "post_id": 1,
    "User": {
      "ID": 1,
      "CreatedAt": "2023-01-01T00:00:00Z",
      "UpdatedAt": "2023-01-01T00:00:00Z",
      "DeletedAt": null,
      "Username": "testuser",
      "Password": "$2a$10$...",
      "Email": "test@example.com"
    }
  }
]
```

**状态码:** 200 OK

### 9. 权限验证测试

**请求:**
```bash
# 使用另一个用户登录获取的token尝试更新不属于自己的文章
curl -X PUT http://localhost:8080/posts/1 \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer <other_user_token>" \
     -d '{"title":"恶意更新","content":"尝试更新别人的文章"}'
```

**预期响应:**
```json
{"error": "You don't have permission to update this post"}
```

**状态码:** 403 Forbidden

### 10. 删除文章

**请求:**
```bash
curl -X DELETE http://localhost:8080/posts/1 \
     -H "Authorization: Bearer <token>"
```

**预期响应:**
```json
{"message": "Post deleted successfully"}
```

**状态码:** 200 OK

## 测试结果分析

通过以上测试用例，我们可以验证博客系统的以下功能：

1. **用户认证功能**
   - 用户可以成功注册和登录
   - 密码被正确加密存储
   - JWT token被正确生成和验证

2. **文章管理功能**
   - 用户可以创建、读取、更新和删除自己的文章
   - 文章列表和详情可以正确获取
   - 文章内容包含关联的用户信息

3. **评论功能**
   - 用户可以对文章发表评论
   - 可以获取指定文章的所有评论
   - 评论包含关联的用户信息

4. **权限控制**
   - 用户只能更新和删除自己创建的文章
   - 未授权用户无法访问需要认证的接口
   - 系统正确拒绝越权操作

## 注意事项

1. 测试前确保数据库服务正常运行
2. 每次测试前可能需要清理数据库中的测试数据
3. JWT token有过期时间（默认24小时），过期后需要重新登录获取
4. 测试过程中需要正确保存和使用token进行认证