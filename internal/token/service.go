package token

import (
	"context"

	modelToken "github.com/AbdulwahabNour/movies/internal/model/token"
	model "github.com/AbdulwahabNour/movies/internal/model/users"
	"github.com/AbdulwahabNour/movies/pkg/utils"
)

type TokenService interface {
	NewPairFromUser(ctx context.Context, u *model.User) (*modelToken.TokenPair, error)
	GenerateActivationToken(ctx context.Context, u *model.User) (*modelToken.Token, error)
	ValidateActivationToken(ctx context.Context, u *model.User, token string) error
	DeleteActivationToken(ctx context.Context, u *model.User) error
	Signout(ctx context.Context, uid int64) error
	ValidateIDToken(tokenString string) (*utils.IDTokenClaims, error)
	ValidateRefreshToken(refreshTokenString string) (*utils.RefreshTokenClaims, error)
}
