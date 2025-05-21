package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func (suite *EndpointsTestSuite) Test_CreateUserHandler(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   *CreateUserRequest
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
			name: "Test create user successfully",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/api/user/create",
				body: &CreateUserRequest{
					Login: "amakarkin",
				},
			},
			expected: Expected{
				code:        http.StatusCreated,
				contentType: "application/json",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.request.body)
			r := httptest.NewRequest(test.request.method, test.request.url, strings.NewReader(string(body)))
			w := httptest.NewRecorder()

			suite.repo.EXPECT().CreateUser(
				gomock.Any(),
				gomock.Any(),
				"2/dkyU+GGwHFGIZE5261413Wy2lmdg5gHf6tq+sT87c=",
				gomock.Any(),
			).Return(nil).MaxTimes(1)

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

func (suite *EndpointsTestSuite) Benchmark_CreateUserHandler(b *testing.B) {
	type Request struct {
		method string
		url    string
		body   *CreateUserRequest
	}

	// Подготовка запроса
	request := Request{
		method: http.MethodPost,
		url:    "http://localhost:8080/api/user/create",
		body: &CreateUserRequest{
			Login: "amakarkin",
		},
	}

	body, _ := json.Marshal(request.body)
	r := httptest.NewRequest(request.method, request.url, strings.NewReader(string(body)))
	w := httptest.NewRecorder()

	// Настройка ожиданий для моков
	suite.repo.EXPECT().CreateUser(
		gomock.Any(),
		gomock.Any(),
		"2/dkyU+GGwHFGIZE5261413Wy2lmdg5gHf6tq+sT87c=",
		gomock.Any(),
	).Return(nil).MaxTimes(1)

	// Запуск бенчмарка
	for i := 0; i < b.N; i++ {
		suite.controller.GetServeMux().ServeHTTP(w, r)
	}
}
