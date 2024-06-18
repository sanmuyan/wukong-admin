package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xutil"
	"sort"
	"strings"
	"wukong/pkg/db"
	"wukong/server/model"
)

func (s *Service) IsAuth(routePath string) bool {
	var resource model.Resource
	if err := db.DB.Select("is_auth").Where("resource_path = ?", routePath).First(&resource).Error; err != nil {
		return true
	}
	if resource.IsAuth != 1 {
		return false
	}
	return true
}

func (s *Service) GetUserRoles(userId int) []model.Role {
	var userRoles []model.Role
	getUserBinds := func() (userBinds []model.UserBind) {
		db.DB.Where("user_id = ?", userId).Find(&userBinds)
		return userBinds
	}
	getUserRole := func(userBind model.UserBind) (role model.Role) {
		db.DB.First(&role, userBind.RoleID)
		return role
	}
	for _, userBind := range getUserBinds() {
		userRoles = append(userRoles, getUserRole(userBind))
	}
	return userRoles
}

func (s *Service) GetUserResources(roles []model.Role) []model.Resource {
	var resources []model.Resource
	getRoleBinds := func(roleID int) (resourceBinds []model.RoleBind) {
		db.DB.Where("role_id = ?", roleID).Find(&resourceBinds)
		return resourceBinds
	}
	getRoleResource := func(resourceBind model.RoleBind) (resource model.Resource) {
		db.DB.First(&resource, resourceBind.ResourceID)
		return resource
	}
	for _, role := range roles {
		getRoleBinds(role.ID)
		for _, resourceBind := range getRoleBinds(role.ID) {
			resources = append(resources, getRoleResource(resourceBind))
		}
	}
	return resources
}

func (s *Service) GetMaxAccessLevel(roles []model.Role) int {
	// 一个用户可能绑定多个 role, 取等级最高的 role 生成 token
	var accessLevels []int
	var accessLevel int
	for _, role := range roles {
		accessLevel = role.AccessLevel
		accessLevels = append(accessLevels, accessLevel)
	}
	if len(accessLevels) != 0 {
		sort.Ints(accessLevels)
		accessLevel = accessLevels[len(accessLevels)-1]
	}
	return accessLevel
}

func (s *Service) IsAccessResource(token *model.Token, c *gin.Context) bool {
	//var user model.User
	//if err := db.DB.Where(&model.User{Username: token.Username}).First(&user).Error; err != nil {
	//	return false
	//}
	//if user.IsActive != 1 {
	//	return false
	//}
	//token.SetUserID(user.ID)
	routePath := c.FullPath()

	// 处理 Oauth refreshToken
	if token.TokenType == model.TokenTypeOauthRefresh {
		if routePath == "/api/oauth/token" {
			return true
		}
		return false
	}

	// 处理 OAuth accessToken
	if token.TokenType == model.TokenTypeOauthAccess {
		scope := strings.Split(token.Scope, ",")
		if xutil.IsContains("profile", scope) && routePath == "/api/account/profile" {
			return true
		}
		if !xutil.IsContains("api", scope) {
			return false
		}
	}

	// 判断是否为管理员, 管理员无需执行下面的流程
	if token.AccessLevel >= 100 {
		return true
	}
	// 判断该资源路径是否需要鉴权
	if !s.IsAuth(c.FullPath()) {
		return true
	}

	resources := s.GetUserResources(s.GetUserRoles(token.GetUserID()))
	for _, resource := range resources {
		// 判断客户端请求的 resource 是否等于 role 列表中的 resource
		// 支持父路径匹配 比如资源中有resourcePath=/api 客户端请求resourcePath=/api/user, 那么认为是有权限的
		rr := strings.ToLower(resource.ResourcePath)
		rp := strings.ToLower(routePath)
		// 前缀匹配
		//if strings.HasPrefix(rp, rr) {
		//	return true
		//}
		if rr == rp {
			return true
		}
		// 子路径匹配
		if xutil.IsSubPath(rr, rp) {
			return true
		}
	}
	return false
}
