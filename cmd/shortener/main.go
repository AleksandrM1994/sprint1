package main

import (
	"net/http"

	"github.com/sprint1/internal/app/shortener/endpoints"
	"github.com/sprint1/internal/app/shortener/service"
)

func main() {
	serviceImpl := service.NewService()
	mux := http.NewServeMux()
	controller := endpoints.NewController(mux, serviceImpl)
	err := http.ListenAndServe(":8080", controller.GetServeMux())
	if err != nil {
		panic(err)
	}
}
