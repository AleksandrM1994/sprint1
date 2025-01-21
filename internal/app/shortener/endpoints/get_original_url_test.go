package endpoints

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sprint1/internal/app/shortener/service"
)

func Test_GetOriginalUrlHandler(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   string
	}

	type Expected struct {
		code     int
		location string
	}
	tests := []struct {
		name     string
		request  Request
		expected Expected
	}{
		{
			name: "Test Get Original URL successfully",
			request: Request{
				method: http.MethodGet,
				url:    "http://localhost:8080/aHR0cHM6Ly9qc29uZm9ybWF0dGVyLm9yZw==",
			},
			expected: Expected{
				code:     http.StatusTemporaryRedirect,
				location: "https://jsonformatter.org",
			},
		},
		{
			name: "Test Get Original URL not find original url",
			request: Request{
				method: http.MethodGet,
				url:    "http://localhost:8080/123",
			},
			expected: Expected{
				code:     http.StatusBadRequest,
				location: "",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.request.method, test.request.url, strings.NewReader(test.request.body))
			w := httptest.NewRecorder()

			mux := http.NewServeMux()
			serviceImpl := service.NewService()
			serviceImpl.OriginalURLsMap = map[string]string{"aHR0cHM6Ly9qc29uZm9ybWF0dGVyLm9yZw==": "https://jsonformatter.org"}
			controller := NewController(mux, serviceImpl)
			controller.GetServeMux().ServeHTTP(w, r)

			result := w.Result()

			assert.Equal(t, test.expected.code, result.StatusCode, "unexpected status code")

			assert.Equal(t, test.expected.location, result.Header.Get("Location"), "unexpected location")
		})
	}
}
