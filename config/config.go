// config - пакет с настройками микросервиса
package config

import (
	"flag"
	"os"
)

// Config - структура с переменными, в которые будут записаны конфиги
type Config struct {
	HTTPAddress        string // адрес, на котором будет доступен микросервис по сети
	BaseShortURL       string // базовый URL
	FileStoragePath    string // путь к файлу хранилища
	DNS                string // строка подключения к БД
	HashSecret         string // хэш секрет
	AuthUserCookieName string // имя куки
}

var cfg Config

// Init функция инициализирующая переменные, хранящие в себе конфигурационные параметры
func Init() Config {
	flag.StringVar(&cfg.HTTPAddress, "a", "localhost:8080", "HTTP address")
	flag.StringVar(&cfg.BaseShortURL, "b", "http://localhost:8080", "base short url")
	flag.StringVar(&cfg.FileStoragePath, "f", "./urls.txt", "file path")
	flag.StringVar(&cfg.DNS, "d", "user=postgres password=postgres dbname=praktikum host=localhost port=5432 sslmode=disable", "db connection")
	flag.StringVar(&cfg.HashSecret, "h", "my_secret", "hash secret")
	flag.StringVar(&cfg.AuthUserCookieName, "c", "auth_cookie", "auth cookie name")

	flag.Parse()

	if httpAddress := os.Getenv("SERVER_ADDRESS"); httpAddress != "" {
		cfg.HTTPAddress = httpAddress
	}

	if baseShortURL := os.Getenv("BASE_URL"); baseShortURL != "" {
		cfg.BaseShortURL = baseShortURL
	}

	if fileStoragePath := os.Getenv("FILE_STORAGE_PATH"); fileStoragePath != "" {
		cfg.FileStoragePath = fileStoragePath
	}

	if dns := os.Getenv("DSN"); dns != "" {
		cfg.DNS = dns
	}

	if hashSecret := os.Getenv("HASH_SECRET"); hashSecret != "" {
		cfg.HashSecret = hashSecret
	}

	if authUserCookieName := os.Getenv("AUTH_USER_COOKIE_NAME"); authUserCookieName != "" {
		cfg.AuthUserCookieName = authUserCookieName
	}

	return cfg
}
