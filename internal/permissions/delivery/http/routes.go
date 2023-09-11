package http

import (
	"github.com/AbdulwahabNour/movies/internal/middlewares"
	"github.com/AbdulwahabNour/movies/internal/permissions"
	"github.com/gin-gonic/gin"
)

func MapMoviesRoutes(r *gin.RouterGroup, app permissions.Handler, mw *middlewares.MiddleWares) {

	g := r.Group("/permission/")

	g.POST("/", mw.RequirePermission("add:permission"), app.AddPermissionHandler)
	g.GET("/:id", mw.RequirePermission("view:permission"), app.GetPermissioHandler)
	g.PUT("/:id", mw.RequirePermission("update:permission"), app.UpdatePermissionHandler)
	g.DELETE("/:id", mw.RequirePermission("delete:permission"), app.DeletePermissionHandler)
	g.GET("/user/:id", mw.RequirePermission("view:userpermission"), app.GetUserPermissionsHandler)
	g.POST("/user/:id", mw.RequirePermission("add:userpermission"), app.SetUserPermissionHandler)
	g.DELETE("/user/:id", mw.RequirePermission("delete:userpermission"), app.DeleteUserPermissionHandler)

}
