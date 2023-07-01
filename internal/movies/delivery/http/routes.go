package http

import (
	"github.com/AbdulwahabNour/movies/internal/movies"
	"github.com/gin-gonic/gin"
)


func MapMoviesRoutes(r *gin.RouterGroup, app movies.Handler) {
 
    r.GET("/healthcheck", app.HealthCheckHandler) 
    r.GET("/movies", app.ListMoviesHandler)
    r.POST("/movies", app.CreateMovieHandler)
    r.GET("/movies/:id", app.ShowMovieHandler) 
 
    // mux.PUT("/v1/movies/:id", app.editMovieHandler) 
    // mux.DELETE("/v1/movies/:id", app.deleteMovieHandler) 

}

