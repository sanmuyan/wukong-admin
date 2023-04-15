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
		logrus.Fatal(err)
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

		apiGroup.GET("/user/profile", GetUserProfile)
		apiGroup.PUT("/user/profile", UpdateUserProfile)

		apiGroup.GET("/role", GetRoles)
		apiGroup.POST("/role", CreateRole)
		apiGroup.PUT("/role", UpdateRole)
		apiGroup.DELETE("/role", DeleteRole)

		apiGroup.GET("/resource", GetResources)
		apiGroup.POST("/resource", CreateResource)
		apiGroup.PUT("/resource", UpdateResource)
		apiGroup.DELETE("/resource", DeleteResource)

		apiGroup.GET("/roleBind", GetRoleBinds)
		apiGroup.POST("/roleBind", CreateRoleBind)
		apiGroup.DELETE("/roleBind", DeleteRoleBind)

		apiGroup.GET("/userBind", GetUserBinds)
		apiGroup.POST("/userBind", CreateUserBind)
		apiGroup.DELETE("/userBind", DeleteUserBind)
	}
	r.POST("/api/login", Login)
}
