package oauthlogin

import (
	"encoding/json"
)

type GitlabUser struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	DisplayName string `json:"name"`
}

func NewGitlabUser() *GitlabUser {
	return &GitlabUser{}
}

func (c *GitlabUser) GetUserInfo(userRaw []byte) (OauthProvider, error) {
	user := &GitlabUser{}
	err := json.Unmarshal(userRaw, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *GitlabUser) GetUsername() string {
	return c.Username
}

func (c *GitlabUser) GetEmail() string {
	return c.Email
}

func (c *GitlabUser) GetDisplayName() string {
	return c.DisplayName
}
