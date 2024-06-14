package passkey

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/sirupsen/logrus"
	"time"
	"wukong/pkg/config"
)

var (
	WebAuthn *webauthn.WebAuthn
	err      error
)

func InitWebAuthnConfig(rpName string, rpID string, rpOrigins []string) {
	conf := &webauthn.Config{
		RPDisplayName: rpName,
		RPID:          rpID,
		RPOrigins:     rpOrigins,
		Timeouts: webauthn.TimeoutsConfig{
			Login:        webauthn.TimeoutConfig{Enforce: true, Timeout: config.PassKeyRegistrationTimeoutMin * time.Minute},
			Registration: webauthn.TimeoutConfig{Enforce: true, Timeout: config.PassKeyRegistrationTimeoutMin * time.Minute},
		},
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			UserVerification: protocol.VerificationDiscouraged,
		},
	}
	if WebAuthn, err = webauthn.New(conf); err != nil {
		logrus.Fatalf("failed to initialize webAuthn: %s", err)
	}
}
