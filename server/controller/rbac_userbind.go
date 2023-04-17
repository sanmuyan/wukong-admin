package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"wukong/pkg/response"
	"wukong/server/model"
)

func GetUserBinds(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"role_id", "user_id"}
	userBinds, err := svc.GetUserBinds(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取用户角色绑定: %s", err.Err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().WithData(userBinds).SetGin(c)
}

func CreateUserBind(c *gin.Context) {
	var userBind model.UserBind
	if err := c.ShouldBindJSON(&userBind); err != nil {
		respf().Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	if err := svc.CreateUserBind(&userBind); err != nil {
		logrus.Errorf("创建用户角色绑定: %s", err.Err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().SetGin(c)
}

func DeleteUserBind(c *gin.Context) {
	var userBind model.UserBind
	if err := c.ShouldBindJSON(&userBind); err != nil {
		respf().Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	if err := svc.DeleteUserBind(&userBind); err != nil {
		logrus.Errorf("删除用户角色绑定: %s", err.Err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().SetGin(c)
}
