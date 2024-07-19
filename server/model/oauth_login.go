package model

// OauthBindStatusResponse 响应对象
type OauthBindStatusResponse struct {
	GitLab   bool `json:"gitlab"`
	Wecom    bool `json:"wecom"`
	Dingtalk bool `json:"dingtalk"`
}

const (
	OauthProviderGitLab   = "gitlab"
	OauthProviderWecom    = "wecom"
	OauthProviderDingtalk = "dingtalk"
)
