package tokenutil

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xjwt"
	"strings"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastore"
	"wukong/pkg/db"
	"wukong/server/model"
)

// ValidToken 验证令牌
func validToken(tokenStr string) (*model.Token, error) {
	var token model.Token
	if tokenStr == "" {
		return nil, errors.New("未携带令牌")
	}
	_, err := xjwt.ParseToken(tokenStr, config.Conf.Secret.TokenKey, &token)
	if err != nil {
		return nil, errors.New("令牌不合法")
	}
	if token.ExpiresAt != nil {
		if time.Now().Unix() > *token.ExpiresAt {
			return nil, errors.New("令牌已过期")
		}
	}
	if !config.Conf.Security.DisableVerifyServerToken {
		if _, ok := datastore.DS.LoadSession(token.TokenID, token.TokenType, nil, token.Username); !ok {
			return nil, errors.New("服务器令牌已过期")
		}
	}
	var user model.User
	db.DB.Select("id,is_active").Where("username", token.Username).First(&user)
	if user.IsActive != 1 {
		return nil, errors.New("用户不存在或未激活")
	}
	token.SetUserID(user.ID)
	return &token, nil
}

// getTokenFromHeader 从请求头中获取令牌
func getTokenFromHeader(c *gin.Context) (string, bool) {
	tokenHeaderSplit := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")
	if len(tokenHeaderSplit) == 2 {
		return tokenHeaderSplit[1], true
	}
	return "", false
}

// getTokenFromCookie 从 Cookie 中获取令牌
func getTokenFromCookie(c *gin.Context) (string, bool) {
	tokenCookie, _ := c.Cookie("Authorization")
	tokenCookieSplit := strings.Split(tokenCookie, "Bearer ")
	if len(tokenCookieSplit) == 2 {
		return tokenCookieSplit[1], true
	}
	return "", false
}

// getToken 从请求头或 Cookie 中获取令牌
func getToken(c *gin.Context) (token string, ok bool) {
	token, ok = getTokenFromHeader(c)
	if ok {
		return token, ok
	}
	token, ok = getTokenFromCookie(c)
	if ok {
		return token, ok
	}
	return token, false
}

// ValidToken 验证令牌
func ValidToken(c *gin.Context) (token *model.Token, err error) {
	tokenStr, _ := getToken(c)
	return validToken(tokenStr)
}

// ValidTokenStr 验证字符串令牌
func ValidTokenStr(tokenStr string) (token *model.Token, err error) {
	return validToken(tokenStr)
}
