package middlewares

import (
	"github.com/AbdulwahabNour/movies/config"
	"github.com/AbdulwahabNour/movies/pkg/logger"
)

type MiddleWares struct {
	config *config.Config
	logger logger.Logger
}

func NewMiddleWares(config *config.Config, logger logger.Logger) *MiddleWares {
	return &MiddleWares{
		config: config,
		logger: logger,
	}
}
