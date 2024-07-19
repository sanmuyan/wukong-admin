package oauthlogin

import (
	"encoding/json"
	"fmt"
	"wukong/pkg/config"
	"wukong/server/model"
)

// 1. 标准 OAuth2 流程

type GitlabUser struct {
	// 用户在 Gitlab 中的 ID
	UserID int `json:"id"`
}

func NewGitlabUser() *GitlabUser {
	return &GitlabUser{}
}

func (c *GitlabUser) GetUserIDField() string {
	return model.UserGitlabIDField
}

func (c *GitlabUser) GetUserID() string {
	return fmt.Sprintf("%d", c.UserID)
}

func (c *GitlabUser) GetAuthURL(conf config.OauthProvider, state string) string {
	return getOauthConfig(conf).AuthCodeURL(state)
}

func (c *GitlabUser) GetUser(conf config.OauthProvider, code string) (OauthUserProvider, error) {
	userRaw, err := getUser(conf, code)
	if err != nil {
		return nil, err
	}
	user := &GitlabUser{}
	err = json.Unmarshal(userRaw, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
