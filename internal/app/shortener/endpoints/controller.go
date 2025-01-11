package endpoints

import (
	"net/http"

	"github.com/sprint1/internal/app/shortener/service"
)

type Controller struct {
	mux     *http.ServeMux
	service service.Service
}

func NewController(mux *http.ServeMux, service service.Service) *Controller {
	controller := &Controller{mux: mux, service: service}
	controller.InitHandlers(mux)
	return controller
}

func (c *Controller) InitHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c.UrlWorkerHandlerHandler(w, r)
	})
}

func (c *Controller) GetServeMux() *http.ServeMux {
	return c.mux
}
