package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/AbdulwahabNour/movies/config"
	model "github.com/AbdulwahabNour/movies/internal/model/users"
	"github.com/AbdulwahabNour/movies/internal/token"
	"github.com/AbdulwahabNour/movies/internal/users"
	"github.com/AbdulwahabNour/movies/pkg/httpError"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/AbdulwahabNour/movies/pkg/utils"
	"github.com/gin-gonic/gin"
)

type apiHandlers struct {
	config      *config.Config
	tokenServ   token.TokenService
	userService users.Service
	logger      logger.Logger
}

func NewTokenHandlers(config *config.Config, tokenServ token.TokenService, userService users.Service, logger logger.Logger) token.Handler {
	return &apiHandlers{
		config:      config,
		tokenServ:   tokenServ,
		userService: userService,
		logger:      logger,
	}
}

func (h *apiHandlers) Token(c *gin.Context) {

}
func (h *apiHandlers) Activate(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "token.handlers.Activate", err)
		utils.ErrorResponse(c, err)
		return
	}
	token := c.Query("token")

	if token == "" {
		utils.ErrorResponse(c, httpError.NewBadRequestError("token not found"))
		return
	}
	user := model.User{
		ID: userID,
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	err = h.tokenServ.ValidateActivationToken(ctx, &user, token)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "token.handlers.Activate.ValidateActivationToken", err)
		utils.ErrorResponse(c, err)
		return
	}
	user.Activated = new(bool)
	*user.Activated = true

	err = h.userService.UpdateUser(ctx, &user)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "token.handlers.Activate.UpdateUser", err)
		utils.ErrorResponse(c, err)
		return
	}
	err = h.tokenServ.DeleteActivationToken(ctx, &user)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "token.handlers.Activate.DeleteActivationToken", err)
	}

	utils.Response(c, http.StatusOK, gin.H{"status": "Activated"})
}
