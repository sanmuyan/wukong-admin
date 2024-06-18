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

func ValidToken(tokenStr string) (token model.Token, err error) {
	if tokenStr == "" {
		return token, errors.New("未携带令牌")
	}
	_, err = xjwt.ParseToken(tokenStr, config.Conf.Secret.TokenKey, &token)
	if err != nil {
		return token, errors.New("令牌不合法")
	}
	if token.ExpiresAt != nil {
		if time.Now().Unix() > *token.ExpiresAt {
			return token, errors.New("令牌已过期")
		}
	}
	if !config.Conf.DisableVerifyServerToken {
		if _, ok := datastore.DS.LoadSession(token.TokenID, token.TokenType, token.Username); !ok {
			return token, errors.New("服务器令牌已过期")
		}
	}
	var user model.User
	db.DB.Select("id,is_active").Where("username", token.Username).First(&user)
	if user.ID == 0 || user.IsActive != 1 {
		return token, errors.New("用户不存在或未激活")
	}
	token.SetUserID(user.ID)
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
