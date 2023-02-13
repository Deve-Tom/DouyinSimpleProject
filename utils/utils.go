package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("SGVsbG9Xb3JsZA")

type CustomClaims struct {
	UserID uint
	jwt.RegisteredClaims
}

func GenToken(uid uint) (string, error) {
	claims := &CustomClaims{
		UserID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 30 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "HelloWorld",
			Subject:   "douyin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*CustomClaims, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}
