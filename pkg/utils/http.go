package utils

import (
	"fmt"
	"net/http"

	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/AbdulwahabNour/movies/pkg/validation"
	"github.com/gin-gonic/gin"
)

// ReadRequestJSON reads and binds the JSON data from the HTTP request body to the given structure.
// It returns a map of descriptive error messages if there is an error in binding, otherwise it returns nil.
func ReadRequestJSON(c *gin.Context, v any) map[string]string {
    // Attempt to bind the JSON data to the given structure.
    if err := c.ShouldBindJSON(v); err != nil {
        // If there is an error, return a map of descriptive error messages.
        return validation.GetDescriptiveMessages(err)
    }
  
    // If there is no error, return nil.
    return nil
}
// Refactored function to log an error using the provided logger.
func LogError(c *gin.Context, log logger.Logger, err error) {
    // Call the ErrorLog method of the logger, passing in the error.
    log.ErrorLog(err)
}
// Response is a function that sends a JSON response with the given status code and message.
func Response(c *gin.Context, status int, message any) {
    c.JSON(status, message)
}
 
func MethodNotAllowedResponse(c *gin.Context) {
    // Generate error message with the HTTP method used
    errorMessage := fmt.Sprintf("the %s method is not supported for this resource", c.Request.Method)

    // Send JSON response with error message and HTTP status code
    c.JSON(http.StatusMethodNotAllowed, gin.H{"error": errorMessage})
}
func NnotfoundResponse(c *gin.Context){
    c.JSON(http.StatusNotFound, gin.H{"error": "the requested resource could not be found"})
}

 