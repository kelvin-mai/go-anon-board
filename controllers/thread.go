package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/models"
	"github.com/kelvin-mai/go-anon-board/utils"
)

func listThreads(c *gin.Context) {
	var b models.Board
	var t []models.Thread
	models.FindOrCreateBoard(c.Param("board"), &b)
	db := models.GetDB()
	db.Where("board_id = ?", b.ID).Preload("Replies").Order("bumped_on DESC").Limit(10).Find(&t)
	c.JSON(http.StatusOK, t)
}

func createThread(c *gin.Context) {
	var b models.Board
	db := models.GetDB()
	models.FindOrCreateBoard(c.Param("board"), &b)
	t := models.Thread{
		Text:           c.PostForm("text"),
		DeletePassword: c.PostForm("password"),
		BoardID:        b.ID,
	}
	db.Create(&t)
	if utils.IsBrowser(c) {
		c.Redirect(http.StatusPermanentRedirect, "/b/"+c.Param("board"))
		return
	} else {
		c.JSON(http.StatusCreated, t)
		return
	}
}

func reportThread(c *gin.Context) {
	var t models.Thread
	db := models.GetDB()
	result := db.Where("id = ?", c.PostForm("id")).First(&t).Update("reported", true)
	if result.RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "thread not found",
		})
		return
	}
	if utils.IsBrowser(c) {
		c.Redirect(http.StatusPermanentRedirect, "/b/"+c.Param("board"))
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"thread":  t,
		})
		return
	}
}

func deleteThread(c *gin.Context) {
	var t models.Thread
	db := models.GetDB()
	result := db.Where("id = ?", c.PostForm("id")).First(&t)
	if result.RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "thread not found",
		})
		return
	}
	if t.DeletePassword == c.PostForm("password") {
		db.Delete(&t)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "incorrect password",
		})
	}
}

func RegisterThreadRoutes(r *gin.RouterGroup) {
	r.GET("/", listThreads)
	r.POST("/", createThread)
	r.PUT("/", reportThread)
	r.DELETE("/", deleteThread)
}
