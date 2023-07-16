package model

import (
	"strings"
	"time"
	"unicode/utf8"
)

// Movie: database model for movies
type Movie struct {
	ID       int64     `json:"id"`                                        // Uniq integer Id for movie
	Title    string    `json:"title" validate:"required,min=5,max=200"`   // movie title
	Year     int       `json:"year" validate:"required,numeric,gte=1888"` //Movie release year
	Runtime  Runtime   `json:"runtime" validate:"required,numeric"`       //Movie runtime
	Genres   []string  `json:"genres" validate:"required"`                // Slice of genres for the movie
	Version  string    `json:"version"`
	CreateAt time.Time `json:"create_at"`
}

func (movie *Movie) ValidateMovie() map[string]string {
	err := make(map[string]string)

	if !(movie.Year <= time.Now().Year()) {
		err["year"] = "year shouldn't be after the current year"
	}

	if !(len(movie.Genres) <= 5 && len(movie.Genres) >= 1) {
		err["genres"] = "genres should be between 1 and 5 genre"
	}

	ok := movie.ValidateGenres()
	if !ok {
		err["genres"] = "genres should be unique and each genre is between 1 to 100 character"
	}
	return err
}
func (movie *Movie) PreCreate() {
	movie.Title = strings.TrimSpace(movie.Title)
}

func (movie *Movie) ValidateGenres() bool {
	seen := make(map[string]bool)
	for _, v := range movie.Genres {
		l := utf8.RuneCountInString(strings.TrimSpace(v))
		if seen[v] || l > 100 || l == 0 {
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

func (m *Movie) IsEmpty() bool {
	return m.Title == "" && m.Year == 0 && m.Runtime == 0 && len(m.Genres) == 0
}

type MovieSearchQuery struct {
	Title  string   `form:"title" json:"title" validate:"max=200"`
	Genres []string `form:"genres" json:"genres" validate:"max=5,unique"`
	Filter Filters
}

type Filters struct {
	Page         int    `form:"page"      json:"page" validate:"required,gte=1,lte=1000000"`
	PageSize     int    `form:"page_size" json:"page_size" validate:"required,gte=1,lte=100"`
	Sort         string `form:"sort"      json:"sort" validate:"oneof=id title year runtime -id -title -year -runtime"`
	TotalRecords int    `form:"-"  json:"-"`
}

func (f Filters) Limit() int {
	return f.PageSize
}
func (f Filters) Offset() int {
	return (f.Page - 1) * f.PageSize
}
func (f Filters) SortColumn() string {
	return strings.TrimPrefix(f.Sort, "-")
}
func (f Filters) SortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func (q *MovieSearchQuery) PrepareForQuery() {
	if q.Filter.Page == 0 {
		q.Filter.Page = 1
	}
	if q.Filter.PageSize == 0 {
		q.Filter.PageSize = 10
	}
	if q.Filter.Sort == "" {
		q.Filter.Sort = "id"
	}

	if q.Genres == nil {
		q.Genres = []string{}
	} else {

		if len(q.Genres) > 0 {
			q.Genres[0] = strings.TrimSpace(strings.Trim(q.Genres[0], ","))
			q.Genres = strings.Split(q.Genres[0], ",")
			q.FilterShortGenres()

		}
	}

}

func (q *MovieSearchQuery) FilterShortGenres() {

	f := make([]string, 0, len(q.Genres))

	for _, g := range q.Genres {
		l := utf8.RuneCountInString(g)
		if l < 100 && l > 0 {
			f = append(f, g)
		}
	}
	q.Genres = f
}
