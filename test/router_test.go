package test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouterDefaults(t *testing.T) {
	r := MockRouter()

	tests := []struct {
		name     string
		method   string
		path     string
		status   int
		response map[string]interface{}
	}{
		{"healthcheck works", "GET", "/", http.StatusOK,
			map[string]interface{}{
				"health": "OK",
			},
		},
		{"not found works", "GET", "/404", http.StatusNotFound,
			map[string]interface{}{
				"status":      (float64)(http.StatusNotFound),
				"message":     "resource not found",
				"description": "resource not found",
			},
		},
		{"not found works for other methods", "POST", "/404", http.StatusNotFound,
			map[string]interface{}{
				"status":      (float64)(http.StatusNotFound),
				"message":     "resource not found",
				"description": "resource not found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := PerformRequest(r, tt.method, tt.path, nil)
			assert.Equal(t, tt.status, w.Code)
			var response interface{}
			err := json.Unmarshal([]byte(w.Body.String()), &response)
			assert.Nil(t, err)
			assert.Equal(t, tt.response, response)
		})
	}
}
