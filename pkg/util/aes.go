package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func EncryptCFB(plaintext string, key string) (string, error) {
	plaintextByte := []byte(plaintext)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return plaintext, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintextByte))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return plaintext, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintextByte)
	return hex.EncodeToString(ciphertext), nil
}

func DecryptCFB(ciphertext string, key string) (string, error) {
	plaintext, err := hex.DecodeString(ciphertext)
	if err != nil {
		return ciphertext, err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ciphertext, err
	}
	iv := plaintext[:aes.BlockSize]
	plaintext = plaintext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, plaintext)
	return string(plaintext), nil
}
