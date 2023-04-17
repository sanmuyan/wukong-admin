package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"wukong/pkg/response"
	"wukong/server/model"
)

func GetResources(c *gin.Context) {
	likeKeys := ""
	mustKeys := []string{"resource_path"}
	resources, err := svc.GetResources(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取API资源: %s", err.Err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().WithData(resources).SetGin(c)
}

func CreateResource(c *gin.Context) {
	var resource model.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		respf().Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	if err := svc.CreateResource(&resource); err != nil {
		logrus.Errorf("创建API资源: %s", err.Err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().SetGin(c)
}

func UpdateResource(c *gin.Context) {
	var resource model.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		respf().Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	if err := svc.UpdateResource(&resource); err != nil {
		logrus.Errorf("更新API资源: %s", err.Err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().SetGin(c)
}

func DeleteResource(c *gin.Context) {
	var resource model.Resource
	if err := c.ShouldBindJSON(&resource); err != nil {
		respf().Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	if err := svc.DeleteResource(&resource); err != nil {
		logrus.Errorf("删除API资源: %s", err.Err)
		respf().Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	respf().Ok().SetGin(c)
}
