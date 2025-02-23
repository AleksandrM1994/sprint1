package endpoints

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *EndpointsTestSuite) Test_GetShortenURLHandler(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   *GetShortenURLRequest
	}

	type Expected struct {
		code        int
		response    *GetShortenURLResponse
		contentType string
	}
	tests := []struct {
		name     string
		request  Request
		expected Expected
	}{
		{
			name: "Test Get Shorten URL successfully",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/api/shorten",
				body: &GetShortenURLRequest{
					URL: "https://practicum.yandex.ru",
				},
			},
			expected: Expected{
				code: http.StatusCreated,
				response: &GetShortenURLResponse{
					Result: "http://localhost:8080/8a9923515b446",
				},
			},
		},
		{
			name: "Test Get Shorten URL empty request body",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/api/shorten",
				body:   &GetShortenURLRequest{},
			},
			expected: Expected{
				code:        http.StatusBadRequest,
				contentType: "",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.request.body)
			r := httptest.NewRequest(test.request.method, test.request.url, strings.NewReader(string(body)))
			w := httptest.NewRecorder()

			suite.controller.GetServeMux().ServeHTTP(w, r)

			result := w.Result()

			assert.Equal(t, test.expected.code, result.StatusCode, "unexpected status code")

			assert.Equal(t, test.expected.contentType, result.Header.Get("Content-Type"), "unexpected content type")

			if result.StatusCode == http.StatusCreated {
				resBody, err := io.ReadAll(result.Body)
				require.NoError(t, err, "error reading response body")
				res := &GetShortenURLResponse{}
				_ = json.Unmarshal(resBody, res)
				assert.Equal(t, test.expected.response, res, "unexpected response body")
			}
		})
	}
}
