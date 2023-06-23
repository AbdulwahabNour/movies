package data

import (
	"time"

	"github.com/AbdulwahabNour/movies/pkg/validation"
)


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

func ValidateMovie(movie *MovieBinding , validate *validation.Validatation){
     
    if !(movie.Year <= time.Now().Year()) {
        validate.Errors = append(validate.Errors, validation.ValidatationError{Field:"year", Message: "year shouldn't be after the current year"})
    }


    if !(len(movie.Genres) <= 5 && len(movie.Genres) >= 1) {
        validate.Errors =append(validate.Errors, validation.ValidatationError{Field:"genres", Message: "genres should be between 1 and 5 genre" }) 
    }
    
    ok :=  validate.IsUnique(movie.Genres)
    if !ok{
      validate.Errors = append(validate.Errors, validation.ValidatationError{Field:"genres", Message: "must not contain duplicate values" })
    }

}