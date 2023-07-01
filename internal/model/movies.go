package model

import (
	"strings"
	"time"
)

type MoviesRepo interface{
    InsertMovie(movie *Movie)  error
    GetMovie(id int64) (*Movie, error)
    UpdateMovie(movie *Movie) error
    DeleteMovie(id int64) error
}

 

// Movie: database model for movies
type Movie struct{
    ID int64 `json:"id"` // Uniq integer Id for movie
    Title string `json:"title"` // movie title
    Year int `json:"year"` //Movie release year
    Runtime Runtime `json:"run_time"` //Movie runtime
    Genres []string `json:"genres"` // Slice of genres for the movie
    Version int `json:"version"`
    CreateAt time.Time `json:"create_at"`
    UpdateAt time.Time `json:"update_at"`
}
// MovieBinding: validation model for movie
type MovieBinding struct{
    Title string `json:"title" binding:"required,min=5,max=200"`
    Year int `json:"year" binding:"required,numeric,gte=1888"`
    Runtime Runtime `json:"run_time" binding:"required,numeric"`
    Genres []string `json:"genres" binding:"required"`
}

// ValidateMovie validates a Movie object and appends errors to a validation object.
//
// movie: a pointer to a Movie object.
// validate: a pointer to a validation object.

func (movie *MovieBinding)ValidateMovie( )map[string]string{
    err := make(map[string]string)

    if !(movie.Year <= time.Now().Year()) {
        err["year"] = "year shouldn't be after the current year"
    }


    if !(len(movie.Genres) <= 5 && len(movie.Genres) >= 1) {
        err["genres"] = "genres should be between 1 and 5 genre"
    }
    
    ok :=  movie.IsGenresUnique()
    if !ok{
        err["genres"] = "genres should be unique"
    }
 return err
}
func (movie *MovieBinding)PreCreate(){
   movie.Title = strings.TrimSpace(movie.Title)
}

func (movie *MovieBinding)IsGenresUnique() bool{
        seen := make(map[string]bool)
        for _, v := range movie.Genres{
            if seen[v]{
                return false
            }
            seen[v]= true
        }
        
        return true
   
}