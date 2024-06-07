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
	authURL, err := svc.OauthLogin(c.Query("provider"))
	if err != nil {
		logrus.Errorf("OAuth 登录失败: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	data := make(map[string]interface{})
	data["auth_url"] = authURL
	util.Respf().Ok().WithData(data).Response(util.GinRespf(c))
}

func OauthCallback(c *gin.Context) {
	if !isMustQuery(c, "code", "state") {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	token, err := svc.OauthCallback(c.Query("code"), c.Query("state"))
	if err != nil {
		logrus.Errorf("OAuth 回调失败: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	data := make(map[string]any)
	data["token"] = token
	util.Respf().Ok().WithData(data).Response(util.GinRespf(c))
}
