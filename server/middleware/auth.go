package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xjwt"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastorage"
	"wukong/pkg/util"
	"wukong/server/model"
	"wukong/server/service"
)

var svc = service.NewService()

func AccessControl() gin.HandlerFunc {
	// 1. 校验 token 是否有效或过期 2. 校验 token 是否有权限访问资源
	return func(c *gin.Context) {
		var token model.Token
		var reqToken string
		reqTokenHeader := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")
		if len(reqTokenHeader) == 2 {
			reqToken = reqTokenHeader[1]
		}
		tokenValid := func() util.RespError {
			if reqToken == "" {
				return util.NewRespError(errors.New("未携带令牌"))
			}
			_, err := xjwt.ParseToken(reqToken, config.Conf.Secret.TokenID, &token)
			if err != nil {
				return util.NewRespError(err)
			}
			if token.ExpiresAt != nil {
				if time.Now().UTC().Unix() > *token.ExpiresAt {
					return util.NewRespError(errors.New("令牌已过期"))
				}
			}
			if !config.Conf.DisableVerifyServerToken && !datastorage.DS.IsTokenExist(token.Username, token.TokenType, reqToken) {
				return util.NewRespError(errors.New("服务器令牌已过期"))
			}
			if !svc.IsAccessResource(&token, c) {
				return util.NewRespError(errors.New("无权访问"), true).WithCode(xresponse.HttpForbidden)
			}
			return nil
		}

		err := tokenValid()
		if err != nil {
			logrus.Infof("身份验证错误: %s", err.Err.Error())
			util.Respf().FailWithError(err, xresponse.HttpUnauthorized).Response(util.GinRespf(c))
			c.Abort()
			return
		}
		c.Set("userToken", &token)
	}
}
