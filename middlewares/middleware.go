package middlewares

import "github.com/gin-gonic/gin"

func check() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path != "/test" {
			c.AbortWithStatusJSON(500, gin.H{"message": "path fail"})
		}
	}
}
