package model

type Role struct {
	Id          int    `json:"id"`
	RoleName    string `json:"role_name"`
	AccessLevel int    `json:"access_level"`
	Comment     string `json:"comment"`
}

func (Role) TableName() string {
	return "wk_rbac_role"
}
