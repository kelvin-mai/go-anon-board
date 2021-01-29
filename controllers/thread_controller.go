package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/models"
	"github.com/kelvin-mai/go-anon-board/response"
	"github.com/kelvin-mai/go-anon-board/services"
)

type ThreadController interface {
	ListThreads(c *gin.Context)
	GetThread(c *gin.Context)
	CreateThread(c *gin.Context)
	ReportThread(c *gin.Context)
	DeleteThread(c *gin.Context)
}

type threadController struct {
	ts services.ThreadService
}

func NewThreadController(ts services.ThreadService) ThreadController {
	return &threadController{ts}
}

func (tc *threadController) ListThreads(c *gin.Context) {
	offset := 0
	pageQuery := c.Query("page")
	if pageQuery != "" {
		page, err := strconv.Atoi(pageQuery)
		if err != nil {
			response.BadRequest(c, err)
			return
		}
		if page > 0 {
			offset = page - 1
		}
	}
	err, threads := tc.ts.List(offset)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.OK(c, threads)
	return
}

func (tc *threadController) GetThread(c *gin.Context) {
	id := c.Param("id")
	notfound, thread := tc.ts.GetByID(id)
	if notfound == true {
		response.ResourceNotFound(c, nil)
		return
	}
	response.OK(c, thread)
	return
}

func (tc *threadController) CreateThread(c *gin.Context) {
	var t models.Thread
	if err := c.ShouldBindJSON(&t); err != nil {
		response.BadRequest(c, err)
		return
	}
	err, thread := tc.ts.Create(t)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.Created(c, thread)
	return
}

func (tc *threadController) ReportThread(c *gin.Context) {
	id := c.Param("id")
	notfound, _ := tc.ts.GetByID(id)
	if notfound == true {
		response.ResourceNotFound(c, nil)
		return
	}
	err := tc.ts.Update(id, models.Thread{
		Reported: true,
	})
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.NoContent(c)
	return
}

func (tc *threadController) DeleteThread(c *gin.Context) {
	id := c.Param("id")
	notfound, _ := tc.ts.GetByID(id)
	if notfound == true {
		response.ResourceNotFound(c, nil)
		return
	}
	err := tc.ts.Update(id, models.Thread{
		Text: "[deleted]",
	})
	if err != nil {
		response.InternalServerError(c, err)
		return
	}
	response.NoContent(c)
	return
}
