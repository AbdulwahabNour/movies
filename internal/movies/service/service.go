package service

import (
	"context"

	"github.com/AbdulwahabNour/movies/config"
	model "github.com/AbdulwahabNour/movies/internal/model/movie"
	"github.com/AbdulwahabNour/movies/internal/movies"
	"github.com/AbdulwahabNour/movies/pkg/httpError"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type movieService struct {
	config   *config.Config
	repo     movies.Repository
	logger   logger.Logger
	validate *validator.Validate
}

func NewMovieService(config *config.Config, repo movies.Repository, logger logger.Logger, validate *validator.Validate) movies.Service {
	return &movieService{
		config:   config,
		repo:     repo,
		logger:   logger,
		validate: validate,
	}
}

func (s *movieService) CreateMovie(ctx context.Context, movie *model.Movie) error {

	if err := s.checkMovie(movie); err != nil {
		return httpError.NewUnprocessableEntityError(err)
	}
	err := s.repo.CreateMovie(ctx, movie)
	if err != nil {
		s.logger.ErrorLogWithFields(logrus.Fields{"method": "movies.service.CreateMovie", "query": *movie}, err)
		return httpError.NewInternalServerError("error happen during create movie try again later")
	}

	return nil
}
func (s *movieService) GetMovie(ctx context.Context, id int64) (*model.Movie, error) {
	if id < 1 {
		return nil, httpError.NewNotFoundError("movie not found")
	}
	movie, err := s.repo.GetMovie(ctx, id)
	if err != nil {
		return nil, httpError.NewInternalServerError(err)
	}

	return movie, nil
}
func (s *movieService) ListMovies(ctx context.Context, query *model.MovieSearchQuery) ([]*model.Movie, error) {

	if err := s.validate.Struct(query); err != nil {
		return nil, httpError.ParseValidationErrors(err)
	}
	query.PrepareForQuery()

	movies, err := s.repo.ListMovies(ctx, query)

	if err != nil {
		return nil, httpError.NewInternalServerError(err)
	}
	return movies, nil
}
func (s *movieService) UpdateMovie(ctx context.Context, movie *model.Movie) error {
	if movie.IsEmpty() {
		return httpError.NewBadRequestError("The JSON payload is empty. Please provide valid data to update the movie.")
	}
	getMovie, err := s.GetMovie(ctx, movie.ID)
	if err != nil {
		return err
	}

	getMovie.PrepareForUpdate(movie)

	if err := s.checkMovie(getMovie); err != nil {
		return err
	}

	if err := s.repo.UpdateMovie(ctx, getMovie); err != nil {
		return httpError.NewInternalServerError(err)
	}

	return nil
}
func (s *movieService) DeleteMovie(ctx context.Context, id int64) error {
	if id < 1 {
		return httpError.NewBadRequestError("movie id less than 1")
	}
	if err := s.repo.DeleteMovie(ctx, id); err != nil {
		return httpError.NewInternalServerError(err)
	}
	return nil
}
func (s *movieService) checkMovie(movie *model.Movie) error {
	validate := movie.ValidateMovie()
	if movie.ID < 0 {
		return httpError.NewBadRequestError("movie id less than 0")
	} else if len(validate) != 0 {
		return httpError.NewUnprocessableEntityError(validate)
	} else if err := s.validate.Struct(movie); err != nil {
		return httpError.ParseValidationErrors(err)
	}
	return nil
}
