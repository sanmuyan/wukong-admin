package model

import "github.com/golang-jwt/jwt"

type Token struct {
	UserId      int    `json:"user_id" binding:"required"`
	Username    string `json:"username" binding:"required"`
	AccessLevel int    `json:"access_level" binding:"required"`
	TokenType   string `json:"token_type" binding:"required"`
	TTL         int    `json:"ttl"`
	Timestamp   int64  `json:"timestamp"`
	jwt.StandardClaims
}

func TokenKeyName(userName string, tokenType string) string {
	return "wukong:token:" + tokenType + ":" + userName
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Password    string `json:"password"`
	Mobile      string `json:"mobile"`
	Email       string `json:"email"`
	Source      string `json:"source"`
	IsActive    int    `json:"is_active"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

func (User) TableName() string {
	return "wk_user"
}
