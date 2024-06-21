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

func GetOauthCodeSession(c *gin.Context) {
	var req model.OauthCodeSessionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(200, model.NewOauthErrorResponse("invalid_request"))
		return
	}
	token, err := tokenutil.ValidToken(c)
	if err != nil {
		c.Redirect(302, fmt.Sprintf("%s%s", config.Conf.App.OauthLoginRedirectURL, c.Request.URL.String()))
		return
	}
	redirectURI, _err := svc.GetOauthCodeSession(token, &req)
	if _err != nil {
		c.JSON(200, _err)
		return
	}
	c.Redirect(302, redirectURI)
}

func GetOauthToken(c *gin.Context) {
	var req model.OauthTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, model.NewOauthErrorResponse("invalid_request"))
		return
	}
	switch req.GrantType {
	case "authorization_code":
		token, err := svc.GetOauthToken(&req)
		if err != nil {
			c.JSON(200, err)
			return
		}
		c.JSON(200, token)
	case "refresh_token":
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
