package service

import (
	"errors"
	"github.com/sanmuyan/xpkg/xbcrypt"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"strings"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) UpdateProfile(user *model.User, token *model.Token) util.RespError {
	user.Username = ""
	user.ID = token.GetUserID()
	if len(user.Password) > 0 {
		if !xbcrypt.IsPasswordComplexity(user.Password, config.Conf.Security.PasswordMinLength, config.Conf.Security.PasswordComplexity) {
			return util.NewRespError(errors.New("密码格式不正确"), true).WithCode(xresponse.HttpBadRequest)
		}
		user.Password = xbcrypt.CreatePassword(user.Password)
	}
	if err := db.DB.Updates(&user).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) GetProfile(token *model.Token) (map[string]any, util.RespError) {
	var user model.User
	userId := token.GetUserID()
	// 获取权限
	var roleComments []string
	var menuNames []string
	roles := s.GetUserRoles(userId)
	maxRoleLevel := getMaxAccessLevel(roles)
	for _, role := range roles {
		roleComments = append(roleComments, role.Comment)
		roleMenus := strings.Split(role.UserMenus, ",")
		menuNames = append(menuNames, roleMenus...)
	}
	menuNames = xutil.Deduplication(menuNames)
	// 获取用户信息
	db.DB.First(&user, userId)
	userInfo := make(map[string]any)
	var resUser model.User
	resUser.Username = user.Username
	resUser.Email = user.Email
	resUser.Mobile = user.Mobile
	resUser.DisplayName = user.DisplayName
	userInfo["roles"] = roleComments
	userInfo["access_level"] = maxRoleLevel
	userInfo["menus"] = menuNames
	userInfo["user"] = resUser

	return userInfo, nil
}
