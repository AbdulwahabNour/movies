package middlewares

import (
	"github.com/AbdulwahabNour/movies/config"
	"github.com/AbdulwahabNour/movies/pkg/logger"
)

type MiddleWares struct {
	cfg    *config.Config
	logger logger.Logger
}

func NewMiddleWares(cfg *config.Config, logger logger.Logger) *MiddleWares {
	return &MiddleWares{
		cfg:    cfg,
		logger: logger,
	}
}
