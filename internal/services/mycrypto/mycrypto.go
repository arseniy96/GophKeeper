package mycrypto

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"

	"github.com/arseniy96/GophKeeper/src/logger"
)

const (
	TokenSize = 16
	SecretKey = "9Xa15pap24"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID int64
}

var (
	ErrInvalidToken = errors.New("invalid token")
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

func BuildJWT(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{},
		UserID:           userID,
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserID(tokenString string, l *logger.Logger) (int64, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(SecretKey), nil
		})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		l.Log.Error("token is not valid")
		return 0, ErrInvalidToken
	}

	l.Log.Info("token is valid")
	return claims.UserID, nil
}
