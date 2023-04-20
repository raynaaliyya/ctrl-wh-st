package router

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("enigmacamp")

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("claims", claims)

		c.Next()
	}
}

func profile(c *gin.Context) {
	// ambil name dari JWT token
	claims := c.MustGet("claims").(jwt.MapClaims)
	name := claims["name"].(string)

	// dapatkan informasi user dari database (dalam hal ini, kita return name)
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to profile",
		"name":    name,
	})
}
