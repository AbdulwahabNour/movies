package http

import (
	"github.com/AbdulwahabNour/movies/internal/middlewares"
	"github.com/AbdulwahabNour/movies/internal/users"
	"github.com/gin-gonic/gin"
)

func MapUsersRoutes(r *gin.RouterGroup, app users.Handler, mw *middlewares.MiddleWares) {

	r.POST("auth/users/login", app.SigInHandler)
	r.POST("auth/users/register", app.SignUpHandler)

	r.GET("/users/:email", mw.RequirePermission("read:user"), app.GetUserByEmailHandler)
	r.PUT("/users/:id", mw.RequirePermission("update:user"), app.UpdateuUserHandler)
	r.DELETE("/users/:id", mw.RequirePermission("delete:user"), app.DeleteUserHandler)

}
