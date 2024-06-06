package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func CreateToken(c *gin.Context) {
	var token model.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(token.Username) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	tokenStr, err := svc.CreateOrSetToken(&token, token.Username, 0)
	if err != nil {
		logrus.Errorf("创建token: %s", err)
		util.Respf().FailWithError(util.NewRespError(err)).Response(util.GinRespf(c))
		return
	}
	util.Respf().WithData(tokenStr).Ok().Response(util.GinRespf(c))
}

func DeleteToken(c *gin.Context) {
	var token model.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(token.TokenType, token.Username) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.DeleteToken(&token); err != nil {
		logrus.Errorf("删除token: %s", err)
		util.Respf().FailWithError(util.NewRespError(err)).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
