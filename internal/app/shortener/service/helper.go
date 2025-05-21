package service

import (
	"crypto/sha256"
	"encoding/hex"

	"time"

	"github.com/sprint1/internal/app/shortener/helpers"
)

func HashString(s string) string {
	hash := sha256.New()

	hash.Write([]byte(s))

	hashInBytes := hash.Sum(nil)

	return hex.EncodeToString(hashInBytes)
}

func DatePtr(date time.Time) *time.Time {
	return &date
}

func CreateShortURL(url string) string {
	url = helpers.RemoveControlCharacters(url)
	hashURL := HashString(url)
	fifthLength := len(hashURL) / 5

	// Обрезаем hashURL до нужной длины
	shortURL := hashURL[:fifthLength]
	return shortURL
}
