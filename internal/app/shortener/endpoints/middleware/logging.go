package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/sprint1/internal/app/shortener/service"
)

// ResponseWriterWrapper - структура для логирования запросов/ответов
type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
	size       int
}

// WriteHeader функция для записи в ответ http кода
func (rw *ResponseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logging - мидлваря, осуществляющая логирование запросов и ответов по ним
func Logging(lg *zap.SugaredLogger, s *service.ServiceImpl, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &ResponseWriterWrapper{ResponseWriter: w}
		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		lg.Infoln(
			"uri", r.RequestURI,
			"method", r.Method,
			"status", rw.statusCode,
			"duration", duration,
			"size", rw.size,
		)
	})
}
