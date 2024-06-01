package tokenclient

import (
	"github.com/sanmuyan/xpkg/xutil"
	"time"
	"wukong/pkg/db"
)

type Token struct {
	TokenKey  string
	TokenType string
	Token     string
	ExpiresAt *time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type MySQLTokenClient struct {
}

func NewMySQLTokenClient() *MySQLTokenClient {
	return &MySQLTokenClient{}
}

func (c *MySQLTokenClient) SetToken(tokenKey string, tokenType string, tokenStr string, e time.Duration) error {
	token := Token{
		TokenKey:  tokenKey,
		TokenType: tokenType,
	}
	var expiresAt *time.Time
	if e > 0 {
		expiresAt = xutil.PtrTo[time.Time](time.Now().UTC().Add(e))
	}
	tx := db.DB.Select("token").Where(&token).First(&Token{})
	if tx.RowsAffected > 0 {
		return db.DB.Where(&token).Updates(&Token{Token: tokenStr, ExpiresAt: expiresAt}).Error
	}
	return db.DB.Create(&Token{TokenKey: tokenKey, TokenType: tokenType, Token: tokenStr, ExpiresAt: expiresAt}).Error
}

func (c *MySQLTokenClient) DeleteToken(tokenKey string, tokenType string) error {
	return db.DB.Where(&Token{TokenKey: tokenKey, TokenType: tokenType}).Delete(&Token{}).Error
}

func (c *MySQLTokenClient) IsTokenExist(tokenKey string, tokenType string, tokenStr string) bool {
	token := Token{
		TokenKey:  tokenKey,
		TokenType: tokenType,
	}
	tx := db.DB.Where(&token).First(&token)
	if tx.RowsAffected == 0 {
		return false
	}
	if token.ExpiresAt != nil {
		if time.Now().UTC().After(*token.ExpiresAt) {
			_ = c.DeleteToken(tokenKey, tokenType)
			return false
		}
	}
	if token.Token != tokenStr {
		return false
	}
	return true
}
