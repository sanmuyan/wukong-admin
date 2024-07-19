package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
)

func OauthLogin(c *gin.Context) {
	if !isMustQuery(c, "provider") {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	res, err := svc.OauthLogin(c.Query("provider"))
	if err != nil {
		logrus.Errorf("OAuth 登录失败: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(res).Response(util.GinRespf(c))
}

func OauthCallback(c *gin.Context) {
	if !isMustQuery(c, "code", "state") {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	res, err := svc.OauthCallback(c.Query("code"), c.Query("state"))
	if err != nil {
		logrus.Errorf("OAuth 回调失败: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(res).Response(util.GinRespf(c))
}

func OauthBindCallback(c *gin.Context) {
	if !isMustQuery(c, "code", "state") {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	err := svc.OauthBindCallback(keysToUserToken(c), c.Query("code"), c.Query("state"))
	if err != nil {
		logrus.Errorf("OAuth 绑定失败: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func GetOauthBindStatus(c *gin.Context) {
	res, err := svc.GetOauthBindStatus(keysToUserToken(c))
	if err != nil {
		logrus.Errorf("获取 OAuth 绑定状态失败: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(res).Response(util.GinRespf(c))
}

func DeleteOauthBind(c *gin.Context) {
	if !isMustQuery(c, "provider") {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	err := svc.DeleteOauthBind(keysToUserToken(c), c.Query("provider"))
	if err != nil {
		logrus.Errorf("删除 OAuth 绑定失败: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
