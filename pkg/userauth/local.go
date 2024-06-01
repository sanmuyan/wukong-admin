package userauth

import (
	"github.com/sanmuyan/xpkg/xbcrypt"
	"wukong/pkg/db"
	"wukong/server/model"
)

type LocalAuth struct {
}

func NewLocalAuth() *LocalAuth {
	return &LocalAuth{}
}

func (l *LocalAuth) Login(username string, password string) bool {
	var user model.User
	user.Username = username
	if err := db.DB.Select("password").Where(&model.User{Username: user.Username}).First(&user).Error; err != nil {
		return false
	}
	return xbcrypt.ComparePassword(user.Password, password)
}
