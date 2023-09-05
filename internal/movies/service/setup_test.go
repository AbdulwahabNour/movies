package service

import (
	"github.com/AbdulwahabNour/movies/config"
	"github.com/AbdulwahabNour/movies/internal/movies"
	"github.com/AbdulwahabNour/movies/internal/movies/mocks"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/go-playground/validator/v10"
)

func setup_test() (movies.Service, *mocks.MockRepository) {
	mocRepo := new(mocks.MockRepository)
	config := new(config.Config)
	logger := logger.NewApiLogger(config)

	service := NewMovieService(config, mocRepo, logger, validator.New())
	return service, mocRepo
}
