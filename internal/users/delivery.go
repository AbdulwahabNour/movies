package users

import "github.com/gin-gonic/gin"

type Handler interface {
	SignUpHandler(c *gin.Context)
	SigInHandler(c *gin.Context)
	GetUserByEmailHandler(c *gin.Context)
	UpdateuUserHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
}
