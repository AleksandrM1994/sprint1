package endpoints

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sprint1/internal/app/shortener/repository"
)

func (suite *EndpointsTestSuite) Test_GetOriginalUrlHandler(t *testing.T) {
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
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.request.method, test.request.url, strings.NewReader(test.request.body))
			w := httptest.NewRecorder()

			suite.repo.EXPECT().GetURLByShortURL(gomock.Any(), "aHR0cHM6Ly9qc29uZm9ybWF0dGVyLm9yZw==").Return(&repository.URL{
				ID:          1,
				ShortURL:    "aHR0cHM6Ly9qc29uZm9ybWF0dGVyLm9yZw==",
				OriginalURL: "https://jsonformatter.org",
			}, nil).MaxTimes(1)

			suite.controller.GetServeMux().ServeHTTP(w, r)

			result := w.Result()
			defer func() {
				if err := result.Body.Close(); err != nil {
					fmt.Println("Body.Close:", err)
				}
			}()

			assert.Equal(t, test.expected.code, result.StatusCode, "unexpected status code")

			assert.Equal(t, test.expected.location, result.Header.Get("Location"), "unexpected location")
		})
	}
}
