package tokenclient

import (
	"github.com/sirupsen/logrus"
	"time"
	"wukong/pkg/config"
)

type TokenClient interface {
	IsTokenExist(string, string, string) bool
	SetToken(string, string, string, time.Duration) error
	DeleteToken(string, string) error
}

var TC TokenClient

func InitTokenClient() {
	tokenClients := map[string]TokenClient{
		"redis": NewRDBTokenClient(),
		"mysql": NewMySQLTokenClient(),
	}
	if _, ok := tokenClients[config.Conf.TokenClient]; !ok {
		logrus.Fatalf("token client %s not supported", config.Conf.TokenClient)
	}
	TC = tokenClients[config.Conf.TokenClient]
}
