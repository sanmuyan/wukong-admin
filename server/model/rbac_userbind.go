package model

type UserBind struct {
	Id     int `json:"id"`
	UserId int `json:"user_id" binding:"required"`
	RoleId int `json:"role_id" binding:"required"`
}

func (UserBind) TableName() string {
	return "wk_rbac_user_bind"
}
