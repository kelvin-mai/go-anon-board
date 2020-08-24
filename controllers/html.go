package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl",
		gin.H{
			"title": "Home Page",
		})
}

func RegisterHtmlRoutes(r *gin.RouterGroup) {
	r.GET("/", index)
}
