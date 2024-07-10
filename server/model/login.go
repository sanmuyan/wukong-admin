package model

// LoginRequest 请求对象
type LoginRequest struct {
	// Username 用户名
	Username string `json:"username" binding:"required"`
	// Password 密码，应使用 RSA 证书加密
	Password string `json:"password" binding:"required"`
}

// LoginResponse 响应对象
type LoginResponse struct {
	// Token 登录成功返回的身份验证令牌
	Token string `json:"token,omitempty"`
	// MFABeginLogin 可选项，需要用户二次验证
	MFABeginLogin *MFABeginLoginResponse `json:"mfa_begin_login,omitempty"`
	// RequireMFA 可选项，要求用户绑定 MFA 应用
	RequireMFA *MFAAppBindResponse `json:"require_mfa,omitempty"`
	// PassKeyBeginLogin 可选项，需要用户二次验证
	PassKeyBeginLogin *PassKeyBeginLoginResponse `json:"pass_key_begin_login,omitempty"`
}
