package model

type RoleBind struct {
	Id         int `json:"id"`
	ResourceId int `json:"resource_id"`
	RoleId     int `json:"role_id"`
}

func (RoleBind) TableName() string {
	return "wk_rbac_role_bind"
}
