package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sprint1/internal/app/shortener/endpoints"
	"github.com/sprint1/internal/app/shortener/repository"
)

func (suite *EndpointsTestSuite) Test_AuthUserHandler(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   *endpoints.AuthUserRequest
	}

	type Expected struct {
		code        int
		contentType string
	}
	tests := []struct {
		name     string
		request  Request
		expected Expected
	}{
		{
			name: "Test auth user successfully",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/api/user/auth",
				body: &endpoints.AuthUserRequest{
					Login:    "test",
					Password: "test",
				},
			},
			expected: Expected{
				code:        http.StatusOK,
				contentType: "application/json",
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
				Value: "MTc0Nzg5MzQzNHxLeWlQV2k2bjBYNm03cEZ5bXl4cFhIRjlzbkJlVE1reTloekRLZWpReHRrbzYtcTA2SXNwUWMyeXJfMC1Zd1luMUh4cEdxb195alE9fFkjI2E35AFwLWjCd8SQZwuTHMeDVVAq4nSrizBe9Xyd",
			})

			cookieFinish := time.Now().AddDate(99, 0, 0)
			suite.repo.EXPECT().GetUser(
				gomock.Any(),
				"Ho8bmla6ULoW2wIaJj0jjj5wKXh/Wtbl5IUmKXaW/3U=",
				"Ho8bmla6ULoW2wIaJj0jjj5wKXh/Wtbl5IUmKXaW/3U=",
			).Return(&repository.User{
				ID:           "39e529f3-7947-4d3b-aee4-d49a3a757c0f",
				Login:        "Ho8bmla6ULoW2wIaJj0jjj5wKXh/Wtbl5IUmKXaW/3U=",
				Password:     "Ho8bmla6ULoW2wIaJj0jjj5wKXh/Wtbl5IUmKXaW/3U=",
				Cookie:       "MTc0Nzg5MzQzNHxLeWlQV2k2bjBYNm03cEZ5bXl4cFhIRjlzbkJlVE1reTloekRLZWpReHRrbzYtcTA2SXNwUWMyeXJfMC1Zd1luMUh4cEdxb195alE9fFkjI2E35AFwLWjCd8SQZwuTHMeDVVAq4nSrizBe9Xyd",
				CookieFinish: &cookieFinish,
			}, nil).Times(1)

			suite.repo.EXPECT().GetUserByID(gomock.Any(), "39e529f3-7947-4d3b-aee4-d49a3a757c0f").Return(
				&repository.User{
					ID:           "39e529f3-7947-4d3b-aee4-d49a3a757c0f",
					Login:        "Ho8bmla6ULoW2wIaJj0jjj5wKXh/Wtbl5IUmKXaW/3U=",
					Password:     "Ho8bmla6ULoW2wIaJj0jjj5wKXh/Wtbl5IUmKXaW/3U=",
					Cookie:       "MTc0Nzg5MzQzNHxLeWlQV2k2bjBYNm03cEZ5bXl4cFhIRjlzbkJlVE1reTloekRLZWpReHRrbzYtcTA2SXNwUWMyeXJfMC1Zd1luMUh4cEdxb195alE9fFkjI2E35AFwLWjCd8SQZwuTHMeDVVAq4nSrizBe9Xyd",
					CookieFinish: &cookieFinish,
				}, nil).Times(1)

			suite.repo.EXPECT().UpdateUser(
				gomock.Any(),
				gomock.Any(),
			).Return(nil).Times(1)

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
