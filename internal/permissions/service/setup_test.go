package service

import (
	"github.com/AbdulwahabNour/movies/config"
	"github.com/AbdulwahabNour/movies/internal/permissions"
	"github.com/AbdulwahabNour/movies/internal/permissions/mocks"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/go-playground/validator/v10"
)

func setupTest() (permissions.Service, *mocks.MockRepository) {
	mockRepo := new(mocks.MockRepository)
	config := config.Config{}
	logger := logger.NewApiLogger(&config)
	sevice := NewPermissionService(&config, mockRepo, logger, validator.New())
	return sevice, mockRepo
}
