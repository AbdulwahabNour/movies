package utils

import "github.com/AbdulwahabNour/movies/pkg/logger"

func BackgroundWithRecover(log logger.Logger, fn func()) {

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.ErrorLog(err)
			}
		}()
		fn()
	}()
}
