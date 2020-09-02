package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/models"
)

func index(c *gin.Context) {
	var b []models.Board
	db := models.GetDB()
	db.Find(&b)
	c.HTML(http.StatusOK, "index.tmpl",
		gin.H{
			"title":  "Anonymous Board",
			"boards": b,
		})
}

func boardPage(c *gin.Context) {
	var b models.Board
	var t []models.Thread
	board := c.Param("board")
	models.FindOrCreateBoard(board, &b)
	db := models.GetDB()
	db.Where("board_id = ?", b.ID).Preload("Replies").Order("bumped_on DESC").Limit(10).Find(&t)
	c.HTML(http.StatusOK, "board.tmpl",
		gin.H{
			"title":   board,
			"threads": t,
		})
}

func RegisterHtmlRoutes(r *gin.RouterGroup) {
	r.GET("/", index)
	b := r.Group("/b/:board")
	b.GET("/", boardPage)
}
