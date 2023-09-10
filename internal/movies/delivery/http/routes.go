package http

import (
	"github.com/AbdulwahabNour/movies/internal/middlewares"
	"github.com/AbdulwahabNour/movies/internal/movies"
	"github.com/gin-gonic/gin"
)

func MapMoviesRoutes(r *gin.RouterGroup, app movies.Handler, mw *middlewares.MiddleWares) {

	r.GET("/movies", app.ListMoviesHandler)
	r.GET("/movies/:id", app.ShowMovieHandler)

	r.POST("/movies", mw.RequirePermission("movie:add"), app.CreateMovieHandler)
	r.PUT("/movies/:id", mw.RequirePermission("movie:update"), app.UpdateMovieHandler)
	r.DELETE("/movies/:id", mw.RequirePermission("movie:delete"), app.DeleteMovieHandler)

}
