package utils

import (
	"github.com/AbdulwahabNour/movies/pkg/httpError"
	"github.com/AbdulwahabNour/movies/pkg/logger"
	"github.com/gin-gonic/gin"
)

func ReadRequestJSON(c *gin.Context, v any) error {
	return c.ShouldBindJSON(v)
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

func ErrorResponse(c *gin.Context, err error) {

	m := httpError.ParseErrors(err)

	// c.AbortWithError(http.StatusBadRequest, m)
	c.JSON(m.Status(), gin.H{"error": m})

}

// func MethodNotAllowedResponse(c *gin.Context) {
//     // Generate error message with the HTTP method used
//     errorMessage := fmt.Sprintf("the %s method is not supported for this resource", c.Request.Method)

//     // Send JSON response with error message and HTTP status code
//     c.JSON(http.StatusMethodNotAllowed, gin.H{"error": errorMessage})
// }
// func NnotfoundResponse(c *gin.Context){
//     c.JSON(http.StatusNotFound, gin.H{"error": "the requested resource could not be found"})
// }
