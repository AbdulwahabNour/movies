package http

import (
	"github.com/AbdulwahabNour/movies/config"
	"github.com/AbdulwahabNour/movies/internal/movies"
	"github.com/AbdulwahabNour/movies/internal/movies/mocks"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/gin-gonic/gin"
)

func setupTest() (*gin.Engine, movies.Handler, *mocks.MockService) {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	// NewMovieHandlers(app *config.Config, serv movies.Service, logger logger.Logger)
	config := config.Config{}
	mockServ := new(mocks.MockService)
	logger := logger.NewApiLogger(&config)
	handlers := NewMovieHandlers(&config, mockServ, logger)

	return router, handlers, mockServ
}
