package permissions

import (
	"context"

	model "github.com/AbdulwahabNour/movies/internal/model/permission"
)

type Repository interface {
	AddPermission(ctx context.Context, p *model.Permission) error
	GetPermission(ctx context.Context, id int64) (*model.Permission, error)
	UpdatePermission(ctx context.Context, p *model.Permission) error
	DeletePermission(ctx context.Context, id int64) error
	UserPermissionsRepo
}

type UserPermissionsRepo interface {
	UserPermissions(ctx context.Context, userId int64) ([]*model.Permission, error)
	AddUserPermissions(ctx context.Context, userId int64, codes ...string) error
	DeleteUserPermission(ctx context.Context, userId int64, codes ...string) error
}
