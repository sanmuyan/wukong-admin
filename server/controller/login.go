package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"wukong/pkg/response"
	"wukong/server/model"
)

func Login(c *gin.Context) {
	resp := response.NewResponse()
	var login model.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		resp.Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	token, err := svc.Login(login)
	if err != nil {
		logrus.Errorf("用户登陆失败: %s", err.Err)
		resp.Fail(response.HttpUnauthorized).SetGin(c)
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	resp.OkWithData(data).SetGin(c)
}

func Logout(c *gin.Context) {
	resp := response.NewResponse()
	if err := svc.Logout(keysToUserToken(c.Keys)); err != nil {
		logrus.Errorf("用户登出失败: %s", err.Err)
		resp.Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	resp.Ok().SetGin(c)
}
