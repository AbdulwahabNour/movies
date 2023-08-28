package http

import (
	"github.com/AbdulwahabNour/movies/config"
	"github.com/AbdulwahabNour/movies/internal/permissions"
	"github.com/AbdulwahabNour/movies/internal/permissions/mocks"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/gin-gonic/gin"
)

func setupTest() (*gin.Engine, permissions.Handler, *mocks.MockService) {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.MockService)
	config := &config.Config{}

	logger := logger.NewApiLogger(config) // Provide your logger instance
	handlers := NewPermissionsHandlers(config, mockService, logger)

	return router, handlers, mockService
}
