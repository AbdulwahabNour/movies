package movies

import (
	"context"

	"github.com/AbdulwahabNour/movies/internal/model"
)



type Service interface {
    CreateMovie(ctx context.Context, movie *model.Movie)  error
    GetMovie(ctx context.Context, id int64) (*model.Movie, error)
    UpdateMovie(ctx context.Context, movie *model.Movie) error
    DeleteMovie(ctx context.Context, id int64) error
}


