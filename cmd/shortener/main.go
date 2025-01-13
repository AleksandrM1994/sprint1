package main

import (
	"net/http"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/endpoints"
	"github.com/sprint1/internal/app/shortener/service"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Init()

	serviceImpl := service.NewService()
	router := mux.NewRouter()
	controller := endpoints.NewController(router, serviceImpl, cfg)
	err := http.ListenAndServe(cfg.HTTPAddress, controller.GetServeMux())
	if err != nil {
		panic(err)
	}
}
