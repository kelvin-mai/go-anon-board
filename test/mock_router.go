package test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/utils"
)

func MockRouter() *gin.Engine {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "OK"})
		return
	})
	r.NoRoute(func(c *gin.Context) {
		utils.CreateApiError(http.StatusNotFound, errors.New("route not found"))
		return
	})

	return r
}

func PerformRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
