package util

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	c := TokenClaims{}
	c.Body = []byte("")
	token, err := CreateToken(c, "")
	t.Log(token, err)
}
