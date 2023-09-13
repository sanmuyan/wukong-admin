package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xjwt"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
	"wukong/server/service"
)

var svc = service.NewService()

func AccessControl() gin.HandlerFunc {
	// 1. 校验 token 是否有效或过期 2. 校验 token 是否有权限访问资源
	return func(c *gin.Context) {
		var token model.Token
		var ctx = context.Background()
		reqToken := c.Request.Header.Get("token")
		res := func() *model.Error {
			if reqToken == "" {
				return model.NewError("未携带令牌")
			}
			_, err := xjwt.ParseToken(reqToken, config.Conf.Secret.TokenKey, &token)
			if err != nil {
				return model.NewError(err.Error())
			}
			if token.ExpiresTime > 0 && token.ExpiresTime < time.Now().UTC().Unix() || token.UserId == 0 {
				return model.NewError("令牌已过期")
			}
			if e := db.RDB.Exists(ctx, model.TokenKeyName(token.Username, token.TokenType)).Val(); e == 0 {
				return model.NewError("令牌已过期")
			}
			if !svc.IsAccessResource(token, c) {
				return model.NewError("无权访问", true)
			}
			return nil
		}

		err := res()
		if err != nil {
			logrus.Infof("身份验证错误: %s", err.Err.Error())
			if err.IsResponseMsg {
				util.Respf().Fail(xresponse.HttpForbidden).WithMsg(err.Err.Error()).Response(util.GinRespf(c))
			} else {
				util.Respf().Fail(xresponse.HttpUnauthorized).Response(util.GinRespf(c))
			}
			c.Abort()
			return
		}
		c.Set("userToken", token)
	}
}
