package model

type OauthCodeSession struct {
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

type OauthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type"`
}

type OauthCodeSessionRequest struct {
	ResponseType string `form:"response_type" binding:"required"`
	ClientID     string `form:"client_id" binding:"required"`
	RedirectURI  string `form:"redirect_uri"`
	Scope        string `form:"scope"`
	State        string `form:"state"`
	RedirectType string `form:"redirect_type"`
}

type OauthTokenRequest struct {
	GrantType    string `form:"grant_type" binding:"required"`
	ClientSecret string `form:"client_secret"`
	ClientID     string `form:"client_id" binding:"required"`
	Code         string `form:"code"`
	RedirectURI  string `form:"redirect_uri"`
	RefreshToken string `form:"refresh_token"`
}

type OauthRevokeTokenRequest struct {
	ClientID     string `form:"client_id" binding:"required"`
	Token        string `form:"token" binding:"required"`
	ClientSecret string `form:"client_secret"`
}

type OauthErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
}

func NewOauthErrorResponse(err string) *OauthErrorResponse {
	return &OauthErrorResponse{
		Error: err,
	}
}

func (c *OauthErrorResponse) WithDes(des string) *OauthErrorResponse {
	c.ErrorDescription = des
	return c
}
