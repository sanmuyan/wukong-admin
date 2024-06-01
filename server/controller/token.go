package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func CreateToken(c *gin.Context) {
	var token model.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).WithError(util.NewRespError(err)).Response(util.GinRespf(c))
		return
	}
	tokenStr, err := svc.CreateOrSetToken(&token)
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
		util.Respf().Fail(xresponse.HttpBadRequest).WithError(util.NewRespError(err)).Response(util.GinRespf(c))
		return
	}
	tokenStr, err := svc.CreateOrSetToken(&token)
	if err != nil {
		logrus.Errorf("创建token: %s", err)
		util.Respf().FailWithError(util.NewRespError(err)).Response(util.GinRespf(c))
		return
	}
	util.Respf().WithData(tokenStr).Ok().Response(util.GinRespf(c))
}
