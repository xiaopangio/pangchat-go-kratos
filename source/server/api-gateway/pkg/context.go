package pkg

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func NewContext(c *gin.Context) (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(c, 5*time.Second)
}
func NewContextWithTimeout(c *gin.Context, timeout time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(c, timeout)
}
