package model

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token             string                     `json:"token,omitempty"`
	MFABeginLogin     *MFABeginLoginResponse     `json:"mfa_begin_login,omitempty"`
	PassKeyBeginLogin *PassKeyBeginLoginResponse `json:"pass_key_begin_login,omitempty"`
}
