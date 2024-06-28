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

		apiGroup.POST("/ldap/user/sync", SyncLDAPUsers)
		apiGroup.POST("/ldap/connTest", LDAPConnTest)

		apiGroup.GET("/oauth/authorize/frontRedirect", GetOauthCodeSessionFrontRedirect)

		apiGroup.GET("/app/oauth", GetOauthApps)
		apiGroup.POST("/app/oauth", CreateOauthApp)
		apiGroup.PUT("/app/oauth", UpdateOauthApp)
		apiGroup.DELETE("/app/oauth", DeleteOauthApp)

		apiGroup.GET("/settings/basic", GetBasicConfig)
		apiGroup.POST("/settings/basic", UpdateBasicConfig)

		apiGroup.GET("/settings/security", GetSecurityConfig)
		apiGroup.POST("/settings/security", UpdateSecurityConfig)

		apiGroup.GET("/settings/ldap", GetLDAPConfig)
		apiGroup.POST("/settings/ldap", UpdateLDAPConfig)

		apiGroup.GET("/settings/oauthProviders", GetOauthProvidersConfig)
		apiGroup.POST("/settings/oauthProviders", UpdateOauthProvidersConfig)

	}
	r.POST("/api/login", Login)
	r.POST("/api/mfaFinishLogin", MFAFinishLogin)
	r.POST("/api/account/mfaAppFinishBind", MFAppFinishBind)
	r.POST("/api/passKeyBeginLogin", PassKeyBeginLogin)
	r.POST("/api/passKeyFinishLogin", PassKeyFinishLogin)
	r.GET("/api/oauth/login", OauthLogin)
	r.GET("/api/oauth/callback", OauthCallback)
	r.POST("/api/oauth/token", GetOauthToken)
	r.POST("/api/oauth/revoke", RevokeOauthToken)
	r.GET("/api/oauth/authorize", GetOauthCodeSession)
	r.GET("/api/clientEncryptPublicKey", GetClientEncryptPublicKey)
}
