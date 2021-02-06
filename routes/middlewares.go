package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiKey(param string, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		attempt := c.Request.Header.Get(param)
		if attempt != key {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Invalid api key",
			})
			return
		}
		c.Next()
	}
}
