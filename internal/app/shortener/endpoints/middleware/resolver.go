package middleware

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/zap"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/service"
)

// Resolver - мидлваря, которая проверяет с разрешенного ли айпи пришел запрос
func Resolver(cfg config.Config, lg *zap.SugaredLogger, s *service.ServiceImpl, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, err := resolveIP(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, ipNet, err := net.ParseCIDR(cfg.TrustedSubnet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !ipNet.Contains(ip) {
			w.WriteHeader(http.StatusForbidden)
			http.Error(w, errors.New("the client address is not in the trusted subnet").Error(), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func resolveIP(r *http.Request) (net.IP, error) {
	// смотрим заголовок запроса X-Real-IP
	ipStr := r.Header.Get("X-Real-IP")
	// парсим ip
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, fmt.Errorf("failed parse ip from http header")
	}
	return ip, nil
}
