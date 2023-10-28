package mycrypto

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const (
	TokenSize = 16
	SecretKey = "9Xa15pap24"
)

func GenRandomToken() (string, error) {
	b := make([]byte, TokenSize)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("create random slise error: %w", err)
	}

	return hex.EncodeToString(b), nil
}

func HashFunc(src string) string {
	initString := fmt.Sprintf("%v:%v", src, SecretKey)
	return fmt.Sprintf("%x", md5.Sum([]byte(initString)))
}
