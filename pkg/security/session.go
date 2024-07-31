package security

import (
	"github.com/google/uuid"
	"github.com/sanmuyan/xpkg/xcrypto"
	"github.com/sanmuyan/xpkg/xutil"
	"wukong/pkg/config"
)

const (
	SessionIDLength = 36
)

func GetRandomID() string {
	return uuid.NewString()
}

func GetSessionID() string {
	id := GetRandomID()
	return xutil.RemoveError(xcrypto.GenerateHmacSha1(id, config.Conf.Secret.SessionKey)) + id
}

func VerifySessionID(sessionID string) bool {
	sl := len(sessionID)
	if sl <= SessionIDLength {
		return false
	}
	hl := sl - SessionIDLength
	message := sessionID[hl:]
	hash := sessionID[:hl]
	return xutil.RemoveError(xcrypto.GenerateHmacSha1(message, config.Conf.Secret.SessionKey)) == hash
}
