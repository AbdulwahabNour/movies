package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/AbdulwahabNour/movies/config"
	model "github.com/AbdulwahabNour/movies/internal/model/users"
	"github.com/AbdulwahabNour/movies/internal/users"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/AbdulwahabNour/movies/pkg/utils"
	"github.com/gin-gonic/gin"
)

type apiHandlers struct {
	config      *config.Config
	logger      logger.Logger
	userService users.Service
}

func NewMovieHandlers(app *config.Config, serv users.Service, logger logger.Logger) users.Handler {
	return &apiHandlers{
		config:      app,
		logger:      logger,
		userService: serv,
	}
}

func (h *apiHandlers) SignUpHandler(c *gin.Context) {
	var user model.SignUpInput

	err := utils.ReadRequestJSON(c, &user)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "users.handlers.SignUpHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	_, err = h.userService.SignUp(ctx, &user)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "users.handlers.SignUpHandler.SignUp", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusOK, gin.H{"status": "created"})
}
func (h *apiHandlers) SigInHandler(c *gin.Context) {
	var userLogin model.SignIn

	err := utils.ReadRequestJSON(c, &userLogin)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "users.handlers.SigInHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	userWithtoken, err := h.userService.SigIn(ctx, &userLogin)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "users.handlers.SigInHandler.SigIn", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusOK, gin.H{"status": "success", "data": userWithtoken})
}
func (h *apiHandlers) GetUserByEmailHandler(c *gin.Context) {

	email := c.Param("email")

	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()
	user, err := h.userService.GetUserByEmail(ctx, email)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "users.handlers.GetUserByEmailHandler.GetUserByEmail", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusOK, user)
}

func (h *apiHandlers) UpdateuUserHandler(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "users.handlers.UpdateuUserHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	err = utils.ReadRequestJSON(c, &user)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "users.handlers.UpdateuUserHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()
	user.ID = id
	err = h.userService.UpdateUser(ctx, &user)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "users.handlers.UpdateuUserHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusOK, gin.H{"status": "updated", "user": user})
}
func (h *apiHandlers) DeleteUserHandler(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "users.handlers.DeleteUserHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()
	err = h.userService.DeleteUser(ctx, id)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "users.handlers.DeleteUserHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusOK, gin.H{"status": "deleted"})
}
