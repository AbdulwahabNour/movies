package model

import (
	"strings"
	"time"
)

type MoviesRepo interface {
	InsertMovie(movie *Movie) error
	GetMovie(id int64) (*Movie, error)
	UpdateMovie(movie *Movie) error
	DeleteMovie(id int64) error
}

// Movie: database model for movies
type Movie struct {
	ID       int64     `json:"id"`                                        // Uniq integer Id for movie
	Title    string    `json:"title" validate:"required,min=5,max=200"`   // movie title
	Year     int       `json:"year" validate:"required,numeric,gte=1888"` //Movie release year
	Runtime  Runtime   `json:"runtime" validate:"required,numeric"`       //Movie runtime
	Genres   []string  `json:"genres" validate:"required"`                // Slice of genres for the movie
	Version  int       `json:"version"`
	CreateAt time.Time `json:"create_at"`
}

// MovieBinding: validation model for movie
// type MovieBinding struct {
// 	Title   string   `json:"title" binding:"required,min=5,max=200"`
// 	Year    int      `json:"year" binding:"required,numeric,gte=1888"`
// 	Runtime Runtime  `json:"runtime" binding:"required,numeric"`
// 	Genres  []string `json:"genres" binding:"required"`
// }

func (movie *Movie) ValidateMovie() map[string]string {
	err := make(map[string]string)

	if !(movie.Year <= time.Now().Year()) {
		err["year"] = "year shouldn't be after the current year"
	}

	if !(len(movie.Genres) <= 5 && len(movie.Genres) >= 1) {
		err["genres"] = "genres should be between 1 and 5 genre"
	}

	ok := movie.IsGenresUnique()
	if !ok {
		err["genres"] = "genres should be unique"
	}
	return err
}
func (movie *Movie) PreCreate() {
	movie.Title = strings.TrimSpace(movie.Title)
}

func (movie *Movie) IsGenresUnique() bool {
	seen := make(map[string]bool)
	for _, v := range movie.Genres {
		if seen[v] {
			return false
		}
		seen[v] = true
	}

	return true

}

// Copy copy the fields of the given Movie to the current non empty fields of the Movie.
func (movie *Movie) Copy(m *Movie) {
	// Set the ID field.
	movie.ID = m.ID

	// If the trimmed Title field is empty, set it to the given Title.
	if strings.TrimSpace(movie.Title) == "" {
		movie.Title = m.Title
	}

	// If the Year field is 0, set it to the given Year.
	if movie.Year == 0 {
		movie.Year = m.Year
	}

	// If the Runtime field is 0, set it to the given Runtime.
	if movie.Runtime == 0 {
		movie.Runtime = m.Runtime
	}

	// If the Genres field is nil, set it to the given Genres.
	if movie.Genres == nil {
		movie.Genres = m.Genres
	}

	// Set the Version field.
	movie.Version = m.Version

	// Set the CreateAt field.
	movie.CreateAt = m.CreateAt

}
