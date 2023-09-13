package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
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
		util.Respf().Fail(xresponse.HttpInternalServerError).Response(util.GinRespf(c))
	}
	util.Respf().Ok().WithData(resources).Response(util.GinRespf(c))
}

func CreateResource(c *gin.Context) {
	var resource model.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	if err := svc.CreateResource(&resource); err != nil {
		logrus.Errorf("创建API资源: %s", err.Err)
		util.Respf().Fail(xresponse.HttpInternalServerError).Response(util.GinRespf(c))
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
	if err := svc.UpdateResource(&resource); err != nil {
		logrus.Errorf("更新API资源: %s", err.Err)
		util.Respf().Fail(xresponse.HttpInternalServerError).Response(util.GinRespf(c))
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
	if err := svc.DeleteResource(&resource); err != nil {
		logrus.Errorf("删除API资源: %s", err.Err)
		util.Respf().Fail(xresponse.HttpInternalServerError).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
