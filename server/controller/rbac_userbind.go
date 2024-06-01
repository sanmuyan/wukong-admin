package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
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
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(userBinds).Response(util.GinRespf(c))
}

func CreateUserBinds(c *gin.Context) {
	var userBinds []model.UserBind
	if err := c.ShouldBindJSON(&userBinds); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).WithError(util.NewRespError(err)).Response(util.GinRespf(c))
		return
	}
	for _, userBind := range userBinds {
		if err := svc.CreateUserBind(&userBind); err != nil {
			logrus.Errorf("创建用户角色绑定: %s", err.Err)
			util.Respf().FailWithError(err).Response(util.GinRespf(c))
			return
		}
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteUserBinds(c *gin.Context) {
	var userBinds []model.UserBind
	if err := c.ShouldBindJSON(&userBinds); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).WithError(util.NewRespError(err)).Response(util.GinRespf(c))
		return
	}
	for _, userBind := range userBinds {
		if err := svc.DeleteUserBind(&userBind); err != nil {
			logrus.Errorf("删除用户角色绑定: %s", err.Err)
			util.Respf().FailWithError(err).Response(util.GinRespf(c))
			return
		}
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
