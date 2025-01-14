package endpoints

import (
	"github.com/sprint1/config"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sprint1/internal/app/shortener/service"
)

func Test_SaveUrlHandler(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   string
	}

	type Expected struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name     string
		request  Request
		expected Expected
	}{
		{
			name: "Test Save URL successfully",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/",
				body:   "http://test",
			},
			expected: Expected{
				code:        http.StatusCreated,
				response:    `http://tes`,
				contentType: "",
			},
		},
		{
			name: "Test Save URL empty body",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/",
				body:   "",
			},
			expected: Expected{
				code:        http.StatusBadRequest,
				response:    "",
				contentType: "",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.request.method, test.request.url, strings.NewReader(test.request.body))
			w := httptest.NewRecorder()

			cfg := config.Init()
			router := mux.NewRouter()
			serviceImpl := service.NewService()
			controller := NewController(router, serviceImpl, cfg)
			controller.GetServeMux().ServeHTTP(w, r)

			result := w.Result()

			assert.Equal(t, test.expected.code, result.StatusCode, "unexpected status code")

			assert.Equal(t, test.expected.contentType, result.Header.Get("Content-Type"), "unexpected content type")

			resBody, err := io.ReadAll(result.Body)
			require.NoError(t, err, "error reading response body")
			assert.Equal(t, test.expected.response, string(resBody), "unexpected response body")
		})
	}
}
