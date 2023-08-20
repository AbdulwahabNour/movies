package http

import (
	"github.com/AbdulwahabNour/movies/internal/token"
	"github.com/gin-gonic/gin"
)

func MapTokenRoutes(r *gin.RouterGroup, app token.Handler) {
	r.GET("/activate", app.Activate)
}
