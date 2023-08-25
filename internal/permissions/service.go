package permissions

import (
	"context"

	model "github.com/AbdulwahabNour/movies/internal/model/permission"
)

type Service interface {
	AddPermission(ctx context.Context, p *model.Permission) error
	GetPermission(ctx context.Context, id int64) (*model.Permission, error)
	UpdatePermission(ctx context.Context, p *model.Permission) error
	DeletePermission(ctx context.Context, id int64) error
	UserPermissionsService
}

type UserPermissionsService interface {
	GetUserPermissions(ctx context.Context, userId int64) ([]*model.Permission, error)
	SetUserPermission(ctx context.Context, userpermission *model.UserPermission) error
	DeleteUserPermission(ctx context.Context, userpermission *model.UserPermission) error
}
