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
		RPDisplayName: config.Conf.Basic.AppName,
		RPID:          config.Conf.Basic.SiteHost,
		RPOrigins:     []string{config.Conf.Basic.SiteURL},
		Timeouts: webauthn.TimeoutsConfig{
			Login:        webauthn.TimeoutConfig{Enforce: true, Timeout: config.PassKeyRegistrationTimeoutMin * time.Minute},
			Registration: webauthn.TimeoutConfig{Enforce: true, Timeout: config.PassKeyRegistrationTimeoutMin * time.Minute},
		},
		AuthenticatorSelection: protocol.AuthenticatorSelection{
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
