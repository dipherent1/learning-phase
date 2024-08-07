package middleware

import (
	"net/http"
	"os"
	"strings"
	"tskmgr/models"

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

	claim := models.Claims{}
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

func IsSuperUser(c *gin.Context) {
	claim, exists := c.Get("claim")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "claim not set"})
		c.Abort()
		return
	}

	// Type assertion to your custom Claims type
	userClaims, ok := claim.(models.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claim type"})
		c.Abort()
		return
	}

	if userClaims.UserRole != "superuser" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "only superuser can acces this page"})
		c.Abort()
		return

	}

}
