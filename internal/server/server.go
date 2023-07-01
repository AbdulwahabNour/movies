package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AbdulwahabNour/movies/config"
	moviesHttp "github.com/AbdulwahabNour/movies/internal/movies/delivery/http"
	psqlRepo "github.com/AbdulwahabNour/movies/internal/movies/repository/postgres"
	"github.com/AbdulwahabNour/movies/internal/movies/service"

	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)




const(
    maxHeaderBytes = 1 << 20//1MB
    ctxTimeout = 5
)
type Server struct{
     ginEngin *gin.Engine
     config *config.Config
     Logger logger.Logger
     db *sqlx.DB
}

func NewServer( config *config.Config, logger logger.Logger, db *sqlx.DB) *Server{
    return &Server{
        ginEngin: gin.Default(),
        config: config,
        Logger: logger,
        db: db,
    }
}

func(s *Server)MapHandler(g *gin.Engine) error{
    
    movieRepo := psqlRepo.NewMovieRepo(s.db)
    movieService :=  service.NewMovieService(movieRepo)
    movieHandler := moviesHttp.NewMovieHandlers(s.config, movieService, s.Logger)

    v1 :=  g.Group("/api/v1")
    moviesHttp.MapMoviesRoutes(v1, movieHandler)

  

    return nil
    
}

func (s *Server) Run()error{


    err :=  s.MapHandler(s.ginEngin)
    
    if err != nil{
         return err
    }
 
    srv := &http.Server{
        Addr: s.config.Server.Port,
        Handler: s.ginEngin,
        ReadTimeout: s.config.Server.ReadTimeout,
        WriteTimeout: s.config.Server.WriteTimeout,
        MaxHeaderBytes: maxHeaderBytes, 
    }

   go func() {

            if err := srv.ListenAndServe(); err != nil {
                s.Logger.ErrorLog(err)
            }
    
   }()
   s.Logger.InfoLog(fmt.Sprintf("Server running on port %s", s.config.Server.Port))

   c := make(chan os.Signal, 1)
   signal.Notify(c, os.Interrupt)
   <-c
   ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
   defer cancle()
   srv.Shutdown(ctx)
   s.Logger.InfoLog("Server exiting")
   return nil
    
}