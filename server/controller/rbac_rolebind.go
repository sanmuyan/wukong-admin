package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func GetRoleBinds(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"role_id", "resource_id"}
	roleBinds, err := svc.GetRoleBinds(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取角色绑定列表: %s", err.Err)
		util.Respf().Fail(xresponse.HttpInternalServerError).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(roleBinds).Response(util.GinRespf(c))
}

func CreateRoleBind(c *gin.Context) {
	var roleBind model.RoleBind
	if err := c.ShouldBindJSON(&roleBind); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.CreateRoleBind(&roleBind); err != nil {
		logrus.Errorf("创建角色绑定: %s", err.Err)
		util.Respf().Fail(xresponse.HttpInternalServerError).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteRoleBind(c *gin.Context) {
	var roleBind model.RoleBind
	if err := c.ShouldBindJSON(&roleBind); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.DeleteRoleBind(&roleBind); err != nil {
		logrus.Errorf("删除角色绑定: %s", err.Err)
		util.Respf().Fail(xresponse.HttpInternalServerError).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
