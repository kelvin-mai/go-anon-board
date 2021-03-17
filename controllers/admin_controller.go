package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/services"
	"github.com/kelvin-mai/go-anon-board/utils"
)

type AdminController interface {
	DeleteThread(c *gin.Context)
	DeleteReply(c *gin.Context)
}

type adminController struct {
	ts services.ThreadService
	rs services.ReplyService
}

func NewAdminController(ts services.ThreadService, rs services.ReplyService) AdminController {
	return &adminController{ts, rs}
}

func (a *adminController) DeleteThread(c *gin.Context) {
	id := c.Param("id")
	err := a.ts.Delete(id)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	return
}

func (a *adminController) DeleteReply(c *gin.Context) {
	id := c.Param("id")
	err := a.rs.Delete(id)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	return
}
