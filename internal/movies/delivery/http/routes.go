package http

import (
	"github.com/AbdulwahabNour/movies/internal/middlewares"
	"github.com/AbdulwahabNour/movies/internal/movies"
	"github.com/gin-gonic/gin"
)

func MapMoviesRoutes(r *gin.RouterGroup, app movies.Handler, mw *middlewares.MiddleWares) {

	r.GET("/healthcheck", app.HealthCheckHandler)

	r.GET("/movies", app.ListMoviesHandler)
	r.GET("/movies/:id", app.ShowMovieHandler)

	r.POST("/movies", app.CreateMovieHandler)
	r.PUT("/movies/:id", app.UpdateMovieHandler)
	r.DELETE("/movies/:id", app.DeleteMovieHandler)

}
