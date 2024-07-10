package passkey

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
	"wukong/pkg/config"
)

var (
	WebAuthn *sync.Map
)

func InitWebAuthnConfig() {
	WebAuthn = new(sync.Map)
	conf := &webauthn.Config{
		// 依赖方系统显示名
		RPDisplayName: config.Conf.Basic.AppName,
		// 依赖方系统 ID
		RPID: config.Conf.Basic.SiteHost,
		// 依赖方允许的源站点 URL
		RPOrigins: []string{config.Conf.Basic.SiteURL},
		Timeouts: webauthn.TimeoutsConfig{
			Login:        webauthn.TimeoutConfig{Enforce: true, Timeout: config.PassKeyRegistrationTimeoutMin * time.Minute},
			Registration: webauthn.TimeoutConfig{Enforce: true, Timeout: config.PassKeyRegistrationTimeoutMin * time.Minute},
		},
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			// 不要求身份验证器再次验证身份
			UserVerification: protocol.VerificationDiscouraged,
		},
	}
	webAuthn, err := webauthn.New(conf)
	if err != nil {
		logrus.Errorf("failed to initialize webAuthn: %s", err)
		return
	}
	WebAuthn.Store(conf.RPID, webAuthn)
}
