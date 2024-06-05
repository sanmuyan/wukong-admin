package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
)

func OAuthLogin(c *gin.Context) {
	if !IsMustQuery(c, "provider") {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	data := make(map[string]interface{})
	data["auth_url"] = svc.OAuthLogin(c.Query("provider"))
	util.Respf().Ok().WithData(data).Response(util.GinRespf(c))
}

func OAuthCallback(c *gin.Context) {
	if !IsMustQuery(c, "code", "state") {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	token, err := svc.OAuthCallback(c.Query("code"), c.Query("state"))
	if err != nil {
		logrus.Errorf("OAuth 回调失败: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	data := make(map[string]any)
	data["token"] = token
	util.Respf().Ok().WithData(data).Response(util.GinRespf(c))
}
