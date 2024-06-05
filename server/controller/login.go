package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func Login(c *gin.Context) {
	var login model.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	token, err := svc.Login(login)
	if err != nil {
		logrus.Errorf("用户登陆失败: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	util.Respf().Ok().WithData(data).Response(util.GinRespf(c))
}

func Logout(c *gin.Context) {
	if err := svc.Logout(keysToUserToken(c)); err != nil {
		logrus.Errorf("用户登出失败: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
