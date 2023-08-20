package token

import "github.com/gin-gonic/gin"

type Handler interface {
	Token(c *gin.Context)
	Activate(c *gin.Context)
}
