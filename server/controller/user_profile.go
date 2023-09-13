package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func UpdateUserProfile(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	err := svc.UpdateUserProfile(&user, keysToUserToken(c.Keys))
	if err != nil {
		logrus.Errorf("更新用户配置: %s", err.Err)
		if err.IsResponseMsg {
			util.Respf().Fail(xresponse.HttpInternalServerError).WithMsg(err.Err.Error()).Response(util.GinRespf(c))
		} else {
			util.Respf().Fail(xresponse.HttpInternalServerError).Response(util.GinRespf(c))
		}
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func GetUserProfile(c *gin.Context) {
	userInfo, err := svc.GetUserProfile(keysToUserToken(c.Keys))
	if err != nil {
		logrus.Errorf("获取用户配置: %s", err.Err)
		util.Respf().Fail(xresponse.HttpInternalServerError).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(userInfo).Response(util.GinRespf(c))
}
