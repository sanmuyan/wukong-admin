package passkey

import (
	"encoding/json"
	"fmt"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/sirupsen/logrus"
	"wukong/pkg/db"
	"wukong/server/model"
)

type WAUser struct {
	*model.User
}

func NewWAUser(user *model.User) *WAUser {
	return &WAUser{User: user}
}

func (c *WAUser) WebAuthnID() []byte {
	return []byte(fmt.Sprintf("%d", c.ID))
}

func (c *WAUser) WebAuthnName() string {
	return c.Username
}

func (c *WAUser) WebAuthnDisplayName() string {
	return c.DisplayName
}

func (c *WAUser) WebAuthnCredentials() []webauthn.Credential {
	var credentials []webauthn.Credential
	var passKeys []model.PassKey
	db.DB.Where("user_id = ?", c.ID).Find(&passKeys)
	for _, passKey := range passKeys {
		var credential webauthn.Credential
		err := json.Unmarshal([]byte(passKey.CredentialRaw), &credential)
		if err != nil {
			logrus.Warnf("unmarshal credential error: %s", err)
			continue
		}
		credentials = append(credentials, credential)
	}
	return credentials
}

func (c *WAUser) WebAuthnIcon() string {
	return ""
}
