package main

import (
	"net/http"

	"github.com/sprint1/internal/app/shortener/endpoints"
	"github.com/sprint1/internal/app/shortener/service"
)

const defaultURL = "write original url, its default"

func main() {
	serviceImpl := service.NewService(defaultURL)
	mux := http.NewServeMux()
	controller := endpoints.NewController(mux, serviceImpl)
	err := http.ListenAndServe(":8080", controller.GetServeMux())
	if err != nil {
		panic(err)
	}
}
