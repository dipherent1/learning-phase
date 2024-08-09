package infrastructure

import (
	"net/http"
	"os"
	"strings"
	domain "tskmgr/Domain"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	key := os.Getenv("JWT_SECRET")
	jwtSecret := []byte(key)

	if jwtSecret == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET environment variable not set"})
		c.Abort()
		return
	}

	claim := domain.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("claim", claim)

	c.Next()
}
