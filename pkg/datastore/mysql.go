package datastore

import (
	"wukong/pkg/db"
	"wukong/server/model"
)

type MySQLStore struct {
}

func NewMySQLStore() *MySQLStore {
	return &MySQLStore{}
}

func (c *MySQLStore) StoreToken(ts *model.StoreToken) error {
	return db.DB.Create(ts).Error
}

func (c *MySQLStore) DeleteToken(ts *model.StoreToken) error {
	return db.DB.Where("uuid = ?", ts.UUID).Delete(&model.StoreToken{}).Error
}

func (c *MySQLStore) IsTokenExist(ts *model.StoreToken) bool {
	tx := db.DB.Select("uuid").Where("uuid = ?", ts.UUID).First(&model.StoreToken{})
	if tx.RowsAffected == 0 {
		return false
	}
	return true
}

func (c *MySQLStore) StoreCode(code *model.OauthCode) error {
	return db.DB.Create(code).Error
}

func (c *MySQLStore) LoadCode(code, clientID string) (*model.OauthCode, error) {
	var codeModel model.OauthCode
	err := db.DB.Where("code = ? AND client_id = ?", code, clientID).First(&codeModel).Error
	return &codeModel, err
}

func (c *MySQLStore) DeleteCode(code, clientID string) error {
	return db.DB.Where("code = ? AND client_id = ?", code, clientID).Delete(&model.OauthCode{}).Error
}
