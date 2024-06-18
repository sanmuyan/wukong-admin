package config

const (
	AppAccessTokenExpiration      = 3600
	MFABindTimeoutMin             = 10
	TOTPSecretLength              = 20
	TOTPTokenInterval             = 30
	PassKeyRegistrationTimeoutMin = 10
	PassKeyMax                    = 5
	MFALoginTimeoutMin            = 5
	PassKeyLoginTimeoutMin        = 10
	OauthCodeTimeoutMin           = 5
	PasswordMinLength             = 8
	PasswordMinIncludeCase        = 3
)
