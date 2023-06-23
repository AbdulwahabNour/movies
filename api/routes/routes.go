package routes

import (
	"github.com/AbdulwahabNour/movies/api/handlers"
	"github.com/gin-gonic/gin"
)

type ApiRouter interface {
    SetRoutes(r *gin.Engine)
}


type apiRouteres struct {
    Handler  handlers.ApiHandler
}

func NewApiHandlers(handler handlers.ApiHandler) ApiRouter {
    return &apiRouteres{
        Handler: handler,
    }
}
 

func(app *apiRouteres)SetRoutes(r *gin.Engine) {
 
    r.GET("/v1/healthcheck", app.Handler.HealthCheckHandler) 
    r.GET("/v1/movies", app.Handler.ListMoviesHandler)
    r.POST("/v1/movies", app.Handler.CreateMovieHandler)
    r.GET("/v1/movies/:id", app.Handler.ShowMovieHandler) 
 
    // mux.PUT("/v1/movies/:id", app.editMovieHandler) 
    // mux.DELETE("/v1/movies/:id", app.deleteMovieHandler) 

}

