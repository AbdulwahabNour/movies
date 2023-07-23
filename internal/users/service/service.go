package service

import (
	"context"

	"github.com/AbdulwahabNour/movies/config"
	model "github.com/AbdulwahabNour/movies/internal/model/users"
	"github.com/AbdulwahabNour/movies/internal/users"
	"github.com/AbdulwahabNour/movies/pkg/httpError"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	config   *config.Config
	repo     users.Repository
	logger   logger.Logger
	validate *validator.Validate
}

func NewUserService(config *config.Config, repo users.Repository, logger logger.Logger, validate *validator.Validate) users.Service {
	return &userService{
		config:   config,
		repo:     repo,
		logger:   logger,
		validate: validate,
	}
}

func (s *userService) InsertUser(ctx context.Context, user *model.User) error {
	if err := s.validate.Struct(user); err != nil {
		return httpError.ParseValidationErrors(err)
	}

	err := user.SetHashPassword()
	if err != nil {
		return httpError.NewBadRequestError(err)
	}
	return s.repo.InsertUser(ctx, user)
}
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {

	return nil, nil
}
func (s *userService) UpdateUser(ctx context.Context, user *model.User) error {

	return nil
}
func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return nil
}
func (s *userService) checkUser(user *model.User) error {
	return nil
}
