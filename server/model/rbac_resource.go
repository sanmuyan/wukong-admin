package model

type Resource struct {
	Id           int    `json:"id"`
	ResourcePath string `json:"resource_path"`
	IsAuth       int    `json:"is_auth"`
	Comment      string `json:"comment"`
}

func (Resource) TableName() string {
	return "wk_rbac_resource"
}
