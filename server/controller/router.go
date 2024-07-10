package controller

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"wukong/server/middleware"
)

func RunServer(ctx context.Context, addr string) {
	r := gin.Default()
	router(r)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			if err != nil {
				logrus.Fatalf("server run error: %s", err)
			}
		}
	}()
	<-ctx.Done()
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("server shutdown error: %s", err)
	}
	logrus.Warn("server has been shutdown")
}

func router(r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.AccessControl())
	{
		// 注销登录
		apiGroup.POST("/logout", Logout)
		// 注销所有登录
		apiGroup.POST("/logout/all", LogoutAll)

		// 用户管理
		apiGroup.GET("/user", GetUsers)
		apiGroup.POST("/user", CreateUser)
		apiGroup.PUT("/user", UpdateUser)
		apiGroup.DELETE("/user", DeleteUser)

		// 账号资料
		apiGroup.GET("/account/profile", GetProfile)
		apiGroup.PUT("/account/profile", UpdateProfile)

		// 获取 MFA 应用状态
		apiGroup.GET("/account/mfaAppStatus", GetMFAAppStatus)
		// 开始绑定 MFA 应用
		apiGroup.GET("/account/mfaAppBeginBind", MFAAppBeginBind)
		// 删除 MFA 应用
		apiGroup.DELETE("/account/mfaApp", DeleteMFAApp)
		// 获取账号通行密钥列表
		apiGroup.GET("/account/passKey", GetPassKeys)
		// 更新通行密钥备注
		apiGroup.PUT("/account/passKey", UpdatePassKey)
		// 删除通行密钥
		apiGroup.DELETE("/account/passKey", DeletePassKey)
		// 开始绑定通行密钥
		apiGroup.GET("/account/passKeyBeginRegistration", PassKeyBeginRegistration)
		// 完成绑定通行密钥
		apiGroup.POST("/account/passKeyFinishRegistration", PassKeyFinishRegistration)
		// 修改密码
		apiGroup.POST("/account/modifyPassword", ModifyPassword)

		// 角色管理
		apiGroup.GET("/role", GetRoles)
		apiGroup.POST("/role", CreateRole)
		apiGroup.PUT("/role", UpdateRole)
		apiGroup.DELETE("/role", DeleteRole)

		// API 资源管理
		apiGroup.GET("/resource", GetResources)
		apiGroup.POST("/resource", CreateResource)
		apiGroup.PUT("/resource", UpdateResource)
		apiGroup.DELETE("/resource", DeleteResource)

		// 角色权限管理
		apiGroup.GET("/roleBind", GetRoleBinds)
		apiGroup.POST("/roleBind", CreateRoleBinds)
		apiGroup.DELETE("/roleBind", DeleteRoleBinds)

		// 用户权限管理
		apiGroup.GET("/userBind", GetUserBinds)
		apiGroup.POST("/userBind", CreateUserBinds)
		apiGroup.DELETE("/userBind", DeleteUserBinds)

		// Token 管理
		apiGroup.POST("/token", CreateToken)
		apiGroup.DELETE("/token", DeleteTokenSession)

		// 同步 LDAP 用户
		apiGroup.POST("/ldap/user/sync", SyncLDAPUsers)
		// LDAP 连接测试
		apiGroup.POST("/ldap/connTest", LDAPConnTest)

		// OAuth 应用登录，从前端跳转
		apiGroup.GET("/oauth/authorize/frontRedirect", GetOauthCodeSessionFrontRedirect)

		// OAuth 应用管理
		apiGroup.GET("/app/oauth", GetOauthApps)
		apiGroup.POST("/app/oauth", CreateOauthApp)
		apiGroup.PUT("/app/oauth", UpdateOauthApp)
		apiGroup.DELETE("/app/oauth", DeleteOauthApp)

		// 基础设置
		apiGroup.GET("/settings/basic", GetBasicConfig)
		apiGroup.POST("/settings/basic", UpdateBasicConfig)

		// 安全设置
		apiGroup.GET("/settings/security", GetSecurityConfig)
		apiGroup.POST("/settings/security", UpdateSecurityConfig)

		// LDAP 设置
		apiGroup.GET("/settings/ldap", GetLDAPConfig)
		apiGroup.POST("/settings/ldap", UpdateLDAPConfig)

		// 第三方 OAuth2 登录设置
		apiGroup.GET("/settings/oauthProviders", GetOauthProvidersConfig)
		apiGroup.POST("/settings/oauthProviders", UpdateOauthProvidersConfig)

	}
	// 登录
	r.POST("/api/login", Login)
	// 完成 MFA 登录
	r.POST("/api/mfaFinishLogin", MFAFinishLogin)
	// 完成 MFA 绑定
	r.POST("/api/account/mfaAppFinishBind", MFAppFinishBind)
	// 开始通行密钥登录
	r.POST("/api/passKeyBeginLogin", PassKeyBeginLogin)
	// 完成通行密钥登录
	r.POST("/api/passKeyFinishLogin", PassKeyFinishLogin)
	// 通过第三方 OAuth2 登录
	r.GET("/api/oauth/login", OauthLogin)
	// 通过第三方OAuth2 登录回调
	r.GET("/api/oauth/callback", OauthCallback)
	// 获取 OAuth2 应用令牌
	r.POST("/api/oauth/token", GetOauthToken)
	// 吊销 OAuth2 应用令牌
	r.POST("/api/oauth/revoke", RevokeOauthToken)
	// OAuth2 应用登录入口
	r.GET("/api/oauth/authorize", GetOauthCodeSession)
	// 获取客户端加密公钥
	r.GET("/api/clientEncryptPublicKey", GetClientEncryptPublicKey)
}
