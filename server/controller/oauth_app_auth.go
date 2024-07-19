package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sanmuyan/xpkg/xresponse"
	"wukong/pkg/config"
	"wukong/pkg/tokenutil"
	"wukong/pkg/util"
	"wukong/server/model"
)

// GetOauthCodeSessionFrontRedirect 获取授权码，并返回回调重定向地址，由浏览器执行重定向
func GetOauthCodeSessionFrontRedirect(c *gin.Context) {
	var req model.OauthCodeSessionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		util.Respf().FailWithError(util.NewRespError(errors.New("invalid_request"), true), xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	redirectURI, _err := svc.GetOauthCodeSession(keysToUserToken(c), &req)
	if _err != nil {
		util.Respf().FailWithError(util.NewRespError(errors.New(_err.Error), true), xresponse.HttpBadRequest).Response(util.GinRespf(c))
		return
	}
	data := make(map[string]interface{})
	data["redirect_uri"] = redirectURI
	util.Respf().Ok().WithData(data).Response(util.GinRespf(c))
}

// GetOauthCodeSession 获取授权码，并重定向回调地址，如果未登录则重定向到登录入口
func GetOauthCodeSession(c *gin.Context) {
	var req model.OauthCodeSessionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(200, model.NewOauthErrorResponse("invalid_request"))
		return
	}
	token, err := tokenutil.ValidToken(c)
	if err != nil {
		siteURL := config.Conf.Basic.SiteURL
		c.Redirect(302, fmt.Sprintf("%s/login?redirect_uri=%s%s", siteURL, siteURL, c.Request.URL.String()))
		return
	}
	redirectURI, _err := svc.GetOauthCodeSession(token, &req)
	if _err != nil {
		c.JSON(200, _err)
		return
	}
	c.Redirect(302, redirectURI)
}

// GetOauthToken 获取 OAuth token
func GetOauthToken(c *gin.Context) {
	var req model.OauthTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, model.NewOauthErrorResponse("invalid_request"))
		return
	}
	switch req.GrantType {
	case "authorization_code":
		// 获取 access token & refresh token
		token, err := svc.GetOauthToken(&req)
		if err != nil {
			c.JSON(200, err)
			return
		}
		c.JSON(200, token)
	case "refresh_token":
		// 刷新 access token
		token, err := svc.RefreshOauthToken(&req)
		if err != nil {
			c.JSON(200, err)
			return
		}
		c.JSON(200, token)
	default:
		c.JSON(200, model.NewOauthErrorResponse("invalid_grant_type"))
	}
}

// RevokeOauthToken 撤销 OAuth token
func RevokeOauthToken(c *gin.Context) {
	var req model.OauthRevokeTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, model.NewOauthErrorResponse("invalid_request"))
		return
	}
	err := svc.RevokeOauthToken(&req)
	if err != nil {
		c.JSON(200, err)
		return
	}
	msg := make(map[string]string)
	msg["message"] = "success"
	c.JSON(200, msg)
}
