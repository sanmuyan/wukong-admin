package service

import (
	"errors"
	"fmt"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastore"
	"wukong/pkg/db"
	"wukong/pkg/oauthlogin"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) OauthLogin(provider string) (any, util.RespError) {
	_opConf, ok := oauthlogin.OauthProviderConfig.Load(provider)
	if !ok {
		return nil, util.NewRespError(errors.New("登录不支持"), true).WithCode(xresponse.HttpBadRequest)
	}
	opConf := _opConf.(*config.OauthProvider)
	if opConf == nil {
		return nil, util.NewRespError(errors.New("登录不支持"), true).WithCode(xresponse.HttpBadRequest)
	}
	if !opConf.Enable {
		return nil, util.NewRespError(errors.New("登录未开启"), true)
	}
	state := util.GetRandomID()
	err := datastore.DS.StoreSession(model.NewSession(state, model.SessionTypeOauthLogin, 0, "", provider).SetTimeout(config.OauthLoginTimeoutMin * time.Minute))
	if err != nil {
		return nil, util.NewRespError(err)
	}
	data := make(map[string]interface{})
	data["auth_url"] = oauthlogin.OauthUserProviders[provider].GetAuthURL(*opConf, state)
	return data, nil
}

func (s *Service) OauthCallback(code string, state string) (res *model.LoginResponse, re util.RespError) {
	session, ok := datastore.DS.LoadSession(state, model.SessionTypeOauthLogin, nil)
	if !ok {
		return nil, util.NewRespError(errors.New("登录超时"), true)
	}
	defer func() {
		// session 只能使用一次
		_ = datastore.DS.DeleteSession(state, model.SessionTypeOauthLogin)
	}()
	_opConf, ok := oauthlogin.OauthProviderConfig.Load(session.SessionRaw)
	if !ok {
		return nil, util.NewRespError(errors.New("provider not found"))
	}
	opConf := _opConf.(*config.OauthProvider)
	if opConf == nil {
		return nil, util.NewRespError(errors.New("登录不支持"), true).WithCode(xresponse.HttpBadRequest)
	}
	oUser, err := oauthlogin.OauthUserProviders[session.SessionRaw].GetUser(*opConf, code)
	if err != nil {
		return nil, util.NewRespError(err).WithCode(xresponse.HttpUnauthorized)
	}
	if oUser.GetUserID() == "" {
		return nil, util.NewRespError(errors.New("user id not found"))
	}
	var user model.User
	tx := db.DB.Where(fmt.Sprintf("%s = ?", oUser.GetUserIDField()), oUser.GetUserID()).Select("id,username,is_active").Where(&model.User{Username: user.Username}).First(&user)
	if tx.RowsAffected == 0 {
		return nil, util.NewRespError(errors.New("未绑定登录"), true).WithCode(xresponse.HttpBadRequest)
	}
	if user.IsActive != 1 {
		return nil, util.NewRespError(errors.New("用户已禁用"), true).WithCode(xresponse.HttpUnauthorized)
	}
	return s.mfaLogin(&user)
}

func (s *Service) OauthBindCallback(token *model.Token, code string, state string) util.RespError {
	session, ok := datastore.DS.LoadSession(state, model.SessionTypeOauthLogin, nil)
	if !ok {
		return util.NewRespError(errors.New("绑定超时"), true)
	}
	defer func() {
		// session 只能使用一次
		_ = datastore.DS.DeleteSession(state, model.SessionTypeOauthLogin)
	}()
	_opConf, ok := oauthlogin.OauthProviderConfig.Load(session.SessionRaw)
	if !ok {
		return util.NewRespError(errors.New("provider not found"))
	}
	opConf := _opConf.(*config.OauthProvider)
	if opConf == nil {
		return util.NewRespError(errors.New("登录不支持"), true).WithCode(xresponse.HttpBadRequest)
	}
	oUser, err := oauthlogin.OauthUserProviders[session.SessionRaw].GetUser(*opConf, code)
	if err != nil {
		return util.NewRespError(err).WithCode(xresponse.HttpUnauthorized)
	}
	if oUser.GetUserID() == "" {
		return util.NewRespError(errors.New("user id not found"))
	}
	var user model.User
	switch oUser.GetUserIDField() {
	case model.UserGitlabIDField:
		tx := db.DB.Where(fmt.Sprintf("%s = ?", model.UserGitlabIDField), oUser.GetUserID()).Select("id").First(&model.User{})
		if tx.RowsAffected > 0 {
			return util.NewRespError(errors.New("不允许多次绑定"), true).WithCode(xresponse.HttpBadRequest)
		}
		user.GitlabID = xutil.PtrTo(oUser.GetUserID())
	case model.UserWecomIDField:
		tx := db.DB.Where(fmt.Sprintf("%s = ?", model.UserWecomIDField), oUser.GetUserID()).Select("id").First(&model.User{})
		if tx.RowsAffected > 0 {
			return util.NewRespError(errors.New("不允许多次绑定"), true).WithCode(xresponse.HttpBadRequest)
		}
		user.WecomID = xutil.PtrTo(oUser.GetUserID())
	case model.UserDingtalkIDField:
		tx := db.DB.Where(fmt.Sprintf("%s = ?", model.UserDingtalkIDField), oUser.GetUserID()).Select("id").First(&model.User{})
		if tx.RowsAffected > 0 {
			return util.NewRespError(errors.New("不允许多次绑定"), true).WithCode(xresponse.HttpBadRequest)
		}
		user.DingtalkID = xutil.PtrTo(oUser.GetUserID())
	default:
		return util.NewRespError(errors.New("不支持的绑定"), true).WithCode(xresponse.HttpBadRequest)
	}
	err = db.DB.Where("id = ?", token.GetUserID()).Updates(&user).Error
	if err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) GetOauthBindStatus(token *model.Token) (*model.OauthBindStatusResponse, util.RespError) {
	var user model.User
	err := db.DB.Where("id = ?", token.GetUserID()).Select("gitlab_id,wecom_id,dingtalk_id").First(&user).Error
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &model.OauthBindStatusResponse{
		GitLab:   user.GitlabID != nil,
		Wecom:    user.WecomID != nil,
		Dingtalk: user.DingtalkID != nil}, nil

}

func (s *Service) DeleteOauthBind(token *model.Token, provider string) util.RespError {
	var column string
	switch provider {
	case model.OauthProviderGitLab:
		column = model.UserGitlabIDField
	case model.OauthProviderWecom:
		column = model.UserWecomIDField
	case model.OauthProviderDingtalk:
		column = model.UserDingtalkIDField
	default:
		return util.NewRespError(errors.New("不支持的绑定"), true).WithCode(xresponse.HttpBadRequest)
	}
	err := db.DB.Model(&model.User{}).Where("id = ?", token.GetUserID()).Update(column, nil).Error
	if err != nil {
		return util.NewRespError(err)
	}
	return nil
}
