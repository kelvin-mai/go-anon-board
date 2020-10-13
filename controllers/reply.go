package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/models"
)

func getReplies(c *gin.Context) {
	var r []models.Reply
	tid := c.Query("thread_id")
	db := models.GetDB()
	if tid == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "thread_id required",
		})
		return
	}
	db.Where("thread_id = ?", tid).Find(&r)
	c.JSON(http.StatusOK, r)
}

func createReply(c *gin.Context) {
	var t models.Thread
	tid, err := strconv.Atoi(c.PostForm("thread_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "thread_id required",
		})
		return
	}
	r := models.Reply{
		ThreadID:       uint(tid),
		Text:           c.PostForm("text"),
		DeletePassword: c.PostForm("password"),
	}
	db := models.GetDB()
	db.Model(&t).Where("id = ?", tid).Update("bumped_on", time.Now())
	db.Create(&r)
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func reportReply(c *gin.Context) {
	var r models.Reply
	db := models.GetDB()
	db.
		Where("thread_id = ?", c.PostForm("thread_id")).
		Where("reply_id = ?", c.PostForm("reply_id")).
		First(&r).
		Update("reported", true)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func softDeleteReply(c *gin.Context) {
	var r models.Reply
	db := models.GetDB()
	db.
		Where("thread_id = ?", c.PostForm("thread_id")).
		Where("reply_id = ?", c.PostForm("reply_id")).
		First(&r)
	if r.DeletePassword == c.PostForm("password") {
		db.Update("text", "[deleted]")
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "incorrect password"})
	}
}

func RegisterReplyRoutes(r *gin.RouterGroup) {
	r.GET("/", getReplies)
	r.POST("/", createReply)
	r.PUT("/", reportThread)
	r.DELETE("/", softDeleteReply)
}
