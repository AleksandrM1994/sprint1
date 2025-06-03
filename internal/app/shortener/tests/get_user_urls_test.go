package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sprint1/internal/app/shortener/endpoints"
	custom_errs "github.com/sprint1/internal/app/shortener/errors"
	"github.com/sprint1/internal/app/shortener/repository"
)

func (suite *EndpointsTestSuite) Test_GetUserUrlsHandler_Success(t *testing.T) {
	type Request struct {
		method string
		url    string
	}

	type Expected struct {
		code        int
		contentType string
		response    []*endpoints.GetUserURLsResponse
	}
	tests := []struct {
		name     string
		request  Request
		expected Expected
	}{
		{
			name: "Test get user urls successfully",
			request: Request{
				method: http.MethodGet,
				url:    "http://localhost:8080/api/user/urls",
			},
			expected: Expected{
				code:        http.StatusOK,
				contentType: "application/json",
				response: []*endpoints.GetUserURLsResponse{
					{
						OriginalURL: "https://go.dev",
						ShortURL:    "http://localhost:8080/6e7f58f6b868",
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.request.method, test.request.url, nil)
			w := httptest.NewRecorder()

			r.AddCookie(&http.Cookie{
				Name:  "auth_cookie",
				Value: "MTc0Nzg5MzQzNHxLeWlQV2k2bjBYNm03cEZ5bXl4cFhIRjlzbkJlVE1reTloekRLZWpReHRrbzYtcTA2SXNwUWMyeXJfMC1Zd1luMUh4cEdxb195alE9fFkjI2E35AFwLWjCd8SQZwuTHMeDVVAq4nSrizBe9Xyd",
			})

			cookieFinish := time.Now().AddDate(99, 0, 0)
			suite.repo.EXPECT().GetUserByID(gomock.Any(), "39e529f3-7947-4d3b-aee4-d49a3a757c0f").Return(
				&repository.User{
					ID:           "39e529f3-7947-4d3b-aee4-d49a3a757c0f",
					Login:        "Ho8bmla6ULoW2wIaJj0jjj5wKXh/Wtbl5IUmKXaW/3U=",
					Password:     "Ho8bmla6ULoW2wIaJj0jjj5wKXh/Wtbl5IUmKXaW/3U=",
					Cookie:       "MTc0Nzg5MzQzNHxLeWlQV2k2bjBYNm03cEZ5bXl4cFhIRjlzbkJlVE1reTloekRLZWpReHRrbzYtcTA2SXNwUWMyeXJfMC1Zd1luMUh4cEdxb195alE9fFkjI2E35AFwLWjCd8SQZwuTHMeDVVAq4nSrizBe9Xyd",
					CookieFinish: &cookieFinish,
				}, nil).Times(1)

			suite.repo.EXPECT().GetURLsByUserID(
				gomock.Any(),
				"39e529f3-7947-4d3b-aee4-d49a3a757c0f",
			).Return([]*repository.URL{
				{
					ID:          1,
					OriginalURL: "https://go.dev",
					ShortURL:    "http://localhost:8080/6e7f58f6b868",
					UserID:      "39e529f3-7947-4d3b-aee4-d49a3a757c0f",
					IsDeleted:   false,
				},
			}, nil).Times(1)

			suite.controller.GetServeMux().ServeHTTP(w, r)

			result := w.Result()
			defer func() {
				if err := result.Body.Close(); err != nil {
					fmt.Println("Body.Close:", err)
				}
			}()

			assert.Equal(t, test.expected.code, result.StatusCode, "unexpected status code")

			if result.StatusCode == http.StatusCreated {
				resBody, err := io.ReadAll(result.Body)
				require.NoError(t, err, "error reading response body")
				var res []*endpoints.GetUserURLsResponse
				_ = json.Unmarshal(resBody, &res)
				assert.Equal(t, test.expected.response, res, "unexpected response body")
			}
		})
	}
}

func (suite *EndpointsTestSuite) Test_GetUserUrlsHandler_NoContent(t *testing.T) {
	type Request struct {
		method string
		url    string
	}

	type Expected struct {
		code        int
		contentType string
		response    []*endpoints.GetUserURLsResponse
	}
	tests := []struct {
		name     string
		request  Request
		expected Expected
	}{
		{
			name: "Test get user urls error, no content",
			request: Request{
				method: http.MethodGet,
				url:    "http://localhost:8080/api/user/urls",
			},
			expected: Expected{
				code: http.StatusNoContent,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.request.method, test.request.url, nil)
			w := httptest.NewRecorder()

			r.AddCookie(&http.Cookie{
				Name:  "auth_cookie",
				Value: "MTc0Nzg5MzQzNHxLeWlQV2k2bjBYNm03cEZ5bXl4cFhIRjlzbkJlVE1reTloekRLZWpReHRrbzYtcTA2SXNwUWMyeXJfMC1Zd1luMUh4cEdxb195alE9fFkjI2E35AFwLWjCd8SQZwuTHMeDVVAq4nSrizBe9Xyd",
			})

			cookieFinish := time.Now().AddDate(99, 0, 0)
			suite.repo.EXPECT().GetUserByID(gomock.Any(), "39e529f3-7947-4d3b-aee4-d49a3a757c0f").Return(
				&repository.User{
					ID:           "39e529f3-7947-4d3b-aee4-d49a3a757c0f",
					Login:        "Ho8bmla6ULoW2wIaJj0jjj5wKXh/Wtbl5IUmKXaW/3U=",
					Password:     "Ho8bmla6ULoW2wIaJj0jjj5wKXh/Wtbl5IUmKXaW/3U=",
					Cookie:       "MTc0Nzg5MzQzNHxLeWlQV2k2bjBYNm03cEZ5bXl4cFhIRjlzbkJlVE1reTloekRLZWpReHRrbzYtcTA2SXNwUWMyeXJfMC1Zd1luMUh4cEdxb195alE9fFkjI2E35AFwLWjCd8SQZwuTHMeDVVAq4nSrizBe9Xyd",
					CookieFinish: &cookieFinish,
				}, nil).Times(1)

			suite.repo.EXPECT().GetURLsByUserID(
				gomock.Any(),
				"39e529f3-7947-4d3b-aee4-d49a3a757c0f",
			).Return(nil, custom_errs.ErrNoContent).Times(1)

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
