package service

import (
	"context"
	"regexp"

	"github.com/AbdulwahabNour/movies/config"
	model "github.com/AbdulwahabNour/movies/internal/model/permission"
	"github.com/AbdulwahabNour/movies/internal/permissions"
	"github.com/AbdulwahabNour/movies/pkg/httpError"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type permissionService struct {
	config   *config.Config
	repo     permissions.Repository
	logger   logger.Logger
	validate *validator.Validate
}

func NewPermissionService(config *config.Config, repo permissions.Repository, logger logger.Logger, validate *validator.Validate) permissions.Service {
	return &permissionService{
		config:   config,
		repo:     repo,
		logger:   logger,
		validate: validate,
	}
}

func (s *permissionService) AddPermission(ctx context.Context, p *model.Permission) error {
	if err := s.validate.Struct(p); err != nil {
		return httpError.ParseValidationErrors(err)
	}

	err := s.repo.AddPermission(ctx, p)
	if err != nil {
		return httpError.ParseErrors(err)
	}
	return nil
}
func (s *permissionService) GetPermission(ctx context.Context, id int64) (*model.Permission, error) {
	perm, err := s.repo.GetPermission(ctx, id)
	if err != nil {
		return nil, httpError.ParseErrors(err)
	}
	return perm, nil
}
func (s *permissionService) UpdatePermission(ctx context.Context, p *model.Permission) error {
	if err := s.validate.Struct(p); err != nil {
		return httpError.ParseValidationErrors(err)
	}
	err := s.repo.UpdatePermission(ctx, p)
	if err != nil {
		return httpError.ParseErrors(err)
	}

	return nil
}
func (s *permissionService) DeletePermission(ctx context.Context, id int64) error {

	err := s.repo.DeletePermission(ctx, id)
	if err != nil {
		return httpError.ParseErrors(err)
	}
	return nil

}
func (s *permissionService) GetUserPermissions(ctx context.Context, userId int64) ([]*model.Permission, error) {
	userPermissions, err := s.repo.UserPermissions(ctx, userId)
	if err != nil {
		return nil, httpError.ParseErrors(err)
	}
	return userPermissions, nil

}

func (s *permissionService) SetUserPermissions(ctx context.Context, userId int64, permissions ...string) error {

	if !isValidPermissions(permissions...) {
		return httpError.NewBadRequestError("invalid permission format")
	}
	if userId < 1 {
		return httpError.NewBadRequestError("user id less than 1")
	}

	err := s.repo.AddUserPermissions(ctx, userId, permissions...)

	if err != nil {
		return httpError.ParseErrors(err)
	}

	return nil
}
func (s *permissionService) DeleteUserPermission(ctx context.Context, userId int64, permissions ...string) error {
	if !isValidPermissions(permissions...) {
		return httpError.NewBadRequestError("invalid permission format")
	}
	if userId < 1 {
		return httpError.NewBadRequestError("user id less than 1")
	}

	err := s.repo.DeleteUserPermission(ctx, userId, permissions...)

	if err != nil {
		return httpError.ParseErrors(err)
	}

	return nil
}

func isValidPermissions(permissions ...string) bool {
	pattern := `^[a-zA-Z]+:[a-zA-Z]+$`
	regex := regexp.MustCompile(pattern)

	for _, permission := range permissions {
		if !regex.MatchString(permission) {
			return false
		}

	}

	return true
}
