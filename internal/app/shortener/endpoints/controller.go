package endpoints

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/endpoints/middleware"
	"github.com/sprint1/internal/app/shortener/service"

	"github.com/gorilla/mux"
)

// Middleware шаблон мидлвари
type Middleware func(*zap.SugaredLogger, *service.ServiceImpl, http.Handler) http.Handler

// Controller структура котроллера
type Controller struct {
	router  *mux.Router
	service *service.ServiceImpl
	cfg     config.Config
	lg      *zap.SugaredLogger
}

// NewController функция по созданию контроллера
func NewController(router *mux.Router, service *service.ServiceImpl, cfg config.Config, lg *zap.SugaredLogger) *Controller {
	controller := &Controller{router: router, service: service, cfg: cfg, lg: lg}
	controller.InitHandlers()
	return controller
}

// InitHandlers - описание эндпоинтов, доступных для вызова
func (c *Controller) InitHandlers() {
	c.router.Handle(
		"/",
		applyMiddlewares(
			http.HandlerFunc(c.SaveURLHandler),
			c.lg,
			c.service,
			middleware.Logging,
			middleware.GzipMiddleware,
			middleware.Authenticate,
		),
	).Methods("POST")
	c.router.Handle(
		"/ping",
		applyMiddlewares(
			http.HandlerFunc(c.PingHandler),
			c.lg,
			c.service,
			middleware.Logging,
		),
	).Methods("GET")
	c.router.Handle("/{id}",
		applyMiddlewares(
			http.HandlerFunc(c.GetOriginalURLHandler),
			c.lg,
			c.service,
			middleware.Logging,
		),
	).Methods("GET")
	c.router.Handle(
		"/api/shorten",
		applyMiddlewares(
			http.HandlerFunc(c.GetShortenURLHandler),
			c.lg,
			c.service,
			middleware.Logging,
			middleware.GzipMiddleware,
		),
	).Methods("POST")
	c.router.Handle(
		"/api/shorten/batch",
		applyMiddlewares(
			http.HandlerFunc(c.SaveURLsBatch),
			c.lg,
			c.service,
			middleware.Logging,
			middleware.Authenticate,
		),
	).Methods("POST")
	c.router.Handle(
		"/api/user/create",
		applyMiddlewares(
			http.HandlerFunc(c.CreateUser),
			c.lg,
			c.service,
			middleware.Logging,
		),
	).Methods("POST")
	c.router.Handle(
		"/api/user/auth",
		applyMiddlewares(
			http.HandlerFunc(c.AuthUser),
			c.lg,
			c.service,
			middleware.Logging,
			middleware.Authenticate,
		),
	).Methods("POST")
	c.router.Handle(
		"/api/user/urls",
		applyMiddlewares(
			http.HandlerFunc(c.GetUserURLs),
			c.lg,
			c.service,
			middleware.Logging,
			middleware.Authenticate,
		),
	).Methods("GET")
	c.router.Handle(
		"/api/user/urls",
		applyMiddlewares(
			http.HandlerFunc(c.DeleteUserURLs),
			c.lg,
			c.service,
			middleware.Logging,
			middleware.Authenticate,
		),
	).Methods("DELETE")
}

// GetServeMux - возвращает экземпляр *mux.Router
func (c *Controller) GetServeMux() *mux.Router {
	return c.router
}

// applyMiddlewares применяет указанные мидлвари
func applyMiddlewares(h http.Handler, lg *zap.SugaredLogger, s *service.ServiceImpl, middlewares ...Middleware) http.Handler {
	for _, mw := range middlewares {
		h = mw(lg, s, h)
	}
	return h
}
