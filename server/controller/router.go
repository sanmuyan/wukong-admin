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

		apiGroup.GET("/user", GetUsers)
		apiGroup.POST("/user", CreateUser)
		apiGroup.PUT("/user", UpdateUser)
		apiGroup.DELETE("/user", DeleteUser)

		apiGroup.GET("/profile", GetProfile)
		apiGroup.PUT("/profile", UpdateProfile)

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
		apiGroup.DELETE("/token", DeleteToken)

		apiGroup.PUT("/ldap/user/sync", SyncLDAPUsers)

		apiGroup.GET("/oauth/authorize", GetOauthCode)

		apiGroup.GET("/oauth/app", GetOauthAPPS)
		apiGroup.POST("/oauth/app", CreateOauthAPP)
		apiGroup.PUT("/oauth/app", UpdateOauthAPP)
		apiGroup.DELETE("/oauth/app", DeleteOauthAPP)

	}
	r.POST("/api/login", Login)
	r.GET("/api/oauth/login", OauthLogin)
	r.GET("/api/oauth/callback", OauthCallback)
	r.GET("/api/oauth/token", GetOauthToken)
	r.POST("/api/oauth/token", RefreshOauthToken)
	r.POST("/api/oauth/revoke", RevokeOauthToken)
}
