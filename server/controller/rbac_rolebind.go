package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"wukong/pkg/response"
	"wukong/server/model"
)

func GetRoleBinds(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"role_id", "resource_id"}
	roleBinds, err := svc.GetRoleBinds(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取角色绑定列表: %s", err.Err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().WithData(roleBinds).SetGin(c)
}

func CreateRoleBind(c *gin.Context) {
	var roleBind model.RoleBind
	if err := c.ShouldBindJSON(&roleBind); err != nil {
		respf().Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	if err := svc.CreateRoleBind(&roleBind); err != nil {
		logrus.Errorf("创建角色绑定: %s", err.Err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().SetGin(c)
}

func DeleteRoleBind(c *gin.Context) {
	var roleBind model.RoleBind
	if err := c.ShouldBindJSON(&roleBind); err != nil {
		respf().Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	if err := svc.DeleteRoleBind(&roleBind); err != nil {
		logrus.Errorf("删除角色绑定: %s", err.Err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().SetGin(c)
}
