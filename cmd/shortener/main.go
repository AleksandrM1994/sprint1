// пакет, который стартует микросервис, отвечающий за работу с сокращенными урлами
package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

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

	var server *http.Server
	if cfg.EnableHTTPS {
		tlsConfig := getTLSConfig(lg)

		// Создаем новый сервер с TLS конфигурацией
		server = &http.Server{
			Addr:      ":443",
			Handler:   controller.GetServeMux(),
			TLSConfig: tlsConfig,
		}

		errListenAndServeTLS := server.ListenAndServeTLS("", "")
		if errListenAndServeTLS != nil {
			lg.Fatal("server.ListenAndServeTLS:", errListenAndServeTLS)
		}
	} else {
		// Создаем новый HTTP сервер
		server = &http.Server{
			Addr:    cfg.HTTPAddress,
			Handler: controller.GetServeMux(),
		}
		errListenAndServe := server.ListenAndServe()
		if errListenAndServe != nil {
			lg.Fatal("server.ListenAndServe:", errListenAndServe)
		}
	}

	// Создание канала для перехвата сигналов
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// Ожидание сигнала
	sig := <-signalChan
	log.Printf("Получен сигнал: %s. Начинаем завершение работы...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Завершение работы сервера
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("server.Shutdown: %v", err)
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

func getTLSConfig(lg *zap.SugaredLogger) *tls.Config {
	// Укажите пути к файлам с сертификатом и приватным ключом
	certFilePath := "./certificate.pem"
	privateKeyFilePath := "./private_key.pem"

	// Чтение файла сертификата
	certPEM, err := os.ReadFile(certFilePath)
	if err != nil {
		log.Fatalf("Error reading certificate file: %v", err)
	}

	// Чтение файла приватного ключа
	privateKeyPEM, err := os.ReadFile(privateKeyFilePath)
	if err != nil {
		log.Fatalf("Error reading private key file: %v", err)
	}

	certificate, err := tls.X509KeyPair(certPEM, privateKeyPEM)
	if err != nil {
		lg.Fatal("Failed to parse certificate and key:", err)
	}

	// так сертификат тестовый, то по факту проверка фиктивная, так как нужен корневой сертификат
	return &tls.Config{
		Certificates:       []tls.Certificate{certificate},
		InsecureSkipVerify: true,
	}
}
