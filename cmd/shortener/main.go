package main

import (
	"net/http"

	"github.com/sprint1/internal/app/shortener/endpoints"
	"github.com/sprint1/internal/app/shortener/service"

	"github.com/gorilla/mux"
)

func main() {
	serviceImpl := service.NewService()
	router := mux.NewRouter()
	controller := endpoints.NewController(router, serviceImpl)
	err := http.ListenAndServe(":8080", controller.GetServeMux())
	if err != nil {
		panic(err)
	}
}
