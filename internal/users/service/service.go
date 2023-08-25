package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/AbdulwahabNour/movies/config"
	"github.com/AbdulwahabNour/movies/internal/mailer"
	model "github.com/AbdulwahabNour/movies/internal/model/users"
	"github.com/AbdulwahabNour/movies/internal/token"
	"github.com/AbdulwahabNour/movies/internal/users"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	"github.com/AbdulwahabNour/movies/pkg/httpError"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/AbdulwahabNour/movies/pkg/utils"
)

type userService struct {
	config    *config.Config
	repo      users.Repository
	tokenServ token.TokenService
	logger    logger.Logger
	validate  *validator.Validate
}

func NewUserService(config *config.Config,
	repo users.Repository,
	tokenServ token.TokenService,
	logger logger.Logger,
	validate *validator.Validate) users.Service {

	return &userService{
		config:    config,
		repo:      repo,
		tokenServ: tokenServ,
		logger:    logger,
		validate:  validate,
	}
}

func (s *userService) SignUp(ctx context.Context, user *model.SignUpInput) (*model.User, error) {

	insertedUser, err := s.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}
	s.sendActivateToken(insertedUser)
	return insertedUser, nil
}

func (s *userService) SigIn(ctx context.Context, userSignIn *model.SignIn) (*model.UserWithToken, error) {

	if err := s.validate.Struct(userSignIn); err != nil {
		return nil, httpError.ParseValidationErrors(err)
	}

	user, err := s.GetUserByEmail(ctx, userSignIn.Email)
	if err != nil {
		return nil, httpError.NewUnAuthorizedError("wrong email")
	}

	err = utils.VerifyPassword(user.HashedPassword, userSignIn.Password, s.config.Server.PepperSecreKey)
	if err != nil {
		return nil, httpError.NewUnAuthorizedError("wrong password")
	}

	if !*user.Activated {
		return nil, httpError.NewHttpError(http.StatusForbidden, "can't login", "user not activated")
	}

	token, err := s.tokenServ.NewPairFromUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &model.UserWithToken{
		User:  user,
		Token: token,
	}, nil
}
func (s *userService) InsertUser(ctx context.Context, user *model.SignUpInput) (*model.User, error) {

	if err := user.Check(); err != nil {
		return nil, httpError.NewBadQueryError(err)
	}
	newuser := user.Map()

	pass, err := utils.HashPassword(newuser.Password, s.config.Server.PepperSecreKey)
	if err != nil {
		return nil, httpError.NewInternalServerError(err)
	}

	newuser.HashedPassword = pass
	newuser.Activated = new(bool)
	*newuser.Activated = false

	if err := s.validate.Struct(newuser); err != nil {
		return nil, httpError.ParseValidationErrors(err)
	}

	err = s.repo.InsertUser(ctx, &newuser)

	if err != nil {
		return nil, httpError.ParseErrors(err)
	}

	return &newuser, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, httpError.ParseErrors(err)
	}
	return user, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, strings.TrimSpace(email))
	if err != nil {

		return nil, httpError.ParseErrors(err)
	}
	return user, err
}
func (s *userService) UpdateUser(ctx context.Context, user *model.User) error {
	userDb, err := s.GetUserByID(ctx, user.ID)
	if err != nil {
		return err
	}
	if user.Password != "" {
		userDb.HashedPassword, err = utils.HashPassword(user.Password, s.config.Server.PepperSecreKey)
		if err != nil {
			return httpError.NewInternalServerError(err)
		}
	}
	if user.Email != "" {
		userDb.Email = user.Email
	}
	if user.Name != "" {
		userDb.Name = user.Name
	}
	if user.Activated != nil {
		userDb.Activated = user.Activated
	}
	user.SanitizePassword()
	err = s.repo.UpdateUser(ctx, userDb)
	if err != nil {
		return httpError.ParseErrors(err)
	}
	return nil
}
func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	if id < 1 {
		return fmt.Errorf("user not found")
	}
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return httpError.ParseErrors(err)
	}
	return nil
}
func (s *userService) sendActivateToken(user *model.User) {
	utils.BackgroundWithRecover(s.logger, func() {
		ctxE, cancle := context.WithTimeout(context.Background(), s.config.Server.CtxDefaultTimeout)
		defer cancle()
		activateToken, err := s.tokenServ.GenerateActivationToken(ctxE, user)
		if err != nil {
			s.logger.ErrorLogWithFields(logrus.Fields{"err": err, "userID": user.ID}, "error happened during GenerateActivationToken")
			return
		}

		data := mailer.MailerData{
			Data: map[string]interface{}{"user": user,
				"appName":      s.config.Server.AppName,
				"activatelink": fmt.Sprintf("%s/activate?id=%d&token=%s", s.config.Server.AppHost, user.ID, activateToken.Plaintext)},
			Recipient: user.Email,
		}

		sendmail := mailer.NewMailer(s.config, "signup")

		err = sendmail.Send(data)
		if err != nil {
			s.logger.ErrorLogWithFields(logrus.Fields{"err": err, "userID": user.ID}, "error happened during send email")
		}
	})
}
