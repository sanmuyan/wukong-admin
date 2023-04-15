package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"wukong/pkg/response"
	"wukong/server/model"
)

func UpdateUserProfile(c *gin.Context) {
	resp := response.NewResponse()
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		resp.Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	err := svc.UpdateUserProfile(&user, keysToUserToken(c.Keys))
	if err != nil {
		logrus.Errorf("更新用户配置: %s", err.Err)
		if err.IsResponseMsg {
			resp.FailWithMsg(response.HttpInternalServerError, err.Err.Error()).SetGin(c)
		} else {
			resp.Fail(response.HttpInternalServerError).SetGin(c)
		}
		return
	}
	resp.Ok().SetGin(c)
}

func GetUserProfile(c *gin.Context) {
	resp := response.NewResponse()
	userInfo, err := svc.GetUserProfile(keysToUserToken(c.Keys))
	if err != nil {
		logrus.Errorf("获取用户配置: %s", err.Err)
		resp.Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	resp.OkWithData(userInfo).SetGin(c)
}
