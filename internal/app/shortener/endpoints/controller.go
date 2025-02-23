package endpoints

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/service"

	"github.com/gorilla/mux"
)

type Controller struct {
	router  *mux.Router
	service service.Service
	cfg     config.Config
	lg      *zap.SugaredLogger
	http.Handler
}

func NewController(router *mux.Router, service service.Service, cfg config.Config, lg *zap.SugaredLogger) *Controller {
	controller := &Controller{router: router, service: service, cfg: cfg, lg: lg}
	controller.InitHandlers()
	return controller
}

func (c *Controller) InitHandlers() {
	c.router.HandleFunc("/", c.SaveURLHandler).Methods("POST")
	c.router.HandleFunc("/{id}", c.GetOriginalURLHandler).Methods("GET")
	c.router.HandleFunc("/api/shorten", c.GetShortenURLHandler).Methods("POST")
}

func (c *Controller) GetServeMux() *mux.Router {
	return c.router
}
