package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sprint1/internal/app/shortener/endpoints"
)

func (suite *EndpointsTestSuite) Example_CreateUserHandler() {
	body := &endpoints.CreateUserRequest{
		Login: "amakarkin",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/user/create", strings.NewReader(string(bodyBytes)))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	suite.repo.EXPECT().CreateUser(
		gomock.Any(),
		gomock.Any(),
		"2/dkyU+GGwHFGIZE5261413Wy2lmdg5gHf6tq+sT87c=",
		gomock.Any(),
	).Return(nil).Times(1)

	suite.controller.GetServeMux().ServeHTTP(recorder, req)

	result := recorder.Result()
	if result.StatusCode != http.StatusCreated {
		log.Fatalf("Expected status code %d, got %d", http.StatusCreated, result.StatusCode)
	}

	var response endpoints.CreateUserResponse
	if err := json.NewDecoder(result.Body).Decode(&response); err != nil {
		log.Fatalf("Failed to decode response body: %v", err)
	}

	log.Println(response)
}

func (suite *EndpointsTestSuite) Test_CreateUserHandler(t *testing.T) {
	type Request struct {
		method string
		url    string
		body   *endpoints.CreateUserRequest
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
				body: &endpoints.CreateUserRequest{
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
