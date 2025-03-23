package service

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
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
