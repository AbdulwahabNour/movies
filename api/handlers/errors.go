package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *apiHandlers) logError(c *gin.Context, err error){
    h.App.ErrorLog.Println(err)
}
func (h *apiHandlers) errorResponse( c *gin.Context, status int, message any){
    c.JSON(status, message)
}

func(h *apiHandlers) serverErrorResponse(c *gin.Context, message any){
    c.JSON(http.StatusInternalServerError, message)
   
}
func(h *apiHandlers) methodNotAllowedResponse(c *gin.Context){
    m:= fmt.Sprintf("the %s method is not supported for this  resource", c.Request.Method)
    c.JSON(http.StatusMethodNotAllowed, gin.H{"error": m})
}
func(h *apiHandlers) notfoundResponse(c *gin.Context){
    c.JSON(http.StatusNotFound, gin.H{"error": "the requested resource could not be found"})
}

 