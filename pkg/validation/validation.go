package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/go-playground/validator/v10"
)


type Validatation struct{
    Errors []ValidatationError
}

func NewValidatation() *Validatation{
    return &Validatation{}
}

type ValidatationError struct{
    Field string
    Message string
}

func(ve *Validatation) DescriptiveValidationMessages(err error){
 
    var syntaxError *json.SyntaxError
    var unmarshalTypeError *json.UnmarshalTypeError
    var validationErrors validator.ValidationErrors
 
    switch{


    case errors.As(err, &syntaxError):
            ve.Errors =append(ve.Errors, ValidatationError{Field: "general", Message: fmt.Sprintf("request body contains invalid formed  Json at position %d", syntaxError.Offset)})
    case errors.Is(err, io.ErrUnexpectedEOF):
            ve.Errors = append(ve.Errors, ValidatationError{Field: "general", Message: ""})
    case errors.As(err, &unmarshalTypeError):
         ve.Errors = append(ve.Errors, ValidatationError{Field: unmarshalTypeError.Field,
                                                                   Message: fmt.Sprintf("Invalid type specified for field %s at position %d", unmarshalTypeError.Field, unmarshalTypeError.Offset),})
    case strings.HasPrefix(err.Error(), "json: unknown field "):
        fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
        ve.Errors = append(ve.Errors, ValidatationError{Field: fieldName,
                                                                  Message: fmt.Sprintf("Request body contains unknown field %s", fieldName), }) 
    case errors.Is(err, io.EOF):
         ve.Errors = append(ve.Errors, ValidatationError{Field: "general",
                                                                    Message: "Request body must not be empty"})
    case err.Error() == "http: request body too large":
            ve.Errors = append(ve.Errors, ValidatationError{Field: "general",
                                                                      Message: "Request body must not be larger than 1MB"}) 
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
            ve.Errors = append(ve.Errors, ValidatationError{Field: f.Field(), Message: msg})
          }
    default:
		ve.Errors = append(ve.Errors, ValidatationError{Field: "general", Message: err.Error()})
    }

   
}
 
func (v *Validatation)IsUnique(slice []string) bool{

    seen := make(map[string]bool)

    for _, v := range slice{
        if seen[v]{
            return false
        }
        seen[v]= true
    }
    
    return true
}
func (V *Validatation)NoErrors() bool {
    return len(V.Errors) == 0
}
