package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, resource interface{}) {
	c.JSON(http.StatusOK, resource)
}

func Created(c *gin.Context, resource interface{}) {
	c.JSON(http.StatusCreated, resource)
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
