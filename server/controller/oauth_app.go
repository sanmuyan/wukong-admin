package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func GetOauthAPPS(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"name"}
	oauthAPPS, err := svc.GetOauthAPPS(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取 OAuth APP: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
	}
	util.Respf().Ok().WithData(oauthAPPS).Response(util.GinRespf(c))
}

func CreateOauthAPP(c *gin.Context) {
	var oauthAPP model.OauthAPP
	if err := c.ShouldBindJSON(&oauthAPP); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(oauthAPP.APPName, oauthAPP.ClientID, oauthAPP.ClientSecret) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.CreateOauthAPP(&oauthAPP); err != nil {
		logrus.Errorf("创建 OAuth APP: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func UpdateOauthAPP(c *gin.Context) {
	var oauthAPP model.OauthAPP
	if err := c.ShouldBindJSON(&oauthAPP); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(oauthAPP.APPName, oauthAPP.ClientID, oauthAPP.ClientSecret) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.UpdateOauthAPP(&oauthAPP); err != nil {
		logrus.Errorf("更新 OAuth APP: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteOauthAPP(c *gin.Context) {
	var oauthAPP model.OauthAPP
	if err := c.ShouldBindJSON(&oauthAPP); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(oauthAPP.ID) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.DeleteOauthAPP(&oauthAPP); err != nil {
		logrus.Errorf("删除 OAuth APP: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
