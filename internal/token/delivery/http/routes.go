package http

import (
	"github.com/AbdulwahabNour/movies/internal/middlewares"
	"github.com/AbdulwahabNour/movies/internal/token"
	"github.com/gin-gonic/gin"
)

func MapTokenRoutes(r *gin.RouterGroup, app token.Handler, mw *middlewares.MiddleWares) {
	r.GET("/activate", app.Activate)
}
