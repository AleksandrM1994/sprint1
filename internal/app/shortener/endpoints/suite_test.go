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

	repo := mocks.NewMockRepo(ctrl)
	serviceImpl := service.NewService(lg, cfg, repo)
	serviceImpl.URLStorage = map[string]string{"aHR0cHM6Ly9qc29uZm9ybWF0dGVyLm9yZw==": "https://jsonformatter.org"}
	controller := NewController(router, serviceImpl, cfg, lg)
	suite := &EndpointsTestSuite{controller: controller}

	suite.Test_SaveUrlHandler(t)
	suite.Test_GetOriginalUrlHandler(t)
	suite.Test_GetShortenURLHandler(t)
}
