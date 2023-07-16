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
func ErrorLogWithFields(log logger.Logger, c *gin.Context, logmsg any, err error) {

	log.ErrorLogWithFields(logrus.Fields{"err": err, "CLIENT_IP": c.ClientIP()}, logmsg)
}

// Response is a function that sends a JSON response with the given status code and message.
func Response(c *gin.Context, status int, message any) {

	c.JSON(status, message)
}

func ErrorResponse(c *gin.Context, err error) {

	m := httpError.ParseErrors(err)

	// c.AbortWithError(http.StatusBadRequest, m)
	c.JSON(m.Status(), gin.H{"error": m})

}
