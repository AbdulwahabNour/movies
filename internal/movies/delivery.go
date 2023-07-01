package movies

import "github.com/gin-gonic/gin"

type Handler interface{
 
    HealthCheckHandler(c *gin.Context)
    CreateMovieHandler(c *gin.Context)
    ShowMovieHandler(c *gin.Context)
    ListMoviesHandler(c *gin.Context)
    
}
