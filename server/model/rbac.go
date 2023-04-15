package model

type RBACUserResource struct {
	ResourcePath string `json:"resource_Path"`
}

type RBAC struct {
	Roles     []Role             `json:"roles"`
	Resources []RBACUserResource `json:"resources"`
	Active    bool               `json:"active"`
}
