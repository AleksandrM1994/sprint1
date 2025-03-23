package endpoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sprint1/internal/app/shortener/repository"
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
					URL: "https://duckduckgo.com",
				},
			},
			expected: Expected{
				code: http.StatusCreated,
				response: &GetShortenURLResponse{
					Result: "http://localhost:8080/c489a87f9b3b",
				},
				contentType: "application/json",
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

			r.AddCookie(&http.Cookie{
				Name:  "auth_cookie",
				Value: "MTc0Mjc2MzgyMnxKN3VTYTkyYmwzc05tYURNNzFDRFFDT3JKakxxRWRsNnJtckZrV3N6R3dCcXk4anptaWxLOV91cHRsUzc0Z2xkamZTbzdfbjNMQ2s9fNEVcpB5EfxIKduWXSW_wvOyM0TWw2k7yV9uIF8qq5K3",
			})

			//cookieFinish := time.Now().AddDate(99, 0, 0)
			//suite.repo.EXPECT().GetUserByID(gomock.Any(), "b371d94a-78d2-4b8d-a5d4-d90e519b42cc").Return(
			//	&repository.User{
			//		ID:           "b371d94a-78d2-4b8d-a5d4-d90e519b42cc",
			//		Login:        "b371d94a",
			//		Password:     "b371d94a",
			//		Cookie:       "MTc0Mjc2MzgyMnxKN3VTYTkyYmwzc05tYURNNzFDRFFDT3JKakxxRWRsNnJtckZrV3N6R3dCcXk4anptaWxLOV91cHRsUzc0Z2xkamZTbzdfbjNMQ2s9fNEVcpB5EfxIKduWXSW_wvOyM0TWw2k7yV9uIF8qq5K3",
			//		CookieFinish: &cookieFinish,
			//	}, nil).MaxTimes(1)
			suite.repo.EXPECT().CreateURL(
				gomock.Any(),
				"c489a87f9b3b",
				"https://duckduckgo.com",
			).Return(nil).MaxTimes(1)
			suite.repo.EXPECT().GetURLByShortURL(gomock.Any(), "c489a87f9b3b").Return(&repository.URL{
				ID:          1,
				ShortURL:    "c489a87f9b3b",
				OriginalURL: "https://duckduckgo.com",
			}, nil).MaxTimes(1)

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
				res := &GetShortenURLResponse{}
				_ = json.Unmarshal(resBody, res)
				assert.Equal(t, test.expected.response, res, "unexpected response body")
			}
		})
	}
}
