package middleware

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/pkg/response"
	"wukong/pkg/util"
	"wukong/server/model"
	"wukong/server/service"
)

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
			tokenClaims, err := util.ExtractToken(reqToken, config.Conf.Secret.TokenKey)
			if err != nil {
				return model.NewError(err.Error())
			}
			if err = json.Unmarshal(tokenClaims.Body, &token); err != nil {
				return model.NewError(err.Error())
			}
			if t := db.RDB.Get(ctx, model.TokenKeyName(token.Username, token.TokenType)).Val(); t != reqToken {
				return model.NewError("令牌已过期")
			}
			if !service.IsAccessResource(token, c) {
				return model.NewError("无权访问", true)
			}
			return nil
		}

		var resp = response.NewResponse()
		err := res()
		if err != nil {
			logrus.Infof("身份验证错误: %s", err.Err.Error())
			if err.IsResponseMsg {
				resp.FailWithMsg(response.HttpUnauthorized, err.Err.Error()).SetGin(c)
			} else {
				resp.Fail(response.HttpUnauthorized).SetGin(c)
			}
			c.Abort()
			return
		}
		c.Set("userToken", token)
	}
}
