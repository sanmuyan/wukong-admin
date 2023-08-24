package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/dao/response"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func GetUserBinds(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"role_id", "user_id"}
	userBinds, err := svc.GetUserBinds(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取用户角色绑定: %s", err.Err)
		util.Respf().Fail(response.HttpInternalServerError).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(userBinds).Response(util.GinRespf(c))
}

func CreateUserBind(c *gin.Context) {
	var userBind model.UserBind
	if err := c.ShouldBindJSON(&userBind); err != nil {
		util.Respf().Fail(response.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.CreateUserBind(&userBind); err != nil {
		logrus.Errorf("创建用户角色绑定: %s", err.Err)
		util.Respf().Fail(response.HttpInternalServerError).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteUserBind(c *gin.Context) {
	var userBind model.UserBind
	if err := c.ShouldBindJSON(&userBind); err != nil {
		util.Respf().Fail(response.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.DeleteUserBind(&userBind); err != nil {
		logrus.Errorf("删除用户角色绑定: %s", err.Err)
		util.Respf().Fail(response.HttpInternalServerError).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
