package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	Body []byte `json:"body"`
	jwt.StandardClaims
}

func CreateToken(claims TokenClaims, key string) (tokenStr string, err error) {
	if len(key) < 16 {
		return tokenStr, errors.New("key length less 16")
	}
	k := []byte(key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(k)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

func ExtractToken(tokenStr string, key string) (*TokenClaims, error) {
	k := []byte(key)
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (any, error) {
		return k, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, errors.New("invalid token")
	}
}
