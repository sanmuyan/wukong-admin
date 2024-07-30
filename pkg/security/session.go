package security

import (
	"github.com/google/uuid"
	"github.com/sanmuyan/xpkg/xcrypto"
	"wukong/pkg/config"
)

const (
	SessionIDLength     = 36
	SessionIDHashLength = 64
)

func GetRandomID() string {
	return uuid.NewString()
}

func GetSessionID() string {
	id := GetRandomID()
	return id + xcrypto.GenerateHmacSha256(id, config.Conf.Secret.TokenKey)
}

func VerifySessionID(sessionID string) bool {
	if len(sessionID) != (SessionIDLength + SessionIDHashLength) {
		return false
	}
	message := sessionID[:SessionIDLength]
	hash := sessionID[SessionIDLength:]
	return xcrypto.GenerateHmacSha256(message, config.Conf.Secret.TokenKey) == hash
}
