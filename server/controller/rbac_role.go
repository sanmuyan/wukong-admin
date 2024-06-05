package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func GetRoles(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"role_name"}
	roles, err := svc.GetRoles(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取角色列表: %v", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(roles).Response(util.GinRespf(c))
}

func CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.CreateRole(&role); err != nil {
		logrus.Errorf("创建角色: %v", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func UpdateRole(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.UpdateRole(&role); err != nil {
		logrus.Errorf("更新角色: %v", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteRole(c *gin.Context) {
	var role model.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(role.ID) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.DeleteRole(&role); err != nil {
		logrus.Errorf("删除角色: %v", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
