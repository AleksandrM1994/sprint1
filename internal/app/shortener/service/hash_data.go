package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// HashData сервисная функция по хэшированию данных
func (s *ServiceImpl) HashData(data []byte) (string, error) {
	hash := hmac.New(sha256.New, []byte(s.cfg.HashSecret))
	_, err := hash.Write(data)
	if err != nil {
		return "", fmt.Errorf("error hashing data: %v", err)
	}
	sum := hash.Sum(nil)
	hashCode := base64.StdEncoding.EncodeToString(sum)
	return hashCode, nil
}
