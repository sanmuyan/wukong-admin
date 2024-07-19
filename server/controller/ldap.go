package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/config"
	"wukong/pkg/util"
	"wukong/server/model"
)

func SyncLDAPUsers(c *gin.Context) {
	msg, err := svc.SyncLDAPUsers()
	if err != nil {
		logrus.Errorf("同步LDAP用户: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithMsg(msg).Response(util.GinRespf(c))
}

func LDAPConnTest(c *gin.Context) {
	var req config.LDAP
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	err := svc.LDAPConnTest(&req)
	if err != nil {
		logrus.Errorf("测试 LDAP 连接: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func LDAPLoginTest(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	err := svc.LDAPLoginTest(&req)
	if err != nil {
		logrus.Errorf("测试 LDAP 用户登录: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
