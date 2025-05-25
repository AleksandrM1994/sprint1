package endpoints

import (
	"context"
	"net/http"
	"time"
)

// PingHandler ручка для health check базы данных
func (c *Controller) PingHandler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
	defer cancel()

	err := c.service.Ping(ctx)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}
