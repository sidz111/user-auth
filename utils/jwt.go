package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateJWT(username string, id uint) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "failed to create jwt", err
	}

	claims := jwt.MapClaims{
		"username": username,
		"user_id":  id,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "failed to create jwt", err
	}
	return tokenString, nil
}
