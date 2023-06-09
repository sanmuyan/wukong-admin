package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) IsAuth(routePath string) bool {
	var resource model.Resource
	resource.ResourcePath = routePath
	if err := dalf().Get(&resource); err != nil {
		return true
	}
	if resource.IsAuth != 1 {
		return false
	}
	return true
}

func (s *Service) GetRBAC(userId int) model.RBAC {
	var rbac model.RBAC
	var rbacResource model.RBACUserResource
	func() (user model.User) {
		user.Id = userId
		_ = dalf().Get(&user)
		if user.IsActive == 1 {
			rbac.Active = true
		}
		return user
	}()
	if !rbac.Active {
		logrus.Warning("用户已禁用")
		return rbac
	}
	getUserBinds := func() (userBinds []model.UserBind) {
		_ = dalf().Where(model.UserBind{UserId: userId}).List(&userBinds)
		return userBinds
	}
	getRoles := func(userBind model.UserBind) (role model.Role) {
		role.Id = userBind.RoleId
		_ = dalf().Get(&role)
		return role
	}
	getRoleBinds := func(userBind model.UserBind) (resourceBinds []model.RoleBind) {
		_ = dalf().Where(model.RoleBind{RoleId: userBind.RoleId}).List(&resourceBinds)
		return resourceBinds
	}
	getResource := func(resourceBind model.RoleBind) (resource model.Resource) {
		resource.Id = resourceBind.ResourceId
		_ = dalf().Get(&resource)
		return resource
	}
	for _, userBind := range getUserBinds() {
		rbac.Roles = append(rbac.Roles, getRoles(userBind))
		for _, roleBind := range getRoleBinds(userBind) {
			resource := getResource(roleBind)
			rbacResource.ResourcePath = resource.ResourcePath
			rbac.Resources = append(rbac.Resources, rbacResource)
		}
	}
	return rbac
}

func (s *Service) GetMaxAccessLevel(userId int) int {
	// 一个用户可能绑定多个 role, 取等级最高的 role 生成 token
	var accessLevels []int
	var accessLevel int
	rbac := s.GetRBAC(userId)
	for _, role := range rbac.Roles {
		accessLevel = role.AccessLevel
		accessLevels = append(accessLevels, accessLevel)
	}
	if len(accessLevels) != 0 {
		sort.Ints(accessLevels)
		accessLevel = accessLevels[len(accessLevels)-1]
	}
	return accessLevel
}

func (s *Service) IsAccessResource(token model.Token, c *gin.Context) bool {
	userId := token.UserId
	routePath := c.FullPath()
	var user model.User
	user.Id = userId
	if err := dalf().Get(&user); err != nil || user.IsActive != 1 {
		return false
	}
	// 判断是否为管理员, 管理员无需执行下面的流程
	if token.AccessLevel >= 100 {
		return true
	}
	// 判断该资源路径是否需要鉴权
	if !s.IsAuth(c.FullPath()) {
		return true
	}
	rbac := s.GetRBAC(userId)
	for _, resource := range rbac.Resources {
		// 判断客户端请求的 resource 是否等于 role 列表中的 resource
		// 支持父路径匹配 比如资源中有resourcePath=/api 客户端请求resourcePath=/api/user, 那么认为是有权限的
		if strings.Compare(strings.ToLower(resource.ResourcePath), strings.ToLower(routePath)) == 0 {
			return true
		}
		if util.IsSubPath(strings.ToLower(resource.ResourcePath), strings.ToLower(routePath)) {
			return true
		}
	}
	return false
}
