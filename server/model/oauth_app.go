package model

import (
	"time"
)

type OauthAPP struct {
	ID           int       `json:"id"`
	APPName      string    `json:"app_name"`
	ClientID     string    `json:"client_id" gorm:"<-:create"`
	ClientSecret string    `json:"client_secret" gorm:"<-:create"`
	Scope        string    `json:"scope,omitempty"`
	RedirectURI  string    `json:"redirect_uri,omitempty"`
	Comment      string    `json:"comment,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

type OauthCode struct {
	ID           int
	Code         string
	Username     string
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scope        string
	ExpiresAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type OauthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type"`
}
