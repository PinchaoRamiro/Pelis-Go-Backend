package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GetJWTPassword() ([]byte, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return []byte(os.Getenv("JWT_PASSWORD")), nil
}

func GenerateToken(email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 4).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey, err := GetJWTPassword()
	if err != nil {
		return "", err
	}

	return token.SignedString(secretKey)
}
