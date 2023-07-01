package service

import (
	"github.com/AbdulwahabNour/movies/config"
	"github.com/AbdulwahabNour/movies/internal/model"
	"github.com/AbdulwahabNour/movies/internal/movies"
	"github.com/AbdulwahabNour/movies/pkg/logger"
)

type movieService struct{
    config *config.Config
    repo movies.Repository
    logger logger.Logger
}

func NewMovieService(repo movies.Repository) movies.Service{
    return &movieService{
        repo: repo,
    }
}

func(s *movieService)InsertMovie(movie *model.Movie)  error{
    return nil
}
func(s *movieService)GetMovie(id int64) (*model.Movie, error){
    return nil, nil
}
func(s *movieService)UpdateMovie(movie *model.Movie) error{
    return nil
}
func(s *movieService)DeleteMovie(id int64) error{
    return nil
}