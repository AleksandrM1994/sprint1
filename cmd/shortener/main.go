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
	defer logger.Sync()

	lg := logger.Sugar()

	cfg := config.Init()

	repo, errNewRepoImpl := repository.NewRepoImpl(lg, cfg)
	if errNewRepoImpl != nil {
		panic(errNewRepoImpl)
	}

	serviceImpl := service.NewService(lg, cfg, repo)
	router := mux.NewRouter()
	controller := endpoints.NewController(router, serviceImpl, cfg, lg)
	err := http.ListenAndServe(cfg.HTTPAddress, controller.GetServeMux())
	if err != nil {
		panic(err)
	}
}
