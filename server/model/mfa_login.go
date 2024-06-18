package model

const (
	MFAProviderMFAApp = "mfa_app"
)

type MFALoginSession struct {
	MFAProvider string `json:"mfa_provider"`
}

type MFAFinishLoginRequest struct {
	SessionID   string `json:"session_id" binding:"required"`
	MFAProvider string `json:"mfa_provider" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Code        string `json:"code" binding:"required"`
}

type MFABeginLoginResponse struct {
	SessionID   string `json:"session_id"`
	Username    string `json:"username"`
	MFAProvider string `json:"mfa_provider"`
}
