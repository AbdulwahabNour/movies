package movies

import "github.com/AbdulwahabNour/movies/internal/model"


type Repository interface{
    InsertMovie(movie *model.Movie)  error
    GetMovie(id int64) (*model.Movie, error)
    UpdateMovie(movie *model.Movie) error
    DeleteMovie(id int64) error
}
