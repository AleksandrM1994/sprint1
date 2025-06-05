// пакет, который стартует микросервис, отвечающий за работу с сокращенными урлами
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os/exec"
	"strings"

	"go.uber.org/zap"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/endpoints"
	"github.com/sprint1/internal/app/shortener/repository"
	"github.com/sprint1/internal/app/shortener/service"
	"github.com/sprint1/internal/app/shortener/workers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// глобальные переменные с информацией о сборке
var (
	buildVersion = "N/A" // версия микросервиса
	buildDate    = "N/A" // дата сборки
	buildCommit  = "N/A" // текст коммита текущей сборки
)

func runShortener() {
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

	workerPool := workers.NewWorkerPool(lg, repo)
	workerPool.Start()

	serviceImpl := service.NewService(lg, cfg, repo, workerPool)
	router := mux.NewRouter()
	controller := endpoints.NewController(router, serviceImpl, cfg, lg)

	errListenAndServe := http.ListenAndServe(cfg.HTTPAddress, controller.GetServeMux())
	if errListenAndServe != nil {
		lg.Fatal("http.ListenAndServe:", errListenAndServe)
	}
}

func main() {
	outTagVersion, err := exec.Command("git", "describe", "--tags").Output()
	if err != nil {
		fmt.Println(err)
	}
	buildVersion = strings.TrimSpace(string(outTagVersion))

	outCommitMessage, err := exec.Command("git", "log", "-1", "--pretty=format:%s").Output()
	if err != nil {
		fmt.Println(err)
	}
	buildCommit = string(outCommitMessage)

	outCommitDate, err := exec.Command("git", "log", "-1", "--pretty=format:%cd").Output()
	if err != nil {
		fmt.Println(err)
	}
	buildDate = string(outCommitDate)

	fmt.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n", buildVersion, buildDate, buildCommit)

	go runShortener()
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		return
	}
}
