package service

import (
	"context"
	"fmt"

	"github.com/AbdulwahabNour/movies/config"
	modelToken "github.com/AbdulwahabNour/movies/internal/model/token"
	model "github.com/AbdulwahabNour/movies/internal/model/users"
	"github.com/AbdulwahabNour/movies/internal/token"
	"github.com/AbdulwahabNour/movies/pkg/httpError"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/AbdulwahabNour/movies/pkg/utils"
	"github.com/sirupsen/logrus"
)

var (
	refreshPrefix    = "refreshtoken"
	activationPrefix = "activationtoken"
)

type tokenService struct {
	config          *config.Config
	tokenRepository token.TokenRepository
	logger          logger.Logger
}

func NewTokenService(config *config.Config, logger logger.Logger, tokenRepository token.TokenRepository) token.TokenService {
	return &tokenService{
		config:          config,
		logger:          logger,
		tokenRepository: tokenRepository,
	}
}

func (s *tokenService) NewPairFromUser(ctx context.Context, u *model.User) (*modelToken.TokenPair, error) {

	idTokenClaims := utils.IDTokenClaims{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
	idToken, err := utils.GenerateIDToken(idTokenClaims, s.config.Server.PrivateKeyToken, s.config.Server.IDExpiration)
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := utils.RefreshTokenClaims{
		UID: fmt.Sprintf("%d", u.ID),
	}

	refreshToken, err := utils.GenerateRefreshToken(refreshTokenClaims, s.config.Server.JwtSecretKey, s.config.Server.RefreshExpiratio)
	if err != nil {
		return nil, err
	}

	if err := s.tokenRepository.SetToken(ctx, u.ID, refreshToken.ID, refreshPrefix, refreshToken.ExpIn); err != nil {
		return nil, err
	}

	return &modelToken.TokenPair{
		IDToken: modelToken.IDToken{Token: idToken},
		RefreshToken: modelToken.RefreshToken{ID: refreshToken.ID,
			UID:   refreshTokenClaims.UID,
			Token: refreshToken.Token},
	}, nil
}
func (s *tokenService) GenerateActivationToken(ctx context.Context, u *model.User) (*modelToken.Token, error) {
	token, err := utils.RandToken()
	if err != nil {
		return nil, err
	}
	fmt.Println(token)
	err = s.tokenRepository.SetToken(ctx, u.ID, token.Hash, activationPrefix, s.config.Server.ActivationTokenExpiratio) //activationtoken:id->rand token
	if err != nil {
		return nil, err
	}

	return &modelToken.Token{
		Plaintext: token.Plaintext,
		Hash:      token.Hash,
	}, nil
}
func (s *tokenService) ValidateActivationToken(ctx context.Context, u *model.User, token string) error {

	dbToken, err := s.tokenRepository.GetToken(ctx, u.ID, activationPrefix)
	if err != nil {
		return err
	}
	hashUserToken := utils.Hash(token)

	if dbToken != hashUserToken {
		return httpError.NewUnAuthorizedError("token not valid ")
	}
	return nil
}
func (s *tokenService) DeleteActivationToken(ctx context.Context, u *model.User) error {
	if err := s.tokenRepository.DeleteToken(ctx, u.ID, activationPrefix); err != nil {
		return httpError.NewInternalServerError(err)
	}
	return nil
}
func (s *tokenService) DeleteUserTokens(ctx context.Context, uid int64) error {
	if err := s.tokenRepository.DeleteToken(ctx, uid, refreshPrefix); err != nil {
		return httpError.NewInternalServerError(err)
	}
	return nil
}
func (s *tokenService) ValidateIDToken(tokenString string) (*utils.IDTokenClaims, error) {

	claims, err := utils.ValidateIDToken(tokenString, s.config.Server.PublicKeyToken)
	if err != nil {
		s.logger.ErrorLogWithFields(logrus.Fields{"err": err, "method": "token.service.ValidateIDToken"}, "unable to validate ID Token")
		return nil, httpError.NewUnAuthorizedError(err)
	}

	return claims, nil
}
func (s *tokenService) ValidateRefreshToken(refreshTokenString string) (*utils.RefreshTokenClaims, error) {

	claims, err := utils.ValidateRefreshToken(refreshTokenString, s.config.Server.JwtSecretKey)
	if err != nil {
		s.logger.ErrorLogWithFields(logrus.Fields{"err": err, "method": "token.service.ValidateRefreshToken "}, "unable to validate Refresh Token")
		return nil, httpError.NewUnAuthorizedError(err)
	}

	return claims, nil
}
