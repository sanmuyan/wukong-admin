package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func GetResources(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"resource_path"}
	resources, err := svc.GetResources(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取API资源: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
	}
	util.Respf().Ok().WithData(resources).Response(util.GinRespf(c))
}

func CreateResource(c *gin.Context) {
	var resource model.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(resource.ResourcePath) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.CreateResource(&resource); err != nil {
		logrus.Errorf("创建API资源: %s", err.Err)
		util.Respf().Fail(xresponse.HttpBadRequest).WithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func UpdateResource(c *gin.Context) {
	var resource model.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(resource.ID) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.UpdateResource(&resource); err != nil {
		logrus.Errorf("更新API资源: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteResource(c *gin.Context) {
	var resource model.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(resource.ID) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.DeleteResource(&resource); err != nil {
		logrus.Errorf("删除API资源: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

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

func GetRoleResourceBinds(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"role_id", "resource_id"}
	roleResourceBinds, err := svc.GetRoleResourceBinds(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取资源绑定列表: %s", err.Err)
		util.Respf().Fail(xresponse.HttpBadRequest).WithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(roleResourceBinds).Response(util.GinRespf(c))
}

func CreateRoleResourceBinds(c *gin.Context) {
	var roleResourceBinds []model.RoleResourceBind
	if err := c.ShouldBindJSON(&roleResourceBinds); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	for _, roleResourceBind := range roleResourceBinds {
		if xutil.IsZero(roleResourceBind.ResourceID, roleResourceBind.RoleID) {
			util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
			return
		}
		if err := svc.CreateRoleResourceBind(&roleResourceBind); err != nil {
			logrus.Errorf("创建资源绑定: %s", err.Err)
			util.Respf().FailWithError(err).Response(util.GinRespf(c))
			return
		}
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteRoleResourceBinds(c *gin.Context) {
	var roleResourceBinds []model.RoleResourceBind
	if err := c.ShouldBindJSON(&roleResourceBinds); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	for _, roleResourceBind := range roleResourceBinds {
		if xutil.IsZero(roleResourceBind.ID) {
			util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
			return
		}
		if err := svc.DeleteRoleResourceBind(&roleResourceBind); err != nil {
			logrus.Errorf("删除资源绑定: %s", err.Err)
			util.Respf().FailWithError(err).Response(util.GinRespf(c))
			return
		}
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func GetUserRoleBinds(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"role_id", "user_id"}
	userRoleBinds, err := svc.GetUserRoleBinds(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取用户角色绑定: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(userRoleBinds).Response(util.GinRespf(c))
}

func CreateUserRoleBinds(c *gin.Context) {
	var userRoleBinds []model.UserRoleBind
	if err := c.ShouldBindJSON(&userRoleBinds); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	for _, userRoleBind := range userRoleBinds {
		if xutil.IsZero(userRoleBind.UserID, userRoleBind.RoleID) {
			util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
			return
		}
		if err := svc.CreateUserRoleBind(&userRoleBind); err != nil {
			logrus.Errorf("创建用户角色绑定: %s", err.Err)
			util.Respf().FailWithError(err).Response(util.GinRespf(c))
			return
		}
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteUserRoleBinds(c *gin.Context) {
	var userRoleBinds []model.UserRoleBind
	if err := c.ShouldBindJSON(&userRoleBinds); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	for _, userRoleBind := range userRoleBinds {
		if xutil.IsZero(userRoleBind.ID) {
			util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
			return
		}
		if err := svc.DeleteUserRoleBind(&userRoleBind); err != nil {
			logrus.Errorf("删除用户角色绑定: %s", err.Err)
			util.Respf().FailWithError(err).Response(util.GinRespf(c))
			return
		}
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func GetGroupRoleBinds(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"role_id", "user_id"}
	groupRoleBinds, err := svc.GetGroupRoleBinds(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取用户组角色绑定: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(groupRoleBinds).Response(util.GinRespf(c))
}

func CreateGroupRoleBinds(c *gin.Context) {
	var groupRoleBinds []model.GroupRoleBind
	if err := c.ShouldBindJSON(&groupRoleBinds); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	for _, groupRoleBind := range groupRoleBinds {
		if xutil.IsZero(groupRoleBind.GroupID, groupRoleBind.RoleID) {
			util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
			return
		}
		if err := svc.CreateGroupRoleBind(&groupRoleBind); err != nil {
			logrus.Errorf("创建用户组资源绑定: %s", err.Err)
			util.Respf().FailWithError(err).Response(util.GinRespf(c))
			return
		}
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteGroupRoleBinds(c *gin.Context) {
	var groupRoleBinds []model.GroupRoleBind
	if err := c.ShouldBindJSON(&groupRoleBinds); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	for _, groupRoleBind := range groupRoleBinds {
		if xutil.IsZero(groupRoleBind.ID) {
			util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
			return
		}
		if err := svc.DeleteGroupRoleBind(&groupRoleBind); err != nil {
			logrus.Errorf("删除用户组资源绑定: %s", err.Err)
			util.Respf().FailWithError(err).Response(util.GinRespf(c))
			return
		}
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
