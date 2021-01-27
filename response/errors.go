package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ApiError struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func createApiError(status int, message string, err error) (int, *ApiError) {
	d := message
	if err != nil {
		log.Error(err.Error())
		d = err.Error()
	}
	return status, &ApiError{
		Status:      status,
		Message:     message,
		Description: d,
	}
}

func InternalServerError(c *gin.Context, err error) {
	c.JSON(createApiError(http.StatusInternalServerError, "internal server error", err))
}

func ResourceNotFound(c *gin.Context, err error) {
	c.JSON(createApiError(http.StatusNotFound, "resource not found", err))
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(createApiError(http.StatusBadRequest, "bad request", err))
}
