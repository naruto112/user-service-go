package request

import (
	"net/http"
	"user-service/src/adapter/code"

	"github.com/gin-gonic/gin"
)

type MiddlewareRequest interface {
	ValidateTokenMiddleware() gin.HandlerFunc
}

type Middleware struct {
}

func ValidateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		// Verifique se o token é válido
		if !code.ValidateToken(&tokenString) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Next()
	}
}
