// Package middleware  @Author xiaobaiio 2023/3/24 16:06:00
package middleware

import (
	"connector/internal/components/auth"
	"connector/pkg"
	"github.com/gin-gonic/gin"
)

func Auth(manager *auth.JwtManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		userToken, err := manager.ParseUser(token)
		if err != nil {
			pkg.Forbidden(c)
		} else {
			c.Set("token", userToken)
			c.Next()
		}
	}
}
