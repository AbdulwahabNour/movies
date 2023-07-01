package postgres

import (
	"github.com/AbdulwahabNour/movies/internal/model"
	"github.com/AbdulwahabNour/movies/internal/movies"
	"github.com/jmoiron/sqlx"
)


type movieRepo struct{
    db *sqlx.DB
}

func NewMovieRepo(db *sqlx.DB) movies.Repository{
    return &movieRepo{
        db: db,
    }
}

func(db *movieRepo)InsertMovie(movie *model.Movie)  error{
 
    return nil
}
func(db *movieRepo)GetMovie(id int64) (*model.Movie, error){

    return nil, nil
}

func(db *movieRepo)UpdateMovie(movie *model.Movie) error{
        return nil
}
func(db *movieRepo)DeleteMovie(id int64) error{
      return nil
}