package controllers

import "github.com/gin-gonic/gin"

func getThread(c *gin.Context) {
	c.JSON(404, gin.H{"thread": "works"})
}

func RegisterThreadRoutes(r *gin.RouterGroup) {
	r.GET("/", getThread)
}
