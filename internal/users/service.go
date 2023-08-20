package users

import (
	"context"

	model "github.com/AbdulwahabNour/movies/internal/model/users"
)

type Service interface {
	InsertUser(ctx context.Context, user *model.SignUpInput) (*model.User, error)
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int64) error
	SignUp(ctx context.Context, user *model.SignUpInput) (*model.User, error)
	SigIn(ctx context.Context, user *model.SignIn) (*model.UserWithToken, error)
}
