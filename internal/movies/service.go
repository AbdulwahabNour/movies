package movies

import (
	"context"

	model "github.com/AbdulwahabNour/movies/internal/model/movie"
)

type Service interface {
	CreateMovie(ctx context.Context, movie *model.Movie) error
	GetMovie(ctx context.Context, id int64) (*model.Movie, error)
	ListMovies(ctx context.Context, query *model.MovieSearchQuery) ([]*model.Movie, error)
	UpdateMovie(ctx context.Context, movie *model.Movie) error
	DeleteMovie(ctx context.Context, id int64) error
}
