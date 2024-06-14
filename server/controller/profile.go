package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func UpdateProfile(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	err := svc.UpdateProfile(&user, keysToUserToken(c))
	if err != nil {
		logrus.Errorf("更新个人资料: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func GetProfile(c *gin.Context) {
	userInfo, err := svc.GetProfile(keysToUserToken(c))
	if err != nil {
		logrus.Errorf("获取个人资料: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(userInfo).Response(util.GinRespf(c))
}
