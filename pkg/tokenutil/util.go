package tokenutil

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xjwt"
	"strings"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastorage"
	"wukong/server/model"
)

func ValidToken(tokenStr string) (token model.Token, err error) {
	if tokenStr == "" {
		return token, errors.New("未携带令牌")
	}
	_, err = xjwt.ParseToken(tokenStr, config.Conf.Secret.TokenID, &token)
	if err != nil {
		return token, errors.New("令牌不合法")
	}
	if token.ExpiresAt != nil {
		if time.Now().UTC().Unix() > *token.ExpiresAt {
			return token, errors.New("令牌已过期")
		}
	}
	if !config.Conf.DisableVerifyServerToken && !datastorage.DS.IsTokenExist(token.Username, token.TokenType, tokenStr) {
		return token, errors.New("服务器令牌已过期")
	}
	return token, nil
}

func ParseHeader(c *gin.Context) (token string) {
	tokenHeaderSplit := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")
	if len(tokenHeaderSplit) == 2 {
		return tokenHeaderSplit[1]
	}
	return token
}

func ParseCookie(c *gin.Context) (token string) {
	tokenCookie, _ := c.Cookie("Authorization")
	tokenCookieSplit := strings.Split(tokenCookie, "Bearer ")
	if len(tokenCookieSplit) == 2 {
		return tokenCookieSplit[1]
	}
	return token
}
