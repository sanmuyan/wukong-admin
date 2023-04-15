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
- go 1.18
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
        proxy_pass http://localhost:8080;
    }
    # 后端接口
    location /api {
        proxy_pass http://localhost:8080;
    }
}
```

### 访问

先导入测试数据
```sql
INSERT INTO `wukong`.`wk_rbac_resource`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/user', 1, '用户');
INSERT INTO `wukong`.`wk_rbac_resource`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/user/owner', 0, '用户个人信息');
INSERT INTO `wukong`.`wk_rbac_resource`(`resource_path`, `is_auth`, `comment`) VALUES ('/api/logout', 0, '退出登录');

INSERT INTO `wukong`.`wk_rbac_role`(`id`, `role_name`, `access_level`, `comment`) VALUES (1, 'base', 1, '基本用户');
INSERT INTO `wukong`.`wk_rbac_role`(`id`, `role_name`, `access_level`, `comment`) VALUES (2, 'admin', 100, '管理员');

INSERT INTO `wukong`.`wk_user`(`id`, `username`, `display_name`, `email`, `mobile`, `password`, `source`, `is_active`, `create_time`, `update_time`) VALUES (1, 'admin', '管理员', 'admin@qq.com', '13888888888', '$2a$04$WXVJ91k1yjGecUgfBgC3COnKstE.h4fdjV0bRc0TUpS4OoAAY0/7K', 'local', 1, '2023-04-14 16:44:07', '2023-04-15 13:17:28');


INSERT INTO `wukong`.`wk_rbac_user_bind`(`role_id`, `user_id`) VALUES (1, 1);
```
http://localhost
