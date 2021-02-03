package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/models"
	"github.com/kelvin-mai/go-anon-board/response"
	"github.com/kelvin-mai/go-anon-board/services"
)

type ReplyController interface {
	GetReply(c *gin.Context)
	CreateReply(c *gin.Context)
	ReportReply(c *gin.Context)
	DeleteReply(c *gin.Context)
}

type replyController struct {
	rs services.ReplyService
}

func NewReplyController(rs services.ReplyService) ReplyController {
	return &replyController{rs}
}

func (rc *replyController) GetReply(c *gin.Context) {
	id := c.Param("id")
	err, reply := rc.rs.GetByID(id)
	if err != nil {
		response.ResourceNotFound(c, nil)
		return
	}
	response.OK(c, reply)
	return
}

func (rc *replyController) CreateReply(c *gin.Context) {
	var r models.Reply
	if err := c.ShouldBindJSON(&r); err != nil {
		response.BadRequest(c, err)
		return
	}
	err, reply := rc.rs.Create(r)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.Created(c, reply)
	return
}

func (rc *replyController) ReportReply(c *gin.Context) {
	id := c.Param("id")
	err, _ := rc.rs.GetByID(id)
	if err != nil {
		response.ResourceNotFound(c, nil)
		return
	}
	err = rc.rs.Update(id, models.Reply{
		Reported: true,
	})
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.NoContent(c)
	return
}

func (rc *replyController) DeleteReply(c *gin.Context) {
	id := c.Param("id")
	err, _ := rc.rs.GetByID(id)
	if err != nil {
		response.ResourceNotFound(c, nil)
		return
	}
	err = rc.rs.Update(id, models.Reply{
		Text: "[deleted]",
	})
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.NoContent(c)
	return
}
