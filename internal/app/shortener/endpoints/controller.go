package endpoints

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/endpoints/middleware"
	"github.com/sprint1/internal/app/shortener/service"

	"github.com/gorilla/mux"
)

type Middleware func(*zap.SugaredLogger, http.Handler) http.Handler

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
	c.router.Handle("/", applyMiddlewares(http.HandlerFunc(c.SaveURLHandler), c.lg, middleware.Logging)).Methods("POST")
	c.router.Handle("/{id}", applyMiddlewares(http.HandlerFunc(c.GetOriginalURLHandler), c.lg, middleware.Logging)).Methods("GET")
	c.router.Handle("/api/shorten", applyMiddlewares(http.HandlerFunc(c.GetShortenURLHandler), c.lg, middleware.Logging, middleware.GzipMiddleware)).Methods("POST")
}

func (c *Controller) GetServeMux() *mux.Router {
	return c.router
}

func applyMiddlewares(h http.Handler, lg *zap.SugaredLogger, middlewares ...Middleware) http.Handler {
	for _, mw := range middlewares {
		h = mw(lg, h)
	}
	return h
}
