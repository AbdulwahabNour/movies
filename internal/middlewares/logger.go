package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// LoggingMiddleware is a middleware function that logs HTTP requests.
//
// It takes a logger.Logger as a parameter and returns a gin.HandlerFunc.
// The gin.HandlerFunc is a function that takes a *gin.Context as a parameter
// and logs information about the HTTP request, including the request method,
// request route, status code, latency time, and client IP.
func (m *MiddleWares) LoggingMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		startTime := time.Now()

		ctx.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := ctx.Request.Method

		reqUri := ctx.Request.RequestURI

		statusCode := ctx.Writer.Status()

		clientIP := ctx.ClientIP()

		m.logger.InfoLogWithFields(log.Fields{
			"METHOD":    reqMethod,
			"URI":       reqUri,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIP,
		}, "HTTP REQUEST")

		ctx.Next()
	}
}
