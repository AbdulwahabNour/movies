package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/AbdulwahabNour/movies/config"
	model "github.com/AbdulwahabNour/movies/internal/model/permission"
	"github.com/AbdulwahabNour/movies/internal/permissions"

	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/AbdulwahabNour/movies/pkg/utils"
	"github.com/gin-gonic/gin"
)

type apiHandlers struct {
	config             *config.Config
	permissionsService permissions.Service
	logger             logger.Logger
}

func NewPermissionsHandlers(app *config.Config, serv permissions.Service, logger logger.Logger) permissions.Handler {
	return &apiHandlers{
		config:             app,
		permissionsService: serv,
		logger:             logger,
	}
}

func (h *apiHandlers) AddPermissionHandler(c *gin.Context) {
	var p model.Permission
	if err := utils.ReadRequestJSON(c, &p); err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.AddPermissionHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	err := h.permissionsService.AddPermission(ctx, &p)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.AddPermissionHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	utils.Response(c, http.StatusCreated, p)
}

func (h *apiHandlers) GetPermissioHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.GetPermissioHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()
	perm, err := h.permissionsService.GetPermission(ctx, id)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.GetPermissioHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	utils.Response(c, http.StatusOK, perm)
}
func (h *apiHandlers) UpdatePermissionHandler(c *gin.Context) {

	var p model.Permission
	if err := utils.ReadRequestJSON(c, &p); err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.UpdatePermissionHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	err := h.permissionsService.UpdatePermission(ctx, &p)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.UpdatePermissionHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	utils.Response(c, http.StatusOK, p)

}
func (h *apiHandlers) DeletePermissionHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.DeletePermissionHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	err = h.permissionsService.DeletePermission(ctx, id)

	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.DeletePermissionHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	utils.Response(c, http.StatusOK, gin.H{"status": "deleted"})

}
func (h *apiHandlers) GetUserPermissionsHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.GetUserPermissionsHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	permissions, err := h.permissionsService.GetUserPermissions(ctx, id)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.GetUserPermissionsHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	utils.Response(c, http.StatusOK, gin.H{"permissions": permissions})
}
func (h *apiHandlers) SetUserPermissionHandler(c *gin.Context) {
	var p model.UserPermission
	if err := utils.ReadRequestJSON(c, &p); err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.SetUserPermissionHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()
	err := h.permissionsService.SetUserPermission(ctx, &p)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.SetUserPermissionHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	utils.Response(c, http.StatusOK, gin.H{"status": "created"})
}

func (h *apiHandlers) DeleteUserPermissionHandler(c *gin.Context) {

	var p model.UserPermission

	if err := utils.ReadRequestJSON(c, &p); err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.DeleteUserPermissionHandler", err)
		utils.ErrorResponse(c, err)
		return
	}

	ctx, cancle := context.WithTimeout(context.Background(), h.config.Server.CtxDefaultTimeout)
	defer cancle()

	err := h.permissionsService.DeleteUserPermission(ctx, &p)
	if err != nil {
		utils.GinErrorLogWithFields(h.logger, c, "permissions.handlers.SetUserPermissionHandler", err)
		utils.ErrorResponse(c, err)
		return
	}
	utils.Response(c, http.StatusOK, gin.H{"status": "deleted"})
}
