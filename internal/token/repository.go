package token

import (
	"context"
	"time"
)

type TokenRepository interface {
	SetToken(ctx context.Context, userID int64, token string, prefix string, expiresIn time.Duration) error
	UpdateToken(ctx context.Context, userID int64, token string, prefix string, expiresIn time.Duration) error
	GetToken(ctx context.Context, userID int64, prefix string) (string, error)
	DeleteToken(ctx context.Context, userID int64, prefix string) error
}
