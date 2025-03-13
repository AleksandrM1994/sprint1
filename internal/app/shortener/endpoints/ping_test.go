package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (suite *EndpointsTestSuite) Test_PingHandler(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   io.Reader
	}

	type Expected struct {
		code int
	}
	tests := []struct {
		name     string
		request  Request
		expected Expected
	}{
		{
			name: "Test Ping successfully",
			request: Request{
				method: http.MethodGet,
				url:    "/ping",
			},
			expected: Expected{
				code: http.StatusOK,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.request.body)
			r := httptest.NewRequest(test.request.method, test.request.url, strings.NewReader(string(body)))
			w := httptest.NewRecorder()

			suite.repo.EXPECT().Ping().Return(nil).Times(1)

			suite.controller.GetServeMux().ServeHTTP(w, r)

			result := w.Result()
			defer func() {
				if err := result.Body.Close(); err != nil {
					fmt.Println("Body.Close:", err)
				}
			}()

			assert.Equal(t, test.expected.code, result.StatusCode, "unexpected status code")
		})
	}
}
