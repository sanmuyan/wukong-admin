package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func MFAFinishLogin(c *gin.Context) {
	var req model.MFAFinishLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	res, err := svc.MFAFinishLogin(&req)
	if err != nil {
		logrus.Errorf("MFA 完成登陆失败: %s", err.Err)
		util.Respf().FailWithError(err, xresponse.HttpUnauthorized).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(res).Response(util.GinRespf(c))
}
