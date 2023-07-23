package users

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateUserHandler(c *gin.Context)
	GetUserByEmailHandler(c *gin.Context)
	UpdateuUserHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
}
