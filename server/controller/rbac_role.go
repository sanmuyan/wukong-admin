package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"wukong/pkg/response"
	"wukong/server/model"
)

func GetRoles(c *gin.Context) {
	resp := response.NewResponse()
	likeKeys := ""
	mustKeys := []string{"role_name"}
	roles, err := svc.GetRoles(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取角色列表: %v", err.Err)
		resp.Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	resp.OkWithData(roles).SetGin(c)
}

func CreateRole(c *gin.Context) {
	resp := response.NewResponse()
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		resp.Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	if err := svc.CreateRole(&role); err != nil {
		logrus.Errorf("创建角色: %v", err.Err)
		resp.Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	resp.Ok().SetGin(c)
}

func UpdateRole(c *gin.Context) {
	resp := response.NewResponse()
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		resp.Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	if err := svc.UpdateRole(&role); err != nil {
		logrus.Errorf("更新角色: %v", err.Err)
		resp.Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	resp.Ok().SetGin(c)
}

func DeleteRole(c *gin.Context) {
	resp := response.NewResponse()
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		resp.Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	if err := svc.DeleteRole(&role); err != nil {
		logrus.Errorf("删除角色: %v", err.Err)
		resp.Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	resp.Ok().SetGin(c)
}
