package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AbdulwahabNour/movies/config"
	"github.com/AbdulwahabNour/movies/internal/middlewares"
	moviesHttp "github.com/AbdulwahabNour/movies/internal/movies/delivery/http"
	moviesRepo "github.com/AbdulwahabNour/movies/internal/movies/repository/postgres"
	moviesService "github.com/AbdulwahabNour/movies/internal/movies/service"
	tokenHttp "github.com/AbdulwahabNour/movies/internal/token/delivery/http"
	tokenRedisRepo "github.com/AbdulwahabNour/movies/internal/token/repository/redis"
	tokenService "github.com/AbdulwahabNour/movies/internal/token/service"
	usersHttp "github.com/AbdulwahabNour/movies/internal/users/delivery/http"
	usersRepo "github.com/AbdulwahabNour/movies/internal/users/repository/postgres"
	usersService "github.com/AbdulwahabNour/movies/internal/users/service"
	"github.com/go-playground/validator/v10"

	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

const (
	maxHeaderBytes = 1 << 20 //1MB
)

type Server struct {
	ginEngin *gin.Engine
	validate *validator.Validate
	config   *config.Config
	Logger   logger.Logger
	db       *sqlx.DB
	RedisDB  *redis.Client
}

func NewServer(config *config.Config, logger logger.Logger, db *sqlx.DB, redisDb *redis.Client) *Server {
	return &Server{
		ginEngin: gin.New(),
		validate: validator.New(),
		config:   config,
		Logger:   logger,
		db:       db,
		RedisDB:  redisDb,
	}
}

func (s *Server) MapHandler(g *gin.Engine) error {

	tokeRepo := tokenRedisRepo.NewTokenRepo(s.RedisDB)
	tokenServ := tokenService.NewTokenService(s.config, s.Logger, tokeRepo)

	movieRepo := moviesRepo.NewMovieRepo(s.db)
	movieService := moviesService.NewMovieService(s.config, movieRepo, s.Logger, s.validate)

	userRepo := usersRepo.NewUserRepo(s.db)
	userService := usersService.NewUserService(s.config, userRepo, tokenServ, s.Logger, s.validate)

	movieHandler := moviesHttp.NewMovieHandlers(s.config, movieService, s.Logger)
	userHandler := usersHttp.NewMovieHandlers(s.config, userService, s.Logger)
	tokenHandler := tokenHttp.NewTokenHandlers(s.config, tokenServ, userService, s.Logger)

	v1 := g.Group("/api/v1")
	usersHttp.MapUsersRoutes(v1, userHandler)
	moviesHttp.MapMoviesRoutes(v1, movieHandler)
	tokenHttp.MapTokenRoutes(v1, tokenHandler)

	return nil

}

func (s *Server) Run() error {
	middleware := middlewares.NewMiddleWares(s.config, s.Logger)
	s.ginEngin.Use(middleware.LoggingMiddleware())

	if s.config.Limiter.Enabled {
		s.ginEngin.Use(middleware.RateLimitMiddleware())
	}

	s.ginEngin.Use(gin.Recovery())

	err := s.MapHandler(s.ginEngin)

	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:           s.config.Server.Port,
		Handler:        s.ginEngin,
		ReadTimeout:    s.config.Server.ReadTimeout,
		WriteTimeout:   s.config.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {

		if err := srv.ListenAndServe(); err != nil {
			s.Logger.ErrorLog(err)
		}

	}()

	s.Logger.InfoLog(fmt.Sprintf("Server running on port %s", s.config.Server.Port))

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	mesg := <-c
	s.Logger.InfoLog(fmt.Sprintf("Server exiting with signal %s", mesg))

	ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancle()
	err = srv.Shutdown(ctx)
	if err != nil {
		return err
	}

	s.Logger.InfoLog("Server exiting")
	return nil

}
