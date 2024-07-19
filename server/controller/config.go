package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"wukong/pkg/config"
	"wukong/pkg/util"
)

func GetBasicConfig(c *gin.Context) {
	util.Respf().Ok().WithData(svc.GetBasicConfig()).Response(util.GinRespf(c))
}

func UpdateBasicConfig(c *gin.Context) {
	var req config.Basic
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	svc.UpdateBasicConfig(&req)
	util.Respf().Ok().Response(util.GinRespf(c))
}

func GetSecurityConfig(c *gin.Context) {
	util.Respf().Ok().WithData(svc.GetSecurityConfig()).Response(util.GinRespf(c))
}

func UpdateSecurityConfig(c *gin.Context) {
	var req config.Security
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	svc.UpdateSecurityConfig(&req)
	util.Respf().Ok().Response(util.GinRespf(c))
}

func GetLDAPConfig(c *gin.Context) {
	util.Respf().Ok().WithData(svc.GetLDAPConfig()).Response(util.GinRespf(c))
}

func UpdateLDAPConfig(c *gin.Context) {
	var req config.LDAP
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	svc.UpdateLDAPConfig(&req)
	util.Respf().Ok().Response(util.GinRespf(c))
}

func GetOauthProvidersConfig(c *gin.Context) {
	util.Respf().Ok().WithData(svc.GetOauthProvidersConfig()).Response(util.GinRespf(c))
}

func UpdateOauthProvidersConfig(c *gin.Context) {
	var req config.OauthProviders
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	svc.UpdateOauthProvidersConfig(&req)
	util.Respf().Ok().Response(util.GinRespf(c))
}
