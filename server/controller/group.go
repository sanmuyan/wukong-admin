package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func GetGroups(c *gin.Context) {
	likeKeys := "group_name,display_name"
	mustKeys := []string{"id", "group_name"}
	groups, err := svc.GetGroups(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取用户组列表: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(groups).Response(util.GinRespf(c))
}

func CreateGroup(c *gin.Context) {
	var Group model.Group
	if err := c.ShouldBindJSON(&Group); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(Group.GroupName) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.CreateGroup(&Group); err != nil {
		logrus.Errorf("创建用户组: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func UpdateGroup(c *gin.Context) {
	var Group model.Group
	if err := c.ShouldBindJSON(&Group); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(Group.ID) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.UpdateGroup(&Group); err != nil {
		logrus.Errorf("更新用户组: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteGroup(c *gin.Context) {
	var Group model.Group
	if err := c.ShouldBindJSON(&Group); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if xutil.IsZero(Group.ID) {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.DeleteGroup(&Group); err != nil {
		logrus.Errorf("删除用户组: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func GetUserGroupBinds(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"user_id", "group_id"}
	UserGroupBinds, err := svc.GetUserGroupBinds(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取用户组绑定: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().WithData(UserGroupBinds).Response(util.GinRespf(c))
}

func CreateUserGroupBinds(c *gin.Context) {
	var UserGroupBinds []model.UserGroupBind
	if err := c.ShouldBindJSON(&UserGroupBinds); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	for _, UserGroupBind := range UserGroupBinds {
		if xutil.IsZero(UserGroupBind.GroupID, UserGroupBind.UserID) {
			util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
			return
		}
		if err := svc.CreateUserGroupBind(&UserGroupBind); err != nil {
			logrus.Errorf("创建用户组绑定: %s", err.Err)
			util.Respf().FailWithError(err).Response(util.GinRespf(c))
			return
		}
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}

func DeleteUserGroupBinds(c *gin.Context) {
	var UserGroupBinds []model.UserGroupBind
	if err := c.ShouldBindJSON(&UserGroupBinds); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	for _, UserGroupBind := range UserGroupBinds {
		if xutil.IsZero(UserGroupBind.ID) {
			util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
			return
		}
		if err := svc.DeleteUserGroupBind(&UserGroupBind); err != nil {
			logrus.Errorf("删除用户组绑定: %s", err.Err)
			util.Respf().FailWithError(err).Response(util.GinRespf(c))
			return
		}
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
