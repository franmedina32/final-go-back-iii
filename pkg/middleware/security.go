package middleware

import (
	"errors"
	"final-go-back-III/pkg/web"
	"github.com/gin-gonic/gin"
	"os"
)

// Authentication manages the security by validating the token
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("auth-key")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			c.Abort()
			return
		}
		if token != os.Getenv("AUTHENTICATION_TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			c.Abort()
			return
		}
		c.Next()
	}
}
