package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	m "github.com/xerardoo/sapip/models"
)

func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")

		// Fix for testing purposes
		env := os.Getenv("APP_ENV")
		if env == "testing" {
			c.Set("USER", "1")
			c.Next()
			return
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "No autorizado"})
			return
		}
		claim, err := m.ValidateTokenJWT(token, c.FullPath(), c.Request.Method)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "No autorizado "})
			return
		}
		c.Set("USER", claim.User)
		c.Next()
	}
}
