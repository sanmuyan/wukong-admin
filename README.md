# 悟空后台管理系统后端 Demo

## 项目介绍

实现了基本的登录鉴权、用户管理、角色管理、权限管理功能，适合作为后台管理系统的基础框架。

- 登录鉴权：基于JWT的登录鉴权，Redis存储Token
- 数据库：使用MySQL
- 数据库模型：使用GORM
- 权限模型：基于RBAC的权限模型

## 配套前端

https://github.com/sanmuyan/wukong-admin-web

## 快速启动

### 前提条件

- mysql 8.0
- redis 5.0
- go 1.21.0
- node 16.14.2
- nginx
- 一个长度为32的随机字符串，用于加密JWT的签名秘钥, 配置位置 secret.token_key
- 一个长度为32的随机字符串，用于CFB加解密JWT的签名秘钥

### 1. 下载项目

```shell
git clone https://github.com/sanmuyan/wukong-admin.git
git clone https://github.com/sanmuyan/wukong-admin-web.git
```

### 2. 初始化数据库

```bash
# 创建数据库
mysql -uroot -p <build/wukong-admin.sql
```

### 3. 启动后端

```shell
cd wukong-admin
go build cmd/server/server.go
./server -config/config.yaml -l -s :8081 --config-secret-key "CFB秘钥"
```

### 4. 启动前端

请参考前端项目的README.md

### 配置Nginx

```nginx
server {
    listen 80;
    server_name localhost;
    # 前端
    location / {
        # 如果是打包好的静态文件
        # root /wukong-admin-web/dist;
        # 历史路由模式需要设置路由回退
        # try_files $uri $uri/ /index.html;
        proxy_pass http://localhost:8080;
    }
    # 后端接口
    location /api {
        proxy_pass http://localhost:8081;
    }
}
```

### 访问

先导入测试数据
```sql
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/user', 1, '用户管理');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/account', -1, '个人账号');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/logout', -1, '退出登录');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/oauth/authorize', -1, 'OAuth 登录');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/account/profile', -1, '账号资料');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/account/mfaAppStatus', -1, 'MFA 应用');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/account/mfaAppBeginBind', -1, 'MFA 应用');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/account/mfaAppFinishBind', -1, 'MFA 应用');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/account/mfaApp', -1, 'MFA 应用');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/account/passKey', -1, '通行密钥');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/account/passKeyBeginRegistration', -1, '通行密钥');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/account/passKeyFinishRegistration', -1, '通行密钥');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/account/modifyPassword', -1, '修改密码');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/role', 1, '角色');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/resource', 1, '权限资源');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/roleBind', 1, '角色绑定资源');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/userBind', 1, '用户绑定角色');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/token', 1, '用户 Token');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/ldap/user/sync', 1, '同步 LDAP 用户');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/oauth/authorize/frontRedirect', -1, '应用登录');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/app/oauth', 1, 'OAuth 应用');
INSERT INTO `resources`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/logout/all', -1, '退出全部会话');

INSERT INTO `roles`(`id`, `role_name`, `access_level`, `comment`) VALUES (1, 'base', 1, '基本用户');
INSERT INTO `roles`(`id`, `role_name`, `access_level`, `comment`) VALUES (2, 'admin', 100, '管理员');

INSERT INTO `users`(`id`, `username`, `display_name`, `email`, `mobile`, `password`, `source`, `is_active`) VALUES (1, 'admin', '管理员', 'admin@qq.com', '13888888888', '$2a$04$WXVJ91k1yjGecUgfBgC3COnKstE.h4fdjV0bRc0TUpS4OoAAY0/7K', 'local', 1);


INSERT INTO `user_binds`(`role_id`, `user_id`) VALUES (1, 1);
```
http://localhost
