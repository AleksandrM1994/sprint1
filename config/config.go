package config

import (
	"flag"
	"os"
)

type Config struct {
	HTTPAddress  string
	BaseShortURL string
}

var cfg Config

func Init() Config {
	flag.StringVar(&cfg.HTTPAddress, "a", "localhost:8080", "HTTP address")
	flag.StringVar(&cfg.BaseShortURL, "b", "http://localhost:8080", "base short url")

	flag.Parse()

	if httpAddress := os.Getenv("SERVER_ADDRESS"); httpAddress != "" {
		cfg.HTTPAddress = httpAddress
	}

	if baseShortURL := os.Getenv("BASE_URL"); baseShortURL != "" {
		cfg.BaseShortURL = baseShortURL
	}

	return cfg
}
