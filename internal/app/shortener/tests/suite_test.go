package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/sprint1/config"
	"github.com/sprint1/internal/app/shortener/endpoints/public"
	"github.com/sprint1/internal/app/shortener/mocks"
	"github.com/sprint1/internal/app/shortener/service"
	"github.com/sprint1/internal/app/shortener/workers"
)

type EndpointsTestSuite struct {
	controller *public.Controller
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
	workerPool := workers.NewWorkerPool(lg, repo)
	workerPool.Start()
	serviceImpl := service.NewService(lg, cfg, repo, workerPool)
	controller := public.NewController(router, serviceImpl, cfg, lg)
	suite := &EndpointsTestSuite{controller: controller, repo: repo}

	suite.Test_GetOriginalUrlHandler_Success(t)
	suite.Test_GetOriginalUrlHandler_ErrorResourceGone(t)
	suite.Test_GetOriginalUrlHandler_ErrorBadRequest(t)
	suite.Test_GetOriginalUrlHandler_Error(t)
	suite.Test_GetShortenURLHandler_Success(t)
	suite.Test_GetShortenURLHandler_ConflictError(t)
	suite.Test_GetShortenURLHandler_BadRequestError(t)
	suite.Test_GetShortenURLHandler_CreateUrlError(t)
	suite.Test_GetShortenURLHandler_UnmarshallError(t)
	suite.Test_PingHandler(t)
	suite.Test_CreateUserHandler_Success(t)
	suite.Test_CreateUserHandler_CreateUserError(t)
	suite.Test_CreateUserHandler_BadRequestError(t)
	suite.Test_SaveUrlHandler_Success(t)
	suite.Test_SaveUrlHandler_ConflictError(t)
	suite.Test_SaveUrlHandler_CreateUrlError(t)
	suite.Test_SaveUrlHandler_BadRequestError(t)
	suite.Test_AuthUserHandler_Success(t)
	suite.Test_AuthUserHandler_UserNoFound(t)
	suite.Test_GetUserUrlsHandler_Success(t)
	suite.Test_GetUserUrlsHandler_NoContent(t)
	suite.Test_SaveUrlsBatchHandler(t)
}
