package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	JWTSecret    []byte
	JWTExpiresIn time.Duration
)

func GetJWTExpiresIn(d string, default_ time.Duration) time.Duration {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return default_
	}
	return duration
}

func GenerateJWT(userID uint64) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(JWTExpiresIn).Unix(),
	})

	return t.SignedString(JWTSecret)
}

func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, ErrInvalidJWT
	}

	return claims, nil
}
