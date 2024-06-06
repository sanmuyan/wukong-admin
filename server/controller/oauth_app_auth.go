package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
)

func GetOauthCode(c *gin.Context) {
	if !isMustQuery(c, "client_id", "redirect_uri", "response_type") {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	redirectURI, err := svc.GetOauthCode(
		keysToUserToken(c),
		c.Query("client_id"),
		c.Query("redirect_uri"),
		c.Query("response_type"),
		c.Query("scope"),
		c.Query("state"))
	if err != nil {
		logrus.Errorf("获取 OAuth code: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	data := make(map[string]string)
	data["redirect_uri"] = redirectURI
	util.Respf().Ok().WithData(data).Response(util.GinRespf(c))
}

func GetOauthToken(c *gin.Context) {
	if !isMustQuery(c, "code", "client_id", "redirect_uri", "grant_type") {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	token, err := svc.GetOauthToken(
		c.Query("code"),
		c.Query("client_id"),
		c.Query("client_secret"),
		c.Query("redirect_uri"),
		c.Query("grant_type"),
	)
	if err != nil {
		logrus.Errorf("获取 OAuth token: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	c.JSON(200, token)
}

func RefreshOauthToken(c *gin.Context) {
	if !isMustQuery(c, "client_id", "refresh_token", "grant_type") {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	token, err := svc.RefreshOauthToken(c.Query("refresh_token"), c.Query("client_id"), c.Query("grant_type"), c.Query("client_secret"))
	if err != nil {
		logrus.Errorf("刷新 OAuth token: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	c.JSON(200, token)
}

func RevokeOauthToken(c *gin.Context) {
	if !isMustQuery(c, "client_id", "token") {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	err := svc.RevokeOauthToken(c.Query("token"), c.Query("client_id"), c.Query("client_secret"))
	if err != nil {
		logrus.Errorf("删除 Oauth token: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
