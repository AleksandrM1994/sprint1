package endpoints

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// ResponseWriterWrapper оборачивает http.ResponseWriter
type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
	size       int
}

// WriteHeader перехватывает вызов метода WriteHeader
func (rw *ResponseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func WithLogging(next http.Handler, lg *zap.SugaredLogger) http.Handler {
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
