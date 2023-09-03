package movies

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateMovieHandler(c *gin.Context)
	ShowMovieHandler(c *gin.Context)
	ListMoviesHandler(c *gin.Context)
	UpdateMovieHandler(c *gin.Context)
	DeleteMovieHandler(c *gin.Context)
}
