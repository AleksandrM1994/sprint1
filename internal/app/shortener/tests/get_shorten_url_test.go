package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sprint1/internal/app/shortener/endpoints/public"
	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/repository"
)

func (suite *EndpointsTestSuite) Test_GetShortenURLHandler_Success(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   interface{}
	}

	type Expected struct {
		code        int
		response    *public.GetShortenURLResponse
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
				body: &public.GetShortenURLRequest{
					URL: "https://duckduckgo.com",
				},
			},
			expected: Expected{
				code: http.StatusCreated,
				response: &public.GetShortenURLResponse{
					Result: "http://localhost:8080/c489a87f9b3b",
				},
				contentType: "application/json",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.request.body)
			r := httptest.NewRequest(test.request.method, test.request.url, strings.NewReader(string(body)))
			w := httptest.NewRecorder()

			suite.repo.EXPECT().CreateURL(
				gomock.Any(),
				"c489a87f9b3b",
				"https://duckduckgo.com",
				"",
			).Return(nil).Times(1)
			suite.repo.EXPECT().GetURLByShortURL(gomock.Any(), "c489a87f9b3b").Return(&repository.URL{
				ID:          1,
				ShortURL:    "c489a87f9b3b",
				OriginalURL: "https://duckduckgo.com",
			}, nil).Times(1)

			suite.controller.GetServeMux().ServeHTTP(w, r)

			result := w.Result()
			defer func() {
				if err := result.Body.Close(); err != nil {
					fmt.Println("Body.Close:", err)
				}
			}()

			assert.Equal(t, test.expected.code, result.StatusCode, "unexpected status code")

			assert.Equal(t, test.expected.contentType, result.Header.Get("Content-Type"), "unexpected content type")

			if result.StatusCode == http.StatusCreated {
				resBody, err := io.ReadAll(result.Body)
				require.NoError(t, err, "error reading response body")
				res := &public.GetShortenURLResponse{}
				_ = json.Unmarshal(resBody, res)
				assert.Equal(t, test.expected.response, res, "unexpected response body")
			}
		})
	}
}

func (suite *EndpointsTestSuite) Test_GetShortenURLHandler_ConflictError(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   interface{}
	}

	type Expected struct {
		code        int
		response    *public.GetShortenURLResponse
		contentType string
	}
	tests := []struct {
		name     string
		request  Request
		expected Expected
	}{
		{
			name: "Test Get Shorten URL conflict error",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/api/shorten",
				body: &public.GetShortenURLRequest{
					URL: "https://duckduckgo.com",
				},
			},
			expected: Expected{
				code: http.StatusConflict,
				response: &public.GetShortenURLResponse{
					Result: "http://localhost:8080/c489a87f9b3b",
				},
				contentType: "application/json",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.request.body)
			r := httptest.NewRequest(test.request.method, test.request.url, strings.NewReader(string(body)))
			w := httptest.NewRecorder()

			suite.repo.EXPECT().CreateURL(
				gomock.Any(),
				"c489a87f9b3b",
				"https://duckduckgo.com",
				"",
			).Return(custom_errs.ErrUniqueViolation).Times(1)
			suite.repo.EXPECT().GetURLByShortURL(gomock.Any(), "c489a87f9b3b").Return(&repository.URL{
				ID:          1,
				ShortURL:    "c489a87f9b3b",
				OriginalURL: "https://duckduckgo.com",
			}, nil).Times(1)

			suite.controller.GetServeMux().ServeHTTP(w, r)

			result := w.Result()
			defer func() {
				if err := result.Body.Close(); err != nil {
					fmt.Println("Body.Close:", err)
				}
			}()

			assert.Equal(t, test.expected.code, result.StatusCode, "unexpected status code")

			assert.Equal(t, test.expected.contentType, result.Header.Get("Content-Type"), "unexpected content type")

			resBody, err := io.ReadAll(result.Body)
			require.NoError(t, err, "error reading response body")
			res := &public.GetShortenURLResponse{}
			_ = json.Unmarshal(resBody, res)
			assert.Equal(t, test.expected.response, res, "unexpected response body")
		})
	}
}

func (suite *EndpointsTestSuite) Test_GetShortenURLHandler_BadRequestError(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   interface{}
	}

	type Expected struct {
		code        int
		response    *public.GetShortenURLResponse
		contentType string
	}
	tests := []struct {
		name     string
		request  Request
		expected Expected
	}{
		{
			name: "Test Get Shorten URL bad request error",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/api/shorten",
				body: &public.GetShortenURLRequest{
					URL: "",
				},
			},
			expected: Expected{
				code: http.StatusBadRequest,
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
			defer func() {
				if err := result.Body.Close(); err != nil {
					fmt.Println("Body.Close:", err)
				}
			}()

			assert.Equal(t, test.expected.code, result.StatusCode, "unexpected status code")

			assert.Equal(t, test.expected.contentType, result.Header.Get("Content-Type"), "unexpected content type")
		})
	}
}

func (suite *EndpointsTestSuite) Test_GetShortenURLHandler_CreateUrlError(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   interface{}
	}

	type Expected struct {
		code        int
		response    *public.GetShortenURLResponse
		contentType string
	}
	tests := []struct {
		name     string
		request  Request
		expected Expected
	}{
		{
			name: "Test Get Shorten URL create url error",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/api/shorten",
				body: &public.GetShortenURLRequest{
					URL: "https://duckduckgo.com",
				},
			},
			expected: Expected{
				code: http.StatusInternalServerError,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.request.body)
			r := httptest.NewRequest(test.request.method, test.request.url, strings.NewReader(string(body)))
			w := httptest.NewRecorder()

			suite.repo.EXPECT().CreateURL(
				gomock.Any(),
				"c489a87f9b3b",
				"https://duckduckgo.com",
				"",
			).Return(errors.New("some saving error")).Times(1)

			suite.controller.GetServeMux().ServeHTTP(w, r)

			result := w.Result()
			defer func() {
				if err := result.Body.Close(); err != nil {
					fmt.Println("Body.Close:", err)
				}
			}()

			assert.Equal(t, test.expected.code, result.StatusCode, "unexpected status code")

			assert.Equal(t, test.expected.contentType, result.Header.Get("Content-Type"), "unexpected content type")
		})
	}
}

func (suite *EndpointsTestSuite) Test_GetShortenURLHandler_UnmarshallError(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   interface{}
	}

	type Expected struct {
		code        int
		response    *public.GetShortenURLResponse
		contentType string
	}
	tests := []struct {
		name     string
		request  Request
		expected Expected
	}{
		{
			name: "Test Get Shorten URL unmarshall error",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/api/shorten",
				body: &public.GetShortenURLRequest{
					URL: "https://duckduckgo.com",
				},
			},
			expected: Expected{
				code: http.StatusInternalServerError,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := []byte(`{"url": "http://example.com"`)
			r := httptest.NewRequest(test.request.method, test.request.url, strings.NewReader(string(request)))
			w := httptest.NewRecorder()

			suite.controller.GetServeMux().ServeHTTP(w, r)

			result := w.Result()
			defer func() {
				if err := result.Body.Close(); err != nil {
					fmt.Println("Body.Close:", err)
				}
			}()

			assert.Equal(t, test.expected.code, result.StatusCode, "unexpected status code")

			assert.Equal(t, test.expected.contentType, result.Header.Get("Content-Type"), "unexpected content type")
		})
	}
}
