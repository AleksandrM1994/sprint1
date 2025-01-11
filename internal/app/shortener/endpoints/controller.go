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
	controller.InitHandlers()
	return controller
}

func (c *Controller) InitHandlers() {
	c.mux.HandleFunc("/", c.SaveURLHandler)
	c.mux.HandleFunc("/url/", c.GetOriginalURLHandler)
}

func (c *Controller) GetServeMux() *http.ServeMux {
	return c.mux
}
