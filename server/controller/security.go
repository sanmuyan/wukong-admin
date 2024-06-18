package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/util"
	"wukong/server/model"
)

func ModifyPassword(c *gin.Context) {
	var req model.ModifyPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.Respf().Fail(xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	err := svc.ModifyPassword(&req, keysToUserToken(c))
	if err != nil {
		logrus.Errorf("修改密码: %s", err.Err)
		util.Respf().FailWithError(err).Response(util.GinRespf(c))
		return
	}
	util.Respf().Ok().Response(util.GinRespf(c))
}
