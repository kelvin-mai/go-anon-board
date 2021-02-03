package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/response"
	"github.com/kelvin-mai/go-anon-board/services"
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
	err, _ := a.ts.GetByID(id)
	if err != nil {
		response.ResourceNotFound(c, nil)
		return
	}
	err = a.ts.Delete(id)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.NoContent(c)
	return
}

func (a *adminController) DeleteReply(c *gin.Context) {
	id := c.Param("id")
	err, _ := a.rs.GetByID(id)
	if err != nil {
		response.ResourceNotFound(c, nil)
		return
	}
	err = a.rs.Delete("id")
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.NoContent(c)
	return
}
