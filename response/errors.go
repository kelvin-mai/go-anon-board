package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func InternalServerError(c *gin.Context, err error) {
	d := "internal server error"
	if err != nil {
		d = err.Error()
	}
	c.JSON(http.StatusInternalServerError, &ApiError{
		Status:      http.StatusInternalServerError,
		Message:     "internal server error",
		Description: d,
	})
}

func ResourceNotFound(c *gin.Context, err error) {
	d := "resource not found"
	if err != nil {
		d = err.Error()
	}
	c.JSON(http.StatusNotFound, &ApiError{
		Status:      http.StatusNotFound,
		Message:     "resource not found",
		Description: d,
	})
}

func BadRequest(c *gin.Context, err error) {
	d := "bad request"
	if err != nil {
		d = err.Error()
	}
	c.JSON(http.StatusBadRequest, &ApiError{
		Status:      http.StatusBadRequest,
		Message:     "bad request",
		Description: d,
	})
}
