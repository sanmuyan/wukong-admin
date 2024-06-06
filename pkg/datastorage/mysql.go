package datastorage

import (
	"github.com/sanmuyan/xpkg/xutil"
	"time"
	"wukong/pkg/db"
	"wukong/server/model"
)

type Token struct {
	TokenID   string `gorm:"<-:create"`
	TokenType string `gorm:"<-:create"`
	Token     string
	ExpiresAt *time.Time
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}

type MySQLStorage struct {
}

func NewMySQLStorage() *MySQLStorage {
	return &MySQLStorage{}
}

func (c *MySQLStorage) StoreToken(tokenID string, tokenType string, tokenStr string, e time.Duration) error {
	token := Token{
		TokenID:   tokenID,
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
	return db.DB.Create(&Token{TokenID: tokenID, TokenType: tokenType, Token: tokenStr, ExpiresAt: expiresAt}).Error
}

func (c *MySQLStorage) DeleteToken(tokenID string, tokenType string) error {
	return db.DB.Where(&Token{TokenID: tokenID, TokenType: tokenType}).Delete(&Token{}).Error
}

func (c *MySQLStorage) IsTokenExist(tokenID string, tokenType string, tokenStr string) bool {
	token := Token{
		TokenID:   tokenID,
		TokenType: tokenType,
	}
	tx := db.DB.Where(&token).First(&token)
	if tx.RowsAffected == 0 {
		return false
	}
	if token.ExpiresAt != nil {
		if time.Now().UTC().After(*token.ExpiresAt) {
			_ = c.DeleteToken(tokenID, tokenType)
			return false
		}
	}
	if token.Token != tokenStr {
		return false
	}
	return true
}

func (c *MySQLStorage) StoreCode(code *model.OauthCode) error {
	return db.DB.Create(code).Error
}

func (c *MySQLStorage) LoadCode(code, clientID string) (*model.OauthCode, error) {
	var codeModel model.OauthCode
	err := db.DB.Where("code = ? AND client_id = ?", code, clientID).First(&codeModel).Error
	return &codeModel, err
}

func (c *MySQLStorage) DeleteCode(code, clientID string) error {
	return db.DB.Where("code = ? AND client_id = ?", code, clientID).Delete(&model.OauthCode{}).Error
}
