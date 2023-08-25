package http

import (
	"github.com/AbdulwahabNour/movies/internal/middlewares"
	"github.com/AbdulwahabNour/movies/internal/permissions"
	"github.com/gin-gonic/gin"
)

func MapMoviesRoutes(r *gin.RouterGroup, app permissions.Handler, mw *middlewares.MiddleWares) {

	g := r.Group("/permission/")

	g.POST("/", app.AddPermissionHandler)  //done
	g.GET("/:id", app.GetPermissioHandler) //done
	g.PUT("/", mw.RequiredAuth(), app.UpdatePermissionHandler)
	g.DELETE("/:id", mw.RequiredAuth(), app.DeletePermissionHandler)
	g.GET("/user/:id", app.GetUserPermissionsHandler)
	g.POST("/user/", app.SetUserPermissionHandler)
	g.DELETE("/user/", mw.RequiredAuth(), app.DeleteUserPermissionHandler)

}
