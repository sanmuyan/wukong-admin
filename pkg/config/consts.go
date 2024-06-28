package config

const (
	AppAccessTokenExpiration      = 3600
	MFABindTimeoutMin             = 10
	TOTPSecretLength              = 20
	TOTPTokenInterval             = 30
	TOTPTokenGracePeriod          = 1
	PassKeyRegistrationTimeoutMin = 10
	PassKeyMax                    = 5
	MFALoginTimeoutMin            = 5
	PassKeyLoginTimeoutMin        = 10
	OauthCodeTimeoutMin           = 5
)
