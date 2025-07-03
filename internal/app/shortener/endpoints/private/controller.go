package private

import (
	"go.uber.org/zap"

	"github.com/sprint1/config"
	pb "github.com/sprint1/internal/app/shortener/endpoints/private/proto"
	"github.com/sprint1/internal/app/shortener/service"
)

// Controller структура котроллера
type Controller struct {
	service *service.ServiceImpl
	cfg     config.Config
	lg      *zap.SugaredLogger
	pb.UnimplementedShortenerServer
}

// NewController функция по созданию контроллера
func NewController(service *service.ServiceImpl, cfg config.Config, lg *zap.SugaredLogger) *Controller {
	controller := &Controller{service: service, cfg: cfg, lg: lg}
	return controller
}
