package utils

import (
	"github.com/AbdulwahabNour/movies/pkg/httpError"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ReadRequestJSON(c *gin.Context, v any) error {

	return c.ShouldBindJSON(v)
}

// Refactored function to log an error using the provided logger.
func GinErrorLogWithFields(log logger.Logger, c *gin.Context, logMethod string, err error) {

	log.ErrorLogWithFields(logrus.Fields{"method": logMethod, "CLIENT_IP": c.ClientIP()}, err)
}

// Response is a function that sends a JSON response with the given status code and message.
func Response(c *gin.Context, status int, message any) {

	c.JSON(status, message)
}

func ErrorResponse(c *gin.Context, err error) {

	httperr, ok := err.(httpError.HttpErr)

	if !ok {
		httperr = httpError.ParseErrors(err)
	}

	c.AbortWithStatusJSON(httperr.Status(), gin.H{"error": httperr})

}
