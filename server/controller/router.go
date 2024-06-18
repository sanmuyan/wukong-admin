package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"wukong/server/middleware"
)

func RunServer(addr string) {
	r := gin.Default()
	router(r)
	err := r.Run(addr)
	if err != nil {
		logrus.Fatalf("server run error: %s", err)
	}
}

func router(r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.AccessControl())
	{
		apiGroup.POST("/logout", Logout)
		apiGroup.POST("/logout/all", LogoutAll)

		apiGroup.GET("/user", GetUsers)
		apiGroup.POST("/user", CreateUser)
		apiGroup.PUT("/user", UpdateUser)
		apiGroup.DELETE("/user", DeleteUser)

		apiGroup.GET("/account/profile", GetProfile)
		apiGroup.PUT("/account/profile", UpdateProfile)

		apiGroup.GET("/account/mfaAppStatus", GetMFAAppStatus)
		apiGroup.GET("/account/mfaAppBeginBind", MFAAppBeginBind)
		apiGroup.POST("/account/mfaAppFinishBind", MFAppFinishBind)
		apiGroup.DELETE("/account/mfaApp", DeleteMFAApp)
		apiGroup.GET("/account/passKey", GetPassKeys)
		apiGroup.PUT("/account/passKey", UpdatePassKey)
		apiGroup.DELETE("/account/passKey", DeletePassKey)
		apiGroup.GET("/account/passKeyBeginRegistration", PassKeyBeginRegistration)
		apiGroup.POST("/account/passKeyFinishRegistration", PassKeyFinishRegistration)
		apiGroup.POST("/account/modifyPassword", ModifyPassword)

		apiGroup.GET("/role", GetRoles)
		apiGroup.POST("/role", CreateRole)
		apiGroup.PUT("/role", UpdateRole)
		apiGroup.DELETE("/role", DeleteRole)

		apiGroup.GET("/resource", GetResources)
		apiGroup.POST("/resource", CreateResource)
		apiGroup.PUT("/resource", UpdateResource)
		apiGroup.DELETE("/resource", DeleteResource)

		apiGroup.GET("/roleBind", GetRoleBinds)
		apiGroup.POST("/roleBind", CreateRoleBinds)
		apiGroup.DELETE("/roleBind", DeleteRoleBinds)

		apiGroup.GET("/userBind", GetUserBinds)
		apiGroup.POST("/userBind", CreateUserBinds)
		apiGroup.DELETE("/userBind", DeleteUserBinds)

		apiGroup.POST("/token", CreateToken)
		apiGroup.DELETE("/token", DeleteTokenSession)

		apiGroup.PUT("/ldap/user/sync", SyncLDAPUsers)

		apiGroup.GET("/oauth/authorize/frontRedirect", GetOauthCodeSessionFrontRedirect)

		apiGroup.GET("/app/oauth", GetOauthApps)
		apiGroup.POST("/app/oauth", CreateOauthApp)
		apiGroup.PUT("/app/oauth", UpdateOauthApp)
		apiGroup.DELETE("/app/oauth", DeleteOauthApp)

	}
	r.POST("/api/login", Login)
	r.POST("/api/mfaFinishLogin", MFAFinishLogin)
	r.POST("/api/passKeyBeginLogin", PassKeyBeginLogin)
	r.POST("/api/passKeyFinishLogin", PassKeyFinishLogin)
	r.GET("/api/oauth/login", OauthLogin)
	r.GET("/api/oauth/callback", OauthCallback)
	r.POST("/api/oauth/token", GetOauthToken)
	r.POST("/api/oauth/revoke", RevokeOauthToken)
	r.GET("/api/oauth/authorize", GetOauthCodeSession)
}
