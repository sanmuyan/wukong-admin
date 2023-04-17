package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"wukong/pkg/response"
	"wukong/server/model"
)

func Login(c *gin.Context) {
	var login model.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		respf().Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	token, err := svc.Login(login)
	if err != nil {
		logrus.Errorf("用户登陆失败: %s", err.Err)
		if err.IsResponseMsg {
			respf().Fail(response.HttpUnauthorized).WithMsg(err.Err.Error()).SetGin(c)
		} else {
			respf().Fail(response.HttpInternalServerError).SetGin(c)
		}
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	respf().Ok().WithData(data).SetGin(c)
}

func Logout(c *gin.Context) {
	if err := svc.Logout(keysToUserToken(c.Keys)); err != nil {
		logrus.Errorf("用户登出失败: %s", err.Err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().SetGin(c)
}
