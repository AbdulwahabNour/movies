package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *MiddleWares) EnableCORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Header("Vary", "Origin")

		ctx.Writer.Header().Add("Vary", "Access-Control-Request-Method")

		origin := ctx.Request.Header.Get("Origin")

		if origin != "" {

			for i := range m.config.CORS.TrustedOrigins {

				if origin == m.config.CORS.TrustedOrigins[i] {

					ctx.Header("Access-Control-Allow-Origin", origin)

					if ctx.Request.Method == http.MethodOptions && ctx.Request.Header.Get("Access-Control-Request-Method") != "" {

						ctx.Header("Access-Control-Request-Method", "OPTIONS, PUT, PATCH, DELETE")
						ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
						ctx.Status(http.StatusOK)

						return

					}

					break
				}

			}
		}
		ctx.Next()
	}
}
