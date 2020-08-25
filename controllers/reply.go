package controllers

import "github.com/gin-gonic/gin"

func getReply(c *gin.Context) {
	c.JSON(200, gin.H{"reply": "works"})
}

func RegisterReplyRoutes(r *gin.RouterGroup) {
	r.GET("/", getReply)
}
