package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/tokenutil"
	"wukong/pkg/util"
	"wukong/server/service"
)

var svc = service.NewService()

func AccessControl() gin.HandlerFunc {
	// 1. 校验 token 是否有效或过期 2. 校验 token 是否有权限访问资源
	return func(c *gin.Context) {
		token, err := tokenutil.ValidToken(c)
		if err != nil {
			logrus.Infof("身份验证错误: %s", err)
			util.Respf().Fail(xresponse.HttpUnauthorized).Response(util.GinRespf(c))
			c.Abort()
			return
		}
		if !svc.IsAccessResource(token, c) {
			util.Respf().Fail(xresponse.HttpForbidden).Response(util.GinRespf(c))
			c.Abort()
			return
		}
		c.Set("userToken", token)
		c.Next()
	}
}
