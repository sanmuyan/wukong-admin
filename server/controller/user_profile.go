package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"wukong/pkg/response"
	"wukong/server/model"
)

func UpdateUserProfile(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		respf().Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	err := svc.UpdateUserProfile(&user, keysToUserToken(c.Keys))
	if err != nil {
		logrus.Errorf("更新用户配置: %s", err.Err)
		if err.IsResponseMsg {
			respf().Fail(response.HttpInternalServerError).WithMsg(err.Err.Error()).SetGin(c)
		} else {
			respf().Fail(response.HttpInternalServerError).SetGin(c)
		}
		return
	}
	respf().Ok().SetGin(c)
}

func GetUserProfile(c *gin.Context) {
	userInfo, err := svc.GetUserProfile(keysToUserToken(c.Keys))
	if err != nil {
		logrus.Errorf("获取用户配置: %s", err.Err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().WithData(userInfo).SetGin(c)
}
