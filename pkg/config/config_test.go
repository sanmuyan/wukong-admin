package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := Config{}
	if config.TokenTTL == 0 {
		t.Error(config)
	}
	t.Log(config)
}
