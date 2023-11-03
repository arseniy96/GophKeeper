package mycrypto

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID int64
}

var (
	ErrInvalidToken = errors.New("invalid token")
)

func HashFunc(src, secret string) (string, error) {
	initString := fmt.Sprintf("%v:%v", src, secret)
	hash, err := bcrypt.GenerateFromPassword([]byte(initString), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt error: %w", err)
	}

	return string(hash), nil
}

func BuildJWT(userID int64, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{},
		UserID:           userID,
	})

	return token.SignedString([]byte(secret))
}

func GetUserID(tokenString, secret string) (int64, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(secret), nil
		})
	if err != nil {
		return 0, fmt.Errorf("jwt parse error: %v", err)
	}

	if !token.Valid {
		return 0, fmt.Errorf("%w: Get user id error", ErrInvalidToken)
	}

	return claims.UserID, nil
}
