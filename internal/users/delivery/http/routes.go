package http

import (
	"github.com/AbdulwahabNour/movies/internal/users"
	"github.com/gin-gonic/gin"
)

func MapUsersRoutes(r *gin.RouterGroup, app users.Handler) {
	r.POST("/users/login", app.SigInHandler)
	r.POST("/users/register", app.SignUpHandler)
	r.GET("/users/:email", app.GetUserByEmailHandler)
	r.PUT("/users/:id", app.UpdateuUserHandler)
	r.DELETE("/users/:id", app.DeleteUserHandler)

}
