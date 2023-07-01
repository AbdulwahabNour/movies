package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/go-playground/validator/v10"
)

 

func GetDescriptiveMessages(err error) map[string]string {

    if err == nil {
        return nil
    }

    errorsSlice := make(map[string]string, 0)
    var syntaxError *json.SyntaxError
    var unmarshalTypeError *json.UnmarshalTypeError
    var validationErrors validator.ValidationErrors
 
    switch{
    case errors.As(err, &syntaxError):
           errorsSlice["general"] = fmt.Sprintf("request body contains invalid formed  Json at position %d", syntaxError.Offset)
    case errors.Is(err, io.ErrUnexpectedEOF):
          errorsSlice["general"] = "request body must not be empty" 
    case errors.As(err, &unmarshalTypeError):
         errorsSlice["general"] = fmt.Sprintf("Invalid type specified for field %s at position %d", unmarshalTypeError.Field, unmarshalTypeError.Offset) 
    case strings.HasPrefix(err.Error(), "json: unknown field "):
        fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
        errorsSlice["general"] =  fmt.Sprintf("Request body contains unknown field %s", fieldName)   
    case errors.Is(err, io.EOF):
        errorsSlice["general"] =   "request body must not be empty" 
    case err.Error() == "http: request body too large":
           errorsSlice["general"] =   "request body must not be larger than 1MB" 
    case errors.As(err, &validationErrors):
            for _, f := range validationErrors{
            msg := "Unkown error"

            switch f.ActualTag() {
                        case "required":
                            msg = fmt.Sprintf("The field %s Required", f.Field())
                        case "email":
                            msg = "Should be a valid email address"
                        case "lte":
                            msg = fmt.Sprintf("%s should be less than %s",f.Field(), f.Param())
                        case "gte":
                            msg = fmt.Sprintf("%s should be greater than %s",f.Field(), f.Param())
                        case "alpha":
                            msg = "Should be alpha characters only"
                        case "numeric":
                            msg = "Should be numbers only"
                        case "oneof":
                            msg = fmt.Sprintf("Should contain one of values %s", f.Param())
                        case "url":
                            msg = "Should be valid web address starting with http(s)://..."
                        case "min":
                            msg = fmt.Sprintf("%s should be minimum %s characters long",f.Field() ,f.Param())
                        case "max":
                            msg = fmt.Sprintf("%s should be maximum %s characters long",f.Field() ,f.Param())
                        case "e164":
                            msg = "Should be valid phone number"
                        case "datetime":
                            msg = fmt.Sprintf("Should be valid Date/Time with %s format", f.Param())
                    }
                    errorsSlice[f.Field()] = msg
            }
            default:
                   errorsSlice["general"] = err.Error()
    }

   return errorsSlice
}
  