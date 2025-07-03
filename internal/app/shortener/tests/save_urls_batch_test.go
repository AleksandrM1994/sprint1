package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sprint1/internal/app/shortener/endpoints/public"
	"github.com/sprint1/internal/app/shortener/repository"
)

func (suite *EndpointsTestSuite) Test_SaveUrlsBatchHandler(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   []*public.URLInBatch
	}

	type Expected struct {
		code        int
		response    []*public.URLInBatch
		contentType string
	}
	tests := []struct {
		name     string
		request  Request
		expected Expected
	}{
		{
			name: "Test Save URLs batch successfully",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/api/shorten/batch",
				body: []*public.URLInBatch{
					{
						CorrelationID: "qwe123",
						OriginalURL:   "https://go.dev",
					},
					{
						CorrelationID: "asd123",
						OriginalURL:   "https://github.com",
					},
					{
						CorrelationID: "zxc123",
						OriginalURL:   "https://ya.ru",
					},
				},
			},
			expected: Expected{
				code: http.StatusCreated,
				response: []*public.URLInBatch{
					{
						CorrelationID: "qwe123",
						ShortURL:      "http://localhost:8080/6e7f58f6b868",
					},
					{
						CorrelationID: "asd123",
						ShortURL:      "http://localhost:8080/996e1f714b08",
					},
					{
						CorrelationID: "zxc123",
						ShortURL:      "http://localhost:8080/7e90a4f9c30b",
					},
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

			r.AddCookie(&http.Cookie{
				Name:  "auth_cookie",
				Value: "MTc0Nzg5MzQzNHxLeWlQV2k2bjBYNm03cEZ5bXl4cFhIRjlzbkJlVE1reTloekRLZWpReHRrbzYtcTA2SXNwUWMyeXJfMC1Zd1luMUh4cEdxb195alE9fFkjI2E35AFwLWjCd8SQZwuTHMeDVVAq4nSrizBe9Xyd",
			})

			cookieFinish := time.Now().AddDate(99, 0, 0)
			suite.repo.EXPECT().GetUserByID(gomock.Any(), "39e529f3-7947-4d3b-aee4-d49a3a757c0f").Return(
				&repository.User{
					ID:           "39e529f3-7947-4d3b-aee4-d49a3a757c0f",
					Login:        "b371d94a",
					Password:     "b371d94a",
					Cookie:       "MTc0Nzg5MzQzNHxLeWlQV2k2bjBYNm03cEZ5bXl4cFhIRjlzbkJlVE1reTloekRLZWpReHRrbzYtcTA2SXNwUWMyeXJfMC1Zd1luMUh4cEdxb195alE9fFkjI2E35AFwLWjCd8SQZwuTHMeDVVAq4nSrizBe9Xyd",
					CookieFinish: &cookieFinish,
				}, nil).MaxTimes(1)
			suite.repo.EXPECT().CreateURLs(
				gomock.Any(),
				[]*repository.URL{
					{
						ShortURL:    "6e7f58f6b868",
						OriginalURL: "https://go.dev",
						UserID:      "39e529f3-7947-4d3b-aee4-d49a3a757c0f",
					},
					{
						ShortURL:    "996e1f714b08",
						OriginalURL: "https://github.com",
						UserID:      "39e529f3-7947-4d3b-aee4-d49a3a757c0f",
					},
					{
						ShortURL:    "7e90a4f9c30b",
						OriginalURL: "https://ya.ru",
						UserID:      "39e529f3-7947-4d3b-aee4-d49a3a757c0f",
					},
				},
			).Return(nil).MaxTimes(1)

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
				var res []*public.URLInBatch
				_ = json.Unmarshal(resBody, &res)
				assert.Equal(t, test.expected.response, res, "unexpected response body")
			}
		})
	}
}
