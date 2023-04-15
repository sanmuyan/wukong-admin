package util

import (
	"os"
	"testing"
)

func TestEncryptCFB(t *testing.T) {
	plaintext := "123"
	res, err := EncryptCFB(plaintext, os.Getenv("CONFIG_SECRET_KEY"))
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestDecryptCFB(t *testing.T) {
	ciphertext := "123"
	res, err := DecryptCFB(ciphertext, os.Getenv("CONFIG_SECRET_KEY"))
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
