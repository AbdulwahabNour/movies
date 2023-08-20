package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	model "github.com/AbdulwahabNour/movies/internal/model/users"
	"github.com/AbdulwahabNour/movies/internal/users"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) users.Repository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) InsertUser(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (name, email, password_hash, activated) VALUES ($1, $2, $3, $4) RETURNING id, create_at, version`
	if user.Activated == nil {
		user.Activated = new(bool)
		*user.Activated = false
	}

	err := u.db.QueryRowContext(ctx,
		query,
		user.Name,
		user.Email,
		user.HashedPassword,
		*user.Activated).Scan(&user.ID, &user.CreateAt, &user.Version)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, create_at, name, email, password_hash, activated, version FROM users WHERE email= $1`
	var user model.User
	err := u.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.CreateAt,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.Activated,
		&user.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, fmt.Errorf("user with email %s not found", email)
		default:
			return nil, err
		}
	}
	return &user, nil
}
func (u *userRepo) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	query := `SELECT id, create_at, name, email, password_hash, activated, version FROM users WHERE id= $1`
	var user model.User
	err := u.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.CreateAt,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.Activated,
		&user.Version)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, fmt.Errorf("user with id %d not found", id)
		default:
			return nil, err
		}

	}
	return &user, nil
}
func (u *userRepo) UpdateUser(ctx context.Context, user *model.User) error {
	query := `UPDATE users SET name=$1, email=$2, password_hash=$3, activated=$4, version=uuid_generate_v4() WHERE id=$5 and version=$6 RETURNING version`
	err := u.db.QueryRowContext(ctx, query,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.Activated,
		&user.ID,
		user.Version).Scan(&user.Version)

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

func (u *userRepo) DeleteUser(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := u.db.ExecContext(ctx, query, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return fmt.Errorf("user with id %d not found", id)
		default:
			return fmt.Errorf("failed to delete user: %w", err)
		}
	}
	return nil
}
