package endpoints

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/mocks"
	"github.com/sprint1/internal/app/shortener/service"
)

type EndpointsTestSuite struct {
	controller *Controller
	repo       *mocks.MockRepoDB
}

func TestEndpointSuite(t *testing.T) {
	router := mux.NewRouter()
	logger, loggerErr := zap.NewDevelopment()
	if loggerErr != nil {
		panic("cannot initialize zap")
	}
	lg := logger.Sugar()
	cfg := config.Init()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockRepoDB(ctrl)
	serviceImpl := service.NewService(lg, cfg, repo)
	controller := NewController(router, serviceImpl, cfg, lg)
	suite := &EndpointsTestSuite{controller: controller, repo: repo}

	suite.Test_SaveUrlHandler(t)
	suite.Test_GetOriginalUrlHandler(t)
	suite.Test_GetShortenURLHandler(t)
	suite.Test_PingHandler(t)
}
