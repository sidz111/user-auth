package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := godotenv.Load(); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Empty Token",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return os.Getenv("SECRET_KEY"), nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "user not valid",
			})
			ctx.Abort()
			return
		}
	}
}
