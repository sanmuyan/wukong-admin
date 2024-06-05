package useroauth

type OAuthProvider interface {
	GetUserInfo(user []byte) (OAuthProvider, error)
	GetUsername() string
	GetEmail() string
	GetDisplayName() string
}

var OAuthProviders = make(map[string]OAuthProvider)

func init() {
	OAuthProviders["gitlab"] = NewGitlabUser()
}
