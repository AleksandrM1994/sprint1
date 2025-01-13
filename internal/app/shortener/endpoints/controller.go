package endpoints

import (
	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/service"

	"github.com/gorilla/mux"
)

type Controller struct {
	router  *mux.Router
	service service.Service
	cfg     config.Config
}

func NewController(router *mux.Router, service service.Service, cfg config.Config) *Controller {
	controller := &Controller{router: router, service: service, cfg: cfg}
	controller.InitHandlers()
	return controller
}

func (c *Controller) InitHandlers() {
	c.router.HandleFunc("/", c.SaveURLHandler).Methods("POST")
	c.router.HandleFunc("/{id}", c.GetOriginalURLHandler).Methods("GET")
}

func (c *Controller) GetServeMux() *mux.Router {
	return c.router
}
