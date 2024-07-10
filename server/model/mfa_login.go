package model

const (
	// MFAProviderMFAApp MFA 应用
	MFAProviderMFAApp = "mfa_app"
)

// MFALoginSession 会话数据
type MFALoginSession struct {
	MFAProvider string `json:"mfa_provider"`
}

// MFAFinishLoginRequest 请求对象
type MFAFinishLoginRequest struct {
	SessionID   string `json:"session_id" binding:"required"`
	MFAProvider string `json:"mfa_provider" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Code        string `json:"code" binding:"required"`
}

// MFABeginLoginResponse 返回对象
type MFABeginLoginResponse struct {
	SessionID   string `json:"session_id"`
	Username    string `json:"username"`
	MFAProvider string `json:"mfa_provider"`
}
