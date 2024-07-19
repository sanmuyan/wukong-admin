package config

// 系统固定配置
const (
	// AppAccessTokenExpiration access token 过期时间，单位秒
	AppAccessTokenExpiration = 3600
	// MFABindTimeoutMin 绑定 session 过期时间，单位分钟
	MFABindTimeoutMin = 10
	// TOTPSecretLength TOTP 密钥长度
	TOTPSecretLength = 20
	// TOTPTokenInterval Token 间隔时间，单位秒
	TOTPTokenInterval = 30
	// TOTPTokenGracePeriod Token 宽限次数
	TOTPTokenGracePeriod = 1
	// PassKeyRegistrationTimeoutMin 注册 session 过期时间，单位分钟
	PassKeyRegistrationTimeoutMin = 10
	// PassKeyMax 用户注册最大数量
	PassKeyMax = 5
	// MFALoginTimeoutMin 登录 session 过期时间，单位分钟
	MFALoginTimeoutMin = 5
	// PassKeyLoginTimeoutMin 登录 session 过期时间，单位分钟
	PassKeyLoginTimeoutMin = 10
	// OauthCodeTimeoutMin 登录 session 过期时间，单位分钟
	OauthCodeTimeoutMin = 5
	// OauthLoginTimeoutMin 第三方登录超时，单位分钟
	OauthLoginTimeoutMin = 5
)
