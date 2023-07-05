package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/AbdulwahabNour/movies/config"
	"github.com/AbdulwahabNour/movies/internal/movies"

	"github.com/AbdulwahabNour/movies/internal/model"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/AbdulwahabNour/movies/pkg/utils"
	"github.com/gin-gonic/gin"
)

type apiHandlers struct {
	config       *config.Config
	movieService movies.Service
	logger       logger.Logger
}

func NewMovieHandlers(app *config.Config, serv movies.Service, logger logger.Logger) movies.Handler {
	return &apiHandlers{
		config:       app,
		movieService: serv,
		logger:       logger,
	}
}

func (h *apiHandlers) HealthCheckHandler(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"status": "available", "enviroment": h.config.Server.Mode, "version": h.config.Server.AppVersion})

}

func (h *apiHandlers) CreateMovieHandler(c *gin.Context) {

	var movie model.Movie
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	if err := utils.ReadRequestJSON(c, &movie); err != nil {

		h.logger.ErrorLog("CreateMovie utils.ReadRequestJSON", err)
		utils.ErrorResponse(c, err)
		return
	}

	err := h.movieService.CreateMovie(ctx, &movie)

	if err != nil {
		h.logger.ErrorLog("CreateMovie h.movieService.CreateMovie", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusCreated, movie)

}

func (h *apiHandlers) ShowMovieHandler(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		h.logger.ErrorLog("ShowMovie strconv.ParseInt: ", err)
		utils.ErrorResponse(c, err)
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	movie, err := h.movieService.GetMovie(ctx, id)
	if err != nil {
		h.logger.ErrorLog("ShowMovie h.movieService.GetMovie:", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusOK, movie)

}

func (h *apiHandlers) ListMoviesHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "create a new movie"})

}
func (h *apiHandlers) UpdateMovieHandler(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	var movie model.Movie

	if err != nil {
		h.logger.ErrorLog("UpdateMovie strconv.ParseInt :", err)
		utils.ErrorResponse(c, err)
		return
	}

	if err := utils.ReadRequestJSON(c, &movie); err != nil {
		h.logger.ErrorLog("UpdateMovie utils.ReadRequestJSON :", err)
		utils.ErrorResponse(c, err)
		return
	}

	movie.ID = id
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	err = h.movieService.UpdateMovie(ctx, &movie)

	if err != nil {
		h.logger.ErrorLog("UpdateMovie h.movieService.UpdateMovie:", err)
		utils.ErrorResponse(c, err)
		return
	}
	utils.Response(c, http.StatusOK, movie)
}
func (h *apiHandlers) DeleteMovieHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.ErrorLog("DeleteMovie strconv.ParseInt :", err)
		utils.ErrorResponse(c, err)
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	err = h.movieService.DeleteMovie(ctx, id)
	if err != nil {
		h.logger.ErrorLog("DeleteMovie h.movieService.DeleteMovie :", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusOK, gin.H{"status": "deleted"})
}
