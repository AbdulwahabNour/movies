package permissions

import "github.com/gin-gonic/gin"

type Handler interface {
	AddPermissionHandler(c *gin.Context)
	GetPermissioHandler(c *gin.Context)
	UpdatePermissionHandler(c *gin.Context)
	DeletePermissionHandler(c *gin.Context)
	GetUserPermissionsHandler(c *gin.Context)
	SetUserPermissionHandler(c *gin.Context)
	DeleteUserPermissionHandler(c *gin.Context)
}
