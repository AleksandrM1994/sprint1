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

	if HTTPAddress := os.Getenv("SERVER_ADDRESS"); HTTPAddress != "" {
		cfg.HTTPAddress = HTTPAddress
	}

	if BaseShortURL := os.Getenv("BASE_URL"); BaseShortURL != "" {
		cfg.BaseShortURL = BaseShortURL
	}

	return cfg
}
