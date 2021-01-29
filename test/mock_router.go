package test

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/kelvin-mai/go-anon-board/response"
)

func MockRouter() *gin.Engine {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		response.OK(c, gin.H{"health": "OK"})
		return
	})
	r.NoRoute(func(c *gin.Context) {
		response.ResourceNotFound(c, nil)
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
