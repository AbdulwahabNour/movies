package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AbdulwahabNour/movies/internal/config"
	"github.com/AbdulwahabNour/movies/internal/data"
	"github.com/AbdulwahabNour/movies/pkg/validation"
	"github.com/gin-gonic/gin"
)

type ApiHandler interface{
 
    HealthCheckHandler(c *gin.Context)
    CreateMovieHandler(c *gin.Context)
    ShowMovieHandler(c *gin.Context)
    ListMoviesHandler(c *gin.Context)
    
}


type apiHandlers struct {
    App  *config.App 
}

func NewApiHandlers(app *config.App) ApiHandler{
    return &apiHandlers{
        App: app,
    }
}



// mux.PUT("/v1/movies/:id", h.editMovieHandler)
// mux.DELETE("/v1/movies/:id", h.deleteMovieHandler)


 
func(h *apiHandlers)HealthCheckHandler(c *gin.Context){
 
    c.IndentedJSON(http.StatusOK, gin.H{"status":"available", "enviroment": h.App.Config.Env, "version": h.App.Config.Version})

}

func(h *apiHandlers) CreateMovieHandler(c *gin.Context){
    var movie data.MovieBinding
    movieValidation :=  validation.NewValidatation()
    
 
    if  err := c.ShouldBindJSON(&movie); err != nil{

        h.logError(c, err)
        
        movieValidation.DescriptiveValidationMessages(err)
        h.errorResponse(
                        c,
                        http.StatusUnprocessableEntity,
                        gin.H{"error": movieValidation.Errors})
        
        return
    }

   data.ValidateMovie(&movie, movieValidation)

   if !movieValidation.NoErrors(){
       h.errorResponse(
                         c, 
                         http.StatusBadRequest,
                         gin.H{"error": movieValidation.Errors})

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
     
        h.notfoundResponse(c)
        return
    }
 
    movie := data.Movie{
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

