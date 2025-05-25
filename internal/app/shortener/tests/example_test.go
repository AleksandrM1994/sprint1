package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/golang/mock/gomock"

	"github.com/sprint1/internal/app/shortener/endpoints"
	"github.com/sprint1/internal/app/shortener/repository"
)

// Example_CreateUserHandler показывает, как использовать CreateUserHandler для создания пользователя.
func (suite *EndpointsTestSuite) Example_CreateUserHandler() {
	body := &endpoints.CreateUserRequest{
		Login: "amakarkin",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("Failed to marshal request body: %v", err)
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
	defer func() {
		if err := result.Body.Close(); err != nil {
			fmt.Println("Body.Close:", err)
		}
	}()
	if result.StatusCode != http.StatusCreated {
		fmt.Printf("Expected status code %d, got %d", http.StatusCreated, result.StatusCode)
	}

	var response endpoints.CreateUserResponse
	if err := json.NewDecoder(result.Body).Decode(&response); err != nil {
		fmt.Printf("Failed to decode response body: %v", err)
	}

	fmt.Println(response)
}

// Example_AuthUserHandler показывает, как использовать AuthUserHandler для аутентификации пользователя.
func (suite *EndpointsTestSuite) Example_AuthUserHandler() {
	// Создаем тестовый запрос
	authRequest := &endpoints.AuthUserRequest{
		Login:    "test",
		Password: "test",
	}
	body, _ := json.Marshal(authRequest)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/user/auth", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	// Настраиваем мок репозитория
	suite.repo.EXPECT().GetUser(gomock.Any(), "test", "test").Return(&repository.User{
		ID:       "user_id",
		Login:    "test",
		Password: "test",
	}, nil).Times(1)

	// Вызываем обработчик
	suite.controller.GetServeMux().ServeHTTP(w, r)

	// Проверяем результат
	result := w.Result()
	defer result.Body.Close()

	// Проверяем статус код
	if result.StatusCode != http.StatusOK {
		panic("unexpected status code")
	}

	fmt.Println(result.Body)
}

// ExampleGetOriginalUrlHandler показывает, как использовать GetOriginalUrlHandler для получения оригинального URL.
func (suite *EndpointsTestSuite) ExampleGetOriginalUrlHandler() {
	// Создаем тестовый запрос
	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/aHR0cHM6Ly9qc29uZm9ybWF0dGVyLm9yZw==", nil)
	w := httptest.NewRecorder()

	// Настраиваем мок репозитория
	suite.repo.EXPECT().GetURLByShortURL(gomock.Any(), "aHR0cHM6Ly9qc29uZm9ybWF0dGVyLm9yZw==").Return(&repository.URL{
		ID:          1,
		ShortURL:    "aHR0cHM6Ly9qc29uZm9ybWF0dGVyLm9yZw==",
		OriginalURL: "https://jsonformatter.org",
	}, nil).Times(1)

	// Вызываем обработчик
	suite.controller.GetServeMux().ServeHTTP(w, r)

	// Проверяем результат
	result := w.Result()
	defer result.Body.Close()

	// Проверяем статус код
	if result.StatusCode != http.StatusTemporaryRedirect {
		fmt.Println("unexpected status code")
	}

	// Проверяем заголовок Location
	location := result.Header.Get("Location")
	if location != "https://jsonformatter.org" {
		fmt.Println("unexpected location")
	}

	fmt.Println(location)
}

// ExampleGetShortenURLHandler показывает, как использовать GetShortenURLHandler для получения сокращенного URL.
func (suite *EndpointsTestSuite) ExampleGetShortenURLHandler() {
	// Создаем тестовый запрос
	requestBody := &endpoints.GetShortenURLRequest{
		URL: "https://duckduckgo.com",
	}
	body, _ := json.Marshal(requestBody)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/shorten", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	// Настраиваем мок репозитория
	suite.repo.EXPECT().CreateURL(gomock.Any(), "c489a87f9b3b", "https://duckduckgo.com", "").Return(nil).Times(1)
	suite.repo.EXPECT().GetURLByShortURL(gomock.Any(), "c489a87f9b3b").Return(&repository.URL{
		ID:          1,
		ShortURL:    "c489a87f9b3b",
		OriginalURL: "https://duckduckgo.com",
	}, nil).Times(1)

	// Вызываем обработчик
	suite.controller.GetServeMux().ServeHTTP(w, r)

	// Проверяем результат
	result := w.Result()
	defer result.Body.Close()

	// Проверяем статус код
	if result.StatusCode != http.StatusCreated {
		fmt.Println("unexpected status code")
	}

	// Проверяем заголовок Content-Type
	if result.Header.Get("Content-Type") != "application/json" {
		fmt.Println("unexpected content type")
	}

	// Читаем тело ответа
	resBody, err := io.ReadAll(result.Body)
	if err != nil {
		fmt.Println("error reading response body")
	}

	// Проверяем ответ
	response := &endpoints.GetShortenURLResponse{}
	if err := json.Unmarshal(resBody, response); err != nil {
		fmt.Println("error unmarshalling response")
	}

	expectedResponse := &endpoints.GetShortenURLResponse{
		Result: "http://localhost:8080/c489a87f9b3b",
	}
	if response.Result != expectedResponse.Result {
		fmt.Println("unexpected response body")
	}

	fmt.Println(response.Result)
}
