package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"wukong/pkg/response"
	"wukong/server/model"
)

func GetUsers(c *gin.Context) {
	resp := response.NewResponse()
	likeKeys := "username,email,mobile,display_name"
	mustKeys := []string{"id", "username"}
	users, err := svc.GetUsers(getQuery(c, likeKeys, mustKeys))
	if err != nil {
		logrus.Errorf("获取用户列表: %s", err.Err)
		resp.Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	resp.OkWithData(users).SetGin(c)
}

func CreateUser(c *gin.Context) {
	resp := response.NewResponse()
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		resp.Fail(response.HttpBadRequest).SetGin(c)
		fmt.Println(err)
		return
	}
	if err := svc.CreateUser(&user); err != nil {
		logrus.Errorf("创建用户: %s", err.Err)
		if err.IsResponseMsg {
			resp.FailWithMsg(response.HttpInternalServerError, err.Err.Error()).SetGin(c)
		} else {
			resp.Fail(response.HttpInternalServerError).SetGin(c)
		}
		return
	}
	resp.Ok().SetGin(c)
}

func UpdateUser(c *gin.Context) {
	resp := response.NewResponse()
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		resp.Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	if err := svc.UpdateUser(&user); err != nil {
		logrus.Errorf("更新用户: %s", err.Err)
		if err.IsResponseMsg {
			resp.FailWithMsg(response.HttpInternalServerError, err.Err.Error()).SetGin(c)
		} else {
			resp.Fail(response.HttpInternalServerError).SetGin(c)
		}
		return
	}
	resp.Ok().SetGin(c)
}

func DeleteUser(c *gin.Context) {
	resp := response.NewResponse()
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		resp.Fail(response.HttpBadRequest).SetGin(c)
		return
	}
	if err := svc.DeleteUser(&user); err != nil {
		logrus.Errorf("删除用户: %s", err.Err)
		resp.Fail(response.HttpInternalServerError).SetGin(c)
		return
	}
	resp.Ok().SetGin(c)
}
