package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}
