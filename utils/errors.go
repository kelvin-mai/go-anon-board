package utils

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ApiError struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func CreateApiError(status int, err error) (int, *ApiError) {
	log.Error(err.Error())
	message := err.Error()
	return status, &ApiError{
		Status:  status,
		Message: message,
	}
}

func ErrorFromDatabase(err error) (int, *ApiError) {
	switch err {
	case gorm.ErrRecordNotFound:
		return CreateApiError(http.StatusNotFound, err)
	default:
		return CreateApiError(http.StatusInternalServerError, err)
	}
}
