package model

import (
	"time"
)

type OauthApp struct {
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
