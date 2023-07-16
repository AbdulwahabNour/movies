package service

import (
	"context"

	"github.com/AbdulwahabNour/movies/config"
	model "github.com/AbdulwahabNour/movies/internal/model/movie"
	"github.com/AbdulwahabNour/movies/internal/movies"
	"github.com/AbdulwahabNour/movies/pkg/httpError"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/go-playground/validator/v10"
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
	err := s.checkMovie(movie)
	if err != nil {
		return err
	}

	return s.repo.CreateMovie(ctx, movie)
}
func (s *movieService) GetMovie(ctx context.Context, id int64) (*model.Movie, error) {
	if id < 1 {
		return nil, httpError.NewBadRequestError(httpError.ErrBadQuery)
	}

	return s.repo.GetMovie(ctx, id)
}
func (s *movieService) ListMoviesHandler(ctx context.Context, query *model.MovieSearchQuery) ([]*model.Movie, error) {
	query.PrepareForQuery()

	if err := s.validate.Struct(query); err != nil {
		return nil, httpError.ParseValidationErrors(err)
	}
	movies, err := s.repo.ListMoviesHandler(ctx, query)

	if err != nil {
		return nil, httpError.NewInternalServerError(err)

	}
	return movies, nil
}
func (s *movieService) UpdateMovie(ctx context.Context, movie *model.Movie) error {

	if movie.ID < 1 {
		return httpError.NewBadRequestError(httpError.ErrBadQuery)
	}
	if movie.IsEmpty() {
		return httpError.NewBadRequestError(httpError.ErrInvalidSyntax)
	}

	getMovie, err := s.GetMovie(ctx, movie.ID)

	if err != nil {
		return httpError.NewNotFoundError(err)
	}

	movie.Copy(getMovie)

	err = s.checkMovie(movie)
	if err != nil {
		return err
	}

	return s.repo.UpdateMovie(ctx, movie)
}
func (s *movieService) DeleteMovie(ctx context.Context, id int64) error {
	if id < 1 {
		return httpError.NewBadRequestError(httpError.ErrBadQuery)
	}

	return s.repo.DeleteMovie(ctx, id)
}
func (s *movieService) checkMovie(movie *model.Movie) error {

	validate := movie.ValidateMovie()

	if len(validate) != 0 {
		return httpError.NewUnprocessableEntityError(validate)
	}

	if err := s.validate.Struct(movie); err != nil {
		return httpError.ParseValidationErrors(err)
	}

	return nil
}
