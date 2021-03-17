package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/utils"
)

func ApiKey(param, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		attempt := c.Request.Header.Get(param)
		if attempt != key {
			c.AbortWithStatusJSON(utils.CreateApiError(http.StatusUnauthorized, errors.New("invalid api key")))
			return
		}
		c.Next()
	}
}
