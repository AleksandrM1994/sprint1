package main

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/endpoints"
	"github.com/sprint1/internal/app/shortener/repository"
	"github.com/sprint1/internal/app/shortener/service"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	logger, loggerErr := zap.NewDevelopment()
	if loggerErr != nil {
		panic("cannot initialize zap")
	}
	defer func() {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}()

	lg := logger.Sugar()

	cfg := config.Init()

	repo, errSelectRepo := repository.SelectRepo(lg, cfg)
	if errSelectRepo != nil {
		lg.Fatal("repository.SelectRepo:", errSelectRepo)
	}

	serviceImpl := service.NewService(lg, cfg, repo)
	router := mux.NewRouter()
	controller := endpoints.NewController(router, serviceImpl, cfg, lg)
	errListenAndServe := http.ListenAndServe(cfg.HTTPAddress, controller.GetServeMux())
	if errListenAndServe != nil {
		lg.Fatal("http.ListenAndServe:", errListenAndServe)
	}
}
