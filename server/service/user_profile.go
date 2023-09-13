package service

import (
	"github.com/sanmuyan/xpkg/xbcrypt"
	"time"
	"wukong/server/model"
)

func (s *Service) UpdateUserProfile(user *model.User, token *model.Token) *model.Error {
	if token == nil {
		return model.NewError("token is nil")
	}
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	user.Username = ""
	user.Id = token.UserId
	user.UpdateTime = nowTime
	if len(user.Password) > 0 {
		if !xbcrypt.IsPasswordComplexity(user.Password, 8, true, true, true, true) {
			return model.NewError("密码格式不正确", true)
		}
		user.Password = xbcrypt.CreatePassword(user.Password)
	}
	if err := dalf().Save(&user); err != nil {
		return &model.Error{Err: err}
	}
	return nil
}

func (s *Service) GetUserProfile(token *model.Token) (map[string]any, *model.Error) {
	if token == nil {
		return nil, model.NewError("token is nil")
	}
	var user model.User
	userId := token.UserId
	// 获取权限
	rbac := s.GetRBAC(userId)

	// 获取用户信息
	user.Id = userId
	_ = dalf().Get(&user)
	userInfo := make(map[string]any)
	var resUser model.User
	resUser.Username = user.Username
	resUser.Email = user.Email
	resUser.Mobile = user.Mobile
	resUser.DisplayName = user.DisplayName
	resUser.IsActive = user.IsActive
	userInfo["rbac"] = rbac
	userInfo["user"] = resUser

	return userInfo, nil
}
