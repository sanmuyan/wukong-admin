package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/pkg/passkey"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) GetPassKeys(token *model.Token) (any, util.RespError) {
	var passKeys []*model.PassKey
	err := db.DB.Select("id,user_id,display_name,last_used_at,created_at,updated_at").Where("user_id = ?", token.GetUserID()).Find(&passKeys).Error
	if err != nil {
		return nil, util.NewRespError(err)
	}
	for _, passKey := range passKeys {
		if passKey.LastUsedAt == nil {
			passKey.LastUsedAt = &passKey.UpdatedAt
		}
	}
	data := make(map[string][]*model.PassKey)
	data["pass_keys"] = passKeys
	return data, nil
}

func (s *Service) UpdatePassKey(passKey *model.PassKey, token *model.Token) util.RespError {
	err := db.DB.Where("user_id = ? AND id = ?", token.GetUserID(), passKey.ID).Updates(&model.PassKey{
		DisplayName: passKey.DisplayName,
	}).Error
	if err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeletePassKey(passKey *model.PassKey, token *model.Token) util.RespError {
	err := db.DB.Where("user_id = ? AND id = ?", token.GetUserID(), passKey.ID).Delete(&model.PassKey{}).Error
	if err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) PassKeyBeginRegistration(token *model.Token) (*protocol.CredentialCreation, util.RespError) {
	db.DB.Where("user_id = ?", token.GetUserID()).Delete(&model.PassKeyRegisterSession{})
	var passKeyCount int64
	db.DB.Model(&model.PassKey{}).Where("user_id = ?", token.GetUserID()).Count(&passKeyCount)
	if passKeyCount >= config.PassKeyMax {
		return nil, util.NewRespError(errors.New("超出最大数量"), true).WithCode(xresponse.HttpBadRequest)
	}
	var user model.User
	db.DB.First(&user, token.GetUserID())
	options, session, err := passkey.WebAuthn.BeginRegistration(passkey.NewWAUser(&user), func(options *protocol.PublicKeyCredentialCreationOptions) {
		options.Parameters = []protocol.CredentialParameter{
			{
				Type:      "public-key",
				Algorithm: webauthncose.AlgES256,
			},
		}
	})
	if err != nil {
		return nil, util.NewRespError(err)
	}
	db.DB.Create(&model.PassKeyRegisterSession{
		UserID:     token.GetUserID(),
		SessionRaw: string(xutil.RemoveError(json.Marshal(session))),
		ExpireAt:   time.Now().UTC().Add(config.PassKeyRegistrationTimeoutMin * time.Minute),
	})
	return options, nil
}

func (s *Service) PassKeyFinishRegistration(token *model.Token, c *gin.Context) util.RespError {
	var user model.User
	db.DB.First(&user, token.GetUserID())
	var passKeySession model.PassKeyRegisterSession
	db.DB.Where("user_id = ?", token.GetUserID()).First(&passKeySession)
	if passKeySession.ExpireAt.Before(time.Now().UTC()) {
		return util.NewRespError(errors.New("注册超时"), true).WithCode(xresponse.HttpBadRequest)
	}
	var session webauthn.SessionData
	if err := json.Unmarshal([]byte(passKeySession.SessionRaw), &session); err != nil {
		return util.NewRespError(err)
	}
	credential, err := passkey.WebAuthn.FinishRegistration(passkey.NewWAUser(&user), session, c.Request)
	if err != nil {
		return util.NewRespError(err)
	}
	if err := db.DB.Create(&model.PassKey{
		UserID:        token.GetUserID(),
		DisplayName:   "安全密钥",
		CredentialID:  base64.RawURLEncoding.EncodeToString(credential.ID),
		CredentialRaw: string(xutil.RemoveError(json.Marshal(credential))),
	}).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) BeginPassKeyLogin(req *model.PassKeyBeginLoginRequest) (*model.PassKeyBeginLoginResponse, util.RespError) {
	db.DB.Where("username = ?", req.Username).Delete(&model.PassKeyLoginSession{})
	var user model.User
	tx := db.DB.Where("username = ?", req.Username).First(&user)
	if tx.RowsAffected == 0 {
		return nil, util.NewRespError(errors.New("用户不存在"), true).WithCode(xresponse.HttpBadRequest)
	}
	options, session, err := passkey.WebAuthn.BeginLogin(passkey.NewWAUser(&user))
	if err != nil {
		return nil, util.NewRespError(err).WithCode(xresponse.HttpBadRequest)
	}
	sessionID := uuid.New().String()
	err = db.DB.Create(&model.PassKeyLoginSession{
		Username:   req.Username,
		UserID:     user.ID,
		SessionID:  sessionID,
		SessionRaw: string(xutil.RemoveError(json.Marshal(session))),
		ExpireAt:   time.Now().UTC().Add(config.PassKeyLoginTimeoutMin * time.Minute),
	}).Error
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &model.PassKeyBeginLoginResponse{
		Username:  req.Username,
		SessionID: sessionID,
		Options:   options,
	}, nil
}

func (s *Service) FinishPassKeyLogin(req *model.PassKeyFinishLoginRequest, c *gin.Context) (*model.LoginResponse, util.RespError) {
	var user model.User
	tx := db.DB.Where("username = ?", req.Username).First(&user)
	if tx.RowsAffected == 0 {
		return nil, util.NewRespError(errors.New("用户不存在"))
	}
	var passKeyLoginSession model.PassKeyLoginSession
	tx = db.DB.Where("username = ? AND session_id = ?", req.Username, req.SessionID).First(&passKeyLoginSession)
	if tx.RowsAffected == 0 {
		return nil, util.NewRespError(errors.New("未知 session"))
	}
	if passKeyLoginSession.ExpireAt.Before(time.Now().UTC()) {
		return nil, util.NewRespError(errors.New("登录超时"), true)
	}
	var session webauthn.SessionData
	if err := json.Unmarshal([]byte(passKeyLoginSession.SessionRaw), &session); err != nil {
		return nil, util.NewRespError(err)
	}
	credential, err := passkey.WebAuthn.FinishLogin(passkey.NewWAUser(&user), session, c.Request)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	db.DB.Model(&model.PassKey{}).Where("user_id = ? AND credential_id = ?", user.ID, base64.RawURLEncoding.EncodeToString(credential.ID)).Update("last_used_at", time.Now().UTC())
	return s.createLoginToken(user.ID, user.Username)
}
