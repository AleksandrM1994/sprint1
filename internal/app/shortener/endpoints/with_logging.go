package endpoints

import (
	"net/http"
	"time"
)

func (c *Controller) WithLogging(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	duration := time.Since(start)

	c.lg.Infoln(
		"uri", r.RequestURI,
		"method", r.Method,
		"status", responseData.status, // получаем перехваченный код статуса ответа
		"duration", duration,
		"size", responseData.size, // получаем перехваченный размер ответа
	)
}
