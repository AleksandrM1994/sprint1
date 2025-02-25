package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (rw *ResponseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logging(lg *zap.SugaredLogger, next http.Handler) http.Handler {
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
