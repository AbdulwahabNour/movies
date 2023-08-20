package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	model "github.com/AbdulwahabNour/movies/internal/model/movie"
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
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&movie.ID,
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

func (r *movieRepo) ListMovies(ctx context.Context, filter *model.MovieSearchQuery) ([]*model.Movie, error) {

	query := fmt.Sprintf(`SELECT count(*) over() ,id, create_at, title, year, runtime, genres, version FROM movies 
	WHERE (to_tsvector('simple',title) @@ plainto_tsquery('simple', $1) OR $1 = '') AND (genres @> $2 OR $2 = '{}')
	order by %s %s, id desc LIMIT $3 OFFSET $4`, filter.Filter.SortColumn(), filter.Filter.SortDirection())

	rows, err := r.db.QueryContext(ctx, query, filter.Title, pq.Array(filter.Genres), filter.Filter.Limit(), filter.Filter.Offset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totalRecord := 0

	movies := make([]*model.Movie, 0)
	for rows.Next() {

		var movie model.Movie

		err := rows.Scan(
			&totalRecord,
			&movie.ID,
			&movie.CreateAt,
			&movie.Title,
			&movie.Year,
			&movie.Runtime,
			pq.Array(&movie.Genres),
			&movie.Version)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	filter.Filter.TotalRecords = totalRecord
	return movies, nil
}

func (r *movieRepo) UpdateMovie(ctx context.Context, movie *model.Movie) error {
	query := `UPDATE movies SET title = $1, year = $2, runtime = $3, genres = $4, version =  uuid_generate_v4() WHERE id = $5 And version = $6 returning  version`

	err := r.db.QueryRowContext(ctx,
		query,
		movie.Title,
		movie.Year,
		movie.Runtime,
		pq.Array(movie.Genres),
		movie.ID,
		movie.Version).Scan(&movie.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return errors.New("edit conflict")
		default:
			return err
		}
	}

	return nil
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
