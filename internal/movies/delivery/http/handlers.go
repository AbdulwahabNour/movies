package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AbdulwahabNour/movies/config"
	"github.com/AbdulwahabNour/movies/internal/movies"

	"github.com/AbdulwahabNour/movies/internal/model"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/AbdulwahabNour/movies/pkg/utils"
	"github.com/gin-gonic/gin"
)

 

type apiHandlers struct {
    config  *config.Config
    movieService movies.Service
    logger logger.Logger
}

func NewMovieHandlers(app *config.Config, serv movies.Service, logger logger.Logger) movies.Handler{
    return &apiHandlers{
        config: app,
        movieService: serv,
        logger: logger,
    }
}



// mux.PUT("/v1/movies/:id", h.editMovieHandler)
// mux.DELETE("/v1/movies/:id", h.deleteMovieHandler)


 
func(h *apiHandlers)HealthCheckHandler(c *gin.Context){
 
    c.IndentedJSON(http.StatusOK, gin.H{"status":"available", "enviroment": h.config.Server.Mode, "version": h.config.Server.AppVersion})

}

func(h *apiHandlers) CreateMovieHandler(c *gin.Context){
    
    var movie model.MovieBinding

    if errs := utils.ReadRequestJSON(c, &movie); errs != nil{
        utils.Response(c, http.StatusUnprocessableEntity, errs)
        return
    }
  
    validate :=  movie.ValidateMovie()
    h.logger.InfoLog("ValidateMovie", validate)
    if len(validate) != 0{
        utils.Response(c, http.StatusBadRequest, validate)
        return
    }
 

 

    c.JSON(
        http.StatusOK,
        gin.H{"status":"create a new movie"},
    )
   
}

 

func(h *apiHandlers) ShowMovieHandler(c *gin.Context){
    
    id, err :=  strconv.ParseInt(c.Param("id"), 10, 64)

    if err != nil{
        utils.NnotfoundResponse(c)
        return
    }
 
    movie := model.Movie{
        ID: id,
        Title: "Vaikings",
        Year: 2020,
        Runtime: 120,
        Genres: []string{"war", "drama"},
        Version: 1,
        CreateAt: time.Now(),
    }

  
    c.JSON(http.StatusOK, gin.H{"movie":movie})
    
}

func(h *apiHandlers) ListMoviesHandler(c *gin.Context){
     
    c.JSON(http.StatusOK, gin.H{"status":"create a new movie"})
 
    
}
func(h *apiHandlers) BadRequestResponse(c *gin.Context, err error){
    c.JSON(http.StatusOK, gin.H{"error": err.Error()})
    c.Abort()
}

// func(h *hlication) createMovieHandler(c *gin.Context){
     
//     c.JSON(http.StatusOK, gin.H{"status":"create a new movie"})
    
// }

