package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func GetOauthApps(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"name"}
	oauthApps, err := svc.GetOauthApps(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取 OAuth APP: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
	}
	util.Respf().Ok().WithData(oauthApps).Response(util.GinRespf(c))
}

func CreateOauthApp(c *gin.Context) {
	var oauthApp model.OauthApp
	if err := c.ShouldBindJSON(&oauthApp); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(oauthApp.AppName) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.CreateOauthApp(&oauthApp); err != nil {
		logrus.Errorf("创建 OAuth APP: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func UpdateOauthApp(c *gin.Context) {
	var oauthApp model.OauthApp
	if err := c.ShouldBindJSON(&oauthApp); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(oauthApp.AppName, oauthApp.ClientID, oauthApp.ClientSecret) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.UpdateOauthApp(&oauthApp); err != nil {
		logrus.Errorf("更新 OAuth APP: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteOauthApp(c *gin.Context) {
	var oauthApp model.OauthApp
	if err := c.ShouldBindJSON(&oauthApp); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(oauthApp.ID) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.DeleteOauthApp(&oauthApp); err != nil {
		logrus.Errorf("删除 OAuth APP: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
