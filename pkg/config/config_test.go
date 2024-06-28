package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := Config{}
	if config.Database.Mysql == "" {
		t.Error(config)
	}
	t.Log(config)
}
