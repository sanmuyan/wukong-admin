package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
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
		util.Respf().Fail(xresponse.HttpBadRequest).WithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(roleBinds).Response(util.GinRespf(c))
}

func CreateRoleBinds(c *gin.Context) {
	var roleBinds []model.RoleBind
	if err := c.ShouldBindJSON(&roleBinds); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	for _, roleBind := range roleBinds {
		if xutil.IsZero(roleBind.ResourceID, roleBind.RoleID) {
			util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
			return
		}
		if err := svc.CreateRoleBind(&roleBind); err != nil {
			logrus.Errorf("创建角色绑定: %s", err.Err)
			util.Respf().FailWithError(err).Response(util.GinRespf(c))
			return
		}
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteRoleBinds(c *gin.Context) {
	var roleBinds []model.RoleBind
	if err := c.ShouldBindJSON(&roleBinds); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	for _, roleBind := range roleBinds {
		if xutil.IsZero(roleBind.ID) {
			util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
			return
		}
		if err := svc.DeleteRoleBind(&roleBind); err != nil {
			logrus.Errorf("删除角色绑定: %s", err.Err)
			util.Respf().FailWithError(err).Response(util.GinRespf(c))
			return
		}
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
