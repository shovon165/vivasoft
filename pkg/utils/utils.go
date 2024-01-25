package utils

import (
	"book-crud/pkg/config"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}
func ComparePassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func GetJwtForUser(username string) (string, error) {
	now := time.Now().UTC()
	ttl := time.Minute * time.Duration(config.LocalConfig.JwtExpireMinutes)
	claims := jwt.StandardClaims{
		ExpiresAt: now.Add(ttl).Unix(),
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
		Subject:   username,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.LocalConfig.JwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}
