package main

import (
	"github.com/gin-gonic/gin"
)



func(app *application)setRoutes(mux *gin.Engine) {

    mux.GET("/v1/healthcheck", app.healthCheckHandler) 
    mux.GET("/v1/movies", app.listMoviesHandler)
    mux.POST("/v1/movies", app.createMovieHandler)
    mux.GET("/v1/movies/:id", app.showMovieHandler) 
    mux.PUT("/v1/movies/:id", app.editMovieHandler) 
    mux.DELETE("/v1/movies/:id", app.deleteMovieHandler) 

}