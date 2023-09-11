package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	model "github.com/AbdulwahabNour/movies/internal/model/permission"
	"github.com/AbdulwahabNour/movies/internal/permissions"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type permissionRepo struct {
	db sqlx.DB
}

func NewPermissionRepo(db *sqlx.DB) permissions.Repository {
	return &permissionRepo{
		db: *db,
	}
}

func (r *permissionRepo) AddPermission(ctx context.Context, p *model.Permission) error {
	query := `INSERT INTO permissions (code) VALUES ($1) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, p.Code).Scan(&p.ID)
	if err != nil {
		return fmt.Errorf("failed to create permission: %w", err)
	}
	return nil
}
func (r *permissionRepo) GetPermission(ctx context.Context, id int64) (*model.Permission, error) {
	query := `SELECT id, code FROM permissions WHERE id= $1`
	var perm model.Permission
	err := r.db.QueryRowContext(ctx, query, id).Scan(&perm.ID, &perm.Code)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, fmt.Errorf("permission with id %d not found", id)
		default:
			return nil, fmt.Errorf("failed to get permission: %w", err)
		}
	}
	return &perm, nil
}

func (r *permissionRepo) UpdatePermission(ctx context.Context, p *model.Permission) error {
	query := `UPDATE permissions SET code = $2 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, p.ID, p.Code)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return fmt.Errorf("permission with id %d not found", p.ID)
		default:
			return fmt.Errorf("failed to get permission: %w", err)
		}
	}
	return nil

}
func (r *permissionRepo) DeletePermission(ctx context.Context, id int64) error {
	query := `DELETE FROM permissions  WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return fmt.Errorf("permission with id %d not found", id)
		default:
			return fmt.Errorf("failed to get permission: %w", err)
		}
	}
	return nil

}
func (r *permissionRepo) UserPermissions(ctx context.Context, userId int64) ([]*model.Permission, error) {
	query := `SELECT permissions.id ,permissions.code 
              FROM permissions
              INNER JOIN users_permissions ON users_permissions.permission_id=permissions.id
              INNER JOIN users ON users_permissions.user_id = users.id 
              WHERE users.id = $1`
	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	perm := make([]*model.Permission, 0)

	for rows.Next() {
		var p model.Permission
		err := rows.Scan(&p.ID, &p.Code)
		if err != nil {
			return nil, err
		}
		perm = append(perm, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return perm, nil
}

func (r *permissionRepo) AddUserPermissions(ctx context.Context, userId int64, codes ...string) error {

	query := `INSERT INTO users_permissions SELECT $1, permissions.id FROM permissions WHERE permissions.code=ANY($2)`
	_, err := r.db.ExecContext(ctx, query, userId, pq.Array(codes))
	if err != nil {
		pqErr, ispGerr := err.(*pq.Error)
		if ispGerr {
			if pqErr.Code == "23503" {
				return fmt.Errorf("foreign key constraint violation")
			}
		}
		return fmt.Errorf("failed to create user permission: %w", err)
	}
	return nil
}

func (r *permissionRepo) DeleteUserPermission(ctx context.Context, userId int64, codes ...string) error {
	query := `DELETE 
	FROM users_permissions 
	WHERE 
	user_id = $1 and permission_id 
	IN(SELECT permissions.id FROM permissions WHERE permissions.code=ANY($2))
	`

	_, err := r.db.ExecContext(ctx, query, userId, pq.Array(codes))
	if err != nil {
		return fmt.Errorf("failed to delete user permission: %w", err)
	}

	return nil
}
