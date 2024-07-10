package model

// OauthCodeSession 会话数据
type OauthCodeSession struct {
	// Code 用于验证用户身份的字符串，使用 UUID 生成
	Code string `json:"code"`
	// ClientID 客户端 ID，用于验证所属应用
	ClientID string `json:"client_id"`
	// ClientSecret 可选项，如果存在，则用于验证客户端密钥
	ClientSecret string `json:"client_secret,omitempty"`
	// RedirectURI 回调 URI
	RedirectURI string `json:"redirect_uri"`
	// Scope 权限范围，多个权限用逗号隔开
	Scope string `json:"scope"`
}

// OauthTokenResponse 响应对象
type OauthTokenResponse struct {
	// AccessToken 身份验证令牌
	AccessToken string `json:"access_token"`
	// AccessToken 令牌过期时间
	ExpiresIn int `json:"expires_in"`
	// RefreshToken 刷新令牌
	RefreshToken string `json:"refresh_token,omitempty"`
	// TokenType 令牌类型，固定为 "Bearer"
	TokenType string `json:"token_type"`
}

// OauthCodeSessionRequest 请求对象
type OauthCodeSessionRequest struct {
	// ResponseType 根据 OAuth2 协议规范，这里固定为 "code"
	ResponseType string `form:"response_type" binding:"required"`
	// ClientID 客户端 ID，用于验证所属应用
	ClientID string `form:"client_id" binding:"required"`
	// RedirectURI 可选项，请求的回调地址
	RedirectURI string `form:"redirect_uri"`
	// Scope 可选项，请求的权限范围，多个权限用逗号隔开
	Scope string `form:"scope"`
	// State 可选项，系统将原样返回，可防止跨站请求伪造
	State string `form:"state"`
}

// OauthTokenRequest 请求对象
type OauthTokenRequest struct {
	// GrantType 根据 OAuth2 协议规范，允许的值是 "authorization_code" "refresh_token"
	// "authorization_code" 用于获取 access token
	// "refresh_token" 用于刷新 access token
	GrantType string `form:"grant_type" binding:"required"`
	// ClientID 客户端 ID，用于验证所属应用
	ClientID string `form:"client_id" binding:"required"`
	// ClientSecret 可选项，如果存在，则用于验证客户端密钥
	ClientSecret string `form:"client_secret"`
	// Code GrantType 是 authorization_code 时用于验证用户身份的字符串
	Code string `form:"code"`
	// RedirectURI 回调 URI
	RedirectURI string `form:"redirect_uri"`
	// RefreshToken GrantType 是 refresh_token 时用于刷新 access token 的令牌
	RefreshToken string `form:"refresh_token"`
}

// OauthRevokeTokenRequest 请求对象
type OauthRevokeTokenRequest struct {
	// ClientID 客户端 ID，用于验证所属应用
	ClientID string `form:"client_id" binding:"required"`
	// Token 要撤销的刷新令牌
	Token string `form:"token" binding:"required"`
	// ClientSecret 可选项，如果存在，则用于验证客户端密钥
	ClientSecret string `form:"client_secret"`
}

// OauthErrorResponse 响应对象
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
