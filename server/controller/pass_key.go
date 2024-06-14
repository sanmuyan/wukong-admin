package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func GetPassKeys(c *gin.Context) {
	passKeys, err := svc.GetPassKeys(keysToUserToken(c))
	if err != nil {
		logrus.Errorf("获取 passKeys 错误: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(passKeys).Response(util.GinRespf(c))
}

func UpdatePassKey(c *gin.Context) {
	var passKey model.PassKey
	if err := c.ShouldBindJSON(&passKey); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(passKey.ID) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(passKey.DisplayName) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	err := svc.UpdatePassKey(&passKey, keysToUserToken(c))
	if err != nil {
		logrus.Errorf("更新 passKey 错误: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeletePassKey(c *gin.Context) {
	var passKey model.PassKey
	if err := c.ShouldBindJSON(&passKey); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(passKey.ID) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	err := svc.DeletePassKey(&passKey, keysToUserToken(c))
	if err != nil {
		logrus.Errorf("删除 passKey 错误: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func PassKeyBeginRegistration(c *gin.Context) {
	res, err := svc.PassKeyBeginRegistration(keysToUserToken(c))
	if err != nil {
		logrus.Errorf("PassKey 开始注册错误: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(res).Response(util.GinRespf(c))
}

func PassKeyFinishRegistration(c *gin.Context) {
	err := svc.PassKeyFinishRegistration(keysToUserToken(c), c)
	if err != nil {
		logrus.Errorf("PassKey 完成注册错误: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func PassKeyBeginLogin(c *gin.Context) {
	var req model.PassKeyBeginLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	res, err := svc.BeginPassKeyLogin(&req)
	if err != nil {
		logrus.Errorf("PassKey 开始登录错误: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(res).Response(util.GinRespf(c))
}

func PassKeyFinishLogin(c *gin.Context) {
	var req model.PassKeyFinishLoginRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	res, err := svc.FinishPassKeyLogin(&req, c)
	if err != nil {
		logrus.Errorf("PassKey 完成登录错误: %s", err.Err)
		util.Respf().FailWithError(err, xresponse.HttpUnauthorized).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(res).Response(util.GinRespf(c))
}
