package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sprint1/internal/app/shortener/endpoints/public"
)

func (suite *EndpointsTestSuite) Test_CreateUserHandler_Success(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   *public.CreateUserRequest
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
				body: &public.CreateUserRequest{
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

func (suite *EndpointsTestSuite) Test_CreateUserHandler_CreateUserError(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   *public.CreateUserRequest
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
			name: "Test create user create user error",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/api/user/create",
				body: &public.CreateUserRequest{
					Login: "amakarkin",
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

			suite.repo.EXPECT().CreateUser(
				gomock.Any(),
				gomock.Any(),
				"2/dkyU+GGwHFGIZE5261413Wy2lmdg5gHf6tq+sT87c=",
				gomock.Any(),
			).Return(errors.New("some error")).Times(1)

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

func (suite *EndpointsTestSuite) Test_CreateUserHandler_BadRequestError(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   *public.CreateUserRequest
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
			name: "Test create user error, bad request",
			request: Request{
				method: http.MethodPost,
				url:    "http://localhost:8080/api/user/create",
				body: &public.CreateUserRequest{
					Login: "",
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
		})
	}
}
