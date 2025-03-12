package config

import (
	"flag"
	"os"
)

type Config struct {
	HTTPAddress     string
	BaseShortURL    string
	FileStoragePath string
	DNS             string
}

var cfg Config

func Init() Config {
	flag.StringVar(&cfg.HTTPAddress, "a", "localhost:8080", "HTTP address")
	flag.StringVar(&cfg.BaseShortURL, "b", "http://localhost:8080", "base short url")
	flag.StringVar(&cfg.FileStoragePath, "f", "./urls.txt", "file path")

	flag.StringVar(&cfg.DNS, "d", "user=postgres password=postgres dbname=praktikum host=postgres port=5432 sslmode=disable", "db connection")

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

	if dns := os.Getenv("DNS"); dns != "" {
		cfg.DNS = dns
	}

	return cfg
}
