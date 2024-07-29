package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xutil"
	"sort"
	"strings"
	"wukong/pkg/db"
	"wukong/pkg/util"
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

func (s *Service) GetUserRoles(userID int) []model.Role {
	var roles []model.Role
	db.DB.Distinct().Joins("LEFT JOIN user_role_binds ON roles.id = user_role_binds.role_id").
		Joins("LEFT JOIN group_role_binds ON roles.id = group_role_binds.role_id").
		Joins("LEFT JOIN user_group_binds ON group_role_binds.group_id = user_group_binds.group_id").
		Where("user_role_binds.user_id = ? OR user_group_binds.user_id = ?", userID, userID).Find(&roles)
	return roles
}

func (s *Service) GetUserResources(userID int) []model.Resource {
	var resources []model.Resource
	db.DB.Distinct().Select("resource_path").
		Joins("JOIN role_resource_binds ON resources.id = role_resource_binds.resource_id").
		Joins("JOIN roles ON role_resource_binds.role_id = roles.id").
		Joins("LEFT JOIN user_role_binds ON roles.id = user_role_binds.role_id").
		Joins("LEFT JOIN group_role_binds ON roles.id = group_role_binds.role_id").
		Joins("LEFT JOIN user_group_binds ON group_role_binds.group_id = user_group_binds.group_id").
		Where("user_role_binds.user_id = ? OR user_group_binds.user_id = ?", userID, userID).Find(&resources)
	return resources
}

func (s *Service) IsUserResources(userID int, resourcePath string) bool {
	if tx := db.DB.Limit(1).Select("resource_path").
		Joins("JOIN role_resource_binds ON resources.id = role_resource_binds.resource_id").
		Joins("JOIN roles ON role_resource_binds.role_id = roles.id").
		Joins("LEFT JOIN user_role_binds ON roles.id = user_role_binds.role_id").
		Joins("LEFT JOIN group_role_binds ON roles.id = group_role_binds.role_id").
		Joins("LEFT JOIN user_group_binds ON group_role_binds.group_id = user_group_binds.group_id").
		Where("(user_role_binds.user_id = ? OR user_group_binds.user_id = ?) AND resources.resource_path = ?", userID, userID, resourcePath).First(&model.Resource{}); tx.RowsAffected > 0 {
		return true
	}
	return false
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

	//resources := s.GetUserResources(token.GetUserID())
	//for _, resource := range resources {
	//	// 判断客户端请求的 resource 是否等于 role 列表中的 resource
	//	// 支持父路径匹配 比如资源中有resourcePath=/api 客户端请求resourcePath=/api/user, 那么认为是有权限的
	//	rr := strings.ToLower(resource.ResourcePath)
	//	rp := strings.ToLower(routePath)
	//	// 前缀匹配
	//	//if strings.HasPrefix(rp, rr) {
	//	//	return true
	//	//}
	//	if rr == rp {
	//		return true
	//	}
	//	// 子路径匹配
	//	if xutil.IsSubPath(rr, rp) {
	//		return true
	//	}
	//}
	return s.IsUserResources(token.GetUserID(), routePath)
}

func (s *Service) GetResources(query *model.Query) (*model.Resources, util.RespError) {
	var resources model.Resources
	err := queryData(query, &resources)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &resources, nil
}

func (s *Service) CreateResource(resource *model.Resource) util.RespError {
	if err := db.DB.Create(&resource).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) UpdateResource(resource *model.Resource) util.RespError {
	if err := db.DB.Updates(&resource).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteResource(resource *model.Resource) util.RespError {
	if err := db.DB.Delete(&model.Resource{}, resource.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) GetRoles(query *model.Query) (*model.Roles, util.RespError) {
	var roles model.Roles
	err := queryData(query, &roles)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &roles, nil
}

func (s *Service) CreateRole(role *model.Role) util.RespError {
	if err := db.DB.Create(&role).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) UpdateRole(role *model.Role) util.RespError {
	if err := db.DB.Updates(&role).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteRole(role *model.Role) util.RespError {
	if err := db.DB.Delete(&model.Role{}, role.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) GetRoleResourceBinds(query *model.Query) (*model.RoleResourceBinds, util.RespError) {
	var roleResourceBinds model.RoleResourceBinds
	err := queryData(query, &roleResourceBinds)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &roleResourceBinds, nil
}

func (s *Service) CreateRoleResourceBind(roleResourceBind *model.RoleResourceBind) util.RespError {
	if err := db.DB.Create(&roleResourceBind).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteRoleResourceBind(roleResourceBind *model.RoleResourceBind) util.RespError {
	if err := db.DB.Delete(&model.RoleResourceBind{}, roleResourceBind.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) GetUserRoleBinds(query *model.Query) (*model.UserRoleBinds, util.RespError) {
	var userRoleBinds model.UserRoleBinds
	err := queryData(query, &userRoleBinds)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &userRoleBinds, nil
}

func (s *Service) CreateUserRoleBind(userRoleBind *model.UserRoleBind) util.RespError {
	if err := db.DB.Create(&userRoleBind).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteUserRoleBind(userRoleBind *model.UserRoleBind) util.RespError {
	if err := db.DB.Delete(&model.UserRoleBind{}, userRoleBind.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) GetGroupRoleBinds(query *model.Query) (*model.GroupRoleBinds, util.RespError) {
	var groupRoleBinds model.GroupRoleBinds
	err := queryData(query, &groupRoleBinds)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &groupRoleBinds, nil
}

func (s *Service) CreateGroupRoleBind(groupRoleBind *model.GroupRoleBind) util.RespError {
	if err := db.DB.Create(&groupRoleBind).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteGroupRoleBind(groupRRoleBind *model.GroupRoleBind) util.RespError {
	if err := db.DB.Delete(&model.UserRoleBind{}, groupRRoleBind.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}
