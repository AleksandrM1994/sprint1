package endpoints

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *EndpointsTestSuite) Test_SaveUrlHandler(t *testing.T) {
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
				body:   "https://practicum.yandex.ru",
			},
			expected: Expected{
				code:        http.StatusCreated,
				response:    "http://localhost:8080/aHR0cHM6Ly9wcmFjdGljdW0ueWFuZGV4LnJ1",
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

			suite.controller.GetServeMux().ServeHTTP(w, r)

			result := w.Result()

			assert.Equal(t, test.expected.code, result.StatusCode, "unexpected status code")

			assert.Equal(t, test.expected.contentType, result.Header.Get("Content-Type"), "unexpected content type")

			resBody, err := io.ReadAll(result.Body)
			require.NoError(t, err, "error reading response body")
			assert.Equal(t, test.expected.response, string(resBody), "unexpected response body")
		})
	}
}
