package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/AbdulwahabNour/movies/config"
	"github.com/AbdulwahabNour/movies/internal/movies"

	model "github.com/AbdulwahabNour/movies/internal/model/movie"
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

func (h *apiHandlers) CreateMovieHandler(c *gin.Context) {

	var movie model.Movie
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	if err := utils.ReadRequestJSON(c, &movie); err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "movies.handlers.CreateMovieHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	err := h.movieService.CreateMovie(ctx, &movie)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "movies.handlers.CreateMovieHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusCreated, movie)

}

func (h *apiHandlers) ShowMovieHandler(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "movies.handlers.ShowMovie", err)
		utils.ErrorResponse(c, err)
		return
	}

	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	movie, err := h.movieService.GetMovie(ctx, id)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "movies.handlers.ShowMovie.GetMovie", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusOK, movie)

}

func (h *apiHandlers) ListMoviesHandler(c *gin.Context) {

	var filter model.MovieSearchQuery

	err := c.ShouldBindQuery(&filter)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "movies.handlers.ListMoviesHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	movies, err := h.movieService.ListMovies(c, &filter)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "movies.handlers.ListMoviesHandler.movieService", err)
		utils.ErrorResponse(c, err)
		return
	}

	meataData := utils.CalculateMetaData(filter.Filter.TotalRecords, filter.Filter.Page, filter.Filter.PageSize)

	c.JSON(http.StatusOK, gin.H{"metadata": meataData, "movies": movies})

}
func (h *apiHandlers) UpdateMovieHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "movies.handlers.UpdateMovieHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	var movie model.Movie
	if err := utils.ReadRequestJSON(c, &movie); err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "movies.handlers.UpdateMovieHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	movie.ID = id

	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	err = h.movieService.UpdateMovie(ctx, &movie)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "movies.handlers.UpdateMovieHandler.UpdateMovie", err)
		utils.ErrorResponse(c, err)
		return
	}
	utils.Response(c, http.StatusOK, movie)
}
func (h *apiHandlers) DeleteMovieHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "movies.handlers.DeleteMovieHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	err = h.movieService.DeleteMovie(ctx, id)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "movies.handlers.DeleteMovieHandler.DeleteMovie", err)
		utils.ErrorResponse(c, err)
		return
	}
	utils.Response(c, http.StatusOK, gin.H{"status": "deleted"})
}
