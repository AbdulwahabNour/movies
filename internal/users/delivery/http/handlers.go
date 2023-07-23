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

func (h *apiHandlers) CreateUserHandler(c *gin.Context) {
	var user model.User

	err := utils.ReadRequestJSON(c, &user)
	if err != nil {
		utils.ErrorLogWithFields(h.logger, c, "CreateUserHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	err = h.userService.InsertUser(ctx, &user)
	if err != nil {
		utils.ErrorLogWithFields(h.logger, c, "CreateUserHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusOK, gin.H{"status": "created", "user": user})
}
func (h *apiHandlers) GetUserByEmailHandler(c *gin.Context) {

	email := c.Param("email")

	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()
	user, err := h.userService.GetUserByEmail(ctx, email)

	if err != nil {
		utils.ErrorLogWithFields(h.logger, c, "GetUserByEmailHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusOK, user)
}

func (h *apiHandlers) UpdateuUserHandler(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		utils.ErrorLogWithFields(h.logger, c, "UpdateuUserHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	err = utils.ReadRequestJSON(c, &user)
	if err != nil {
		utils.ErrorLogWithFields(h.logger, c, "UpdateuUserHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()
	user.ID = id
	err = h.userService.UpdateUser(ctx, &user)

	if err != nil {
		utils.ErrorLogWithFields(h.logger, c, "UpdateuUserHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusOK, gin.H{"status": "updated", "user": user})
}
func (h *apiHandlers) DeleteUserHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		utils.ErrorLogWithFields(h.logger, c, "DeleteUserHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()
	err = h.userService.DeleteUser(ctx, id)

	if err != nil {
		utils.ErrorLogWithFields(h.logger, c, "DeleteUserHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	utils.Response(c, http.StatusOK, gin.H{"status": "deleted"})
}
