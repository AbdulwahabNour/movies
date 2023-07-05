package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/AbdulwahabNour/movies/internal/model"
	"github.com/AbdulwahabNour/movies/internal/movies"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type movieRepo struct {
	db *sqlx.DB
}

func NewMovieRepo(db *sqlx.DB) movies.Repository {
	return &movieRepo{
		db: db,
	}
}

func (r *movieRepo) CreateMovie(ctx context.Context, movie *model.Movie) error {

	query := `INSERT INTO movies (title, year, runtime, genres) VALUES ($1, $2, $3, $4) RETURNING id, create_at, version`

	err := r.db.QueryRowContext(ctx,
		query,
		movie.Title,
		movie.Year,
		movie.Runtime,
		pq.Array(movie.Genres)).Scan(&movie.ID, &movie.CreateAt, &movie.Version)
	if err != nil {
		return fmt.Errorf("failed to insert movie: %w", err)
	}

	return nil

}
func (r *movieRepo) GetMovie(ctx context.Context, id int64) (*model.Movie, error) {

	query := `SELECT id, title, year, runtime, genres, create_at, version FROM movies WHERE id = $1`
	var movie model.Movie
	err := r.db.QueryRowContext(ctx, query, id).Scan(&movie.ID,
		&movie.Title,
		&movie.Year,
		&movie.Runtime,
		pq.Array(&movie.Genres),
		&movie.CreateAt,
		&movie.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, fmt.Errorf("movie with id %d not found", id)
		default:
			return nil, fmt.Errorf("failed to get movie: %w", err)
		}
	}

	return &movie, nil
}

func (r *movieRepo) UpdateMovie(ctx context.Context, movie *model.Movie) error {
	query := `UPDATE movies SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1 WHERE id = $5 returning id, version`

	return r.db.QueryRowContext(ctx,
		query,
		movie.Title,
		movie.Year,
		movie.Runtime,
		pq.Array(movie.Genres),
		movie.ID).Scan(&movie.ID, &movie.Version)
}
func (r *movieRepo) DeleteMovie(ctx context.Context, id int64) error {
	query := `DELETE FROM movies WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return fmt.Errorf("movie with id %d not found", id)
		default:
			return fmt.Errorf("failed to delete movie: %w", err)
		}
	}
	return nil
}
