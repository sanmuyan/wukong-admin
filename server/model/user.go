package model

import "errors"

type Token struct {
	UserId      int    `json:"user_id" binding:"required"`
	Username    string `json:"username" binding:"required"`
	AccessLevel int    `json:"access_level" binding:"required"`
	TokenType   string `json:"token_type" binding:"required"`
	ExpiresTime int64  `json:"expires_time"`
}

func (t Token) Valid() error {
	err := errors.New("required is nil")
	if t.UserId == 0 {
		return err
	}
	if t.Username == "" {
		return err
	}
	if t.TokenType == "" {
		return err
	}
	if t.TokenType != "api" && t.TokenType != "session" {
		return err
	}
	if t.TokenType == "session" && t.ExpiresTime == 0 {
		return err
	}
	return nil
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
