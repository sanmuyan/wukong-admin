package oauthlogin

type OauthProvider interface {
	GetUserInfo(user []byte) (OauthProvider, error)
	GetUsername() string
	GetEmail() string
	GetDisplayName() string
}

var OauthProviders = make(map[string]OauthProvider)

func init() {
	OauthProviders["gitlab"] = NewGitlabUser()
}
