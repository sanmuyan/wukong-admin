package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func GetUsers(c *gin.Context) {
	likeKeys := "username,email,mobile,display_name"
	mustKeys := []string{"id", "username"}
	users, err := svc.GetUsers(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取用户列表: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(users).Response(util.GinRespf(c))
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(user.Username) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.CreateUser(&user); err != nil {
		logrus.Errorf("创建用户: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(user.ID) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.UpdateUser(&user); err != nil {
		logrus.Errorf("更新用户: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(user.ID) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.DeleteUser(&user); err != nil {
		logrus.Errorf("删除用户: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
