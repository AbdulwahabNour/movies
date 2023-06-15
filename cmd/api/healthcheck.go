package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)




func(app *application)healthCheckHandler(c *gin.Context){
    c.IndentedJSON(http.StatusOK, gin.H{"status":"available", "enviroment": app.config.env, "version": version})

}