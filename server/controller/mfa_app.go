package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func GetMFAAppStatus(c *gin.Context) {
	res, err := svc.GetMFAAppStatus(keysToUserToken(c))
	if err != nil {
		logrus.Errorf("获取用户 MFA 绑定状态错误: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(res).Response(util.GinRespf(c))
}

func MFAAppBeginBind(c *gin.Context) {
	res, err := svc.MFAAppBeginBind(keysToUserToken(c))
	if err != nil {
		logrus.Errorf("MFA开始绑定错误: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(res).Response(util.GinRespf(c))
}

func MFAppFinishBind(c *gin.Context) {
	var req model.MFAAppBindRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.MFAAppFinishBind(&req); err != nil {
		logrus.Errorf("MFA 完成绑定错误: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteMFAApp(c *gin.Context) {
	if err := svc.DeleteMFAApp(keysToUserToken(c)); err != nil {
		logrus.Errorf("删除用户 MFA 绑定错误: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
