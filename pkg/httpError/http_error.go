package httpError

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	ErrNotFound          = errors.New("not Found")
	ErrBadQuery          = errors.New("invalid Query parameter")
	ErrBadRequest        = errors.New("bad request")
	ErrRequestTimeout    = errors.New("request Timeout")
	ErrInternalServer    = errors.New("internal Server Error")
	ErrUnSupportedEntity = errors.New("unsupported Entity")
	ErrInvalidSyntax     = errors.New("invalid syntax")
	ErrRecordNotFound    = errors.New("record not found")
	ErrInvalidJsonFormat = errors.New("request body contains invalid formed  Json")
	ErrUnexpectedEOF     = errors.New("an unexpected end of input occurred. The data provided is incomplete or truncated")
	ErrEditConflict      = errors.New("unable to update the record due to an edit conflict, please try again")
	ErrDuplicateEmail    = errors.New("user already exist")
)

type HttpErr interface {
	Status() int
	Error() string
	Description() any
}

type HttpError struct {
	ErrStatus      int    `json:"status"`
	ErrError       string `json:"error"`
	ErrDescription any    `json:"description"`
}

func (e HttpError) Status() int {
	return e.ErrStatus
}
func (e HttpError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - description: %v", e.ErrStatus, e.ErrError, e.ErrDescription)
}

func (e HttpError) Description() any {
	return e.ErrDescription
}
func NewHttpError(status int, err string, description any) HttpErr {
	if e, ok := description.(error); ok {
		description = e.Error()
	}
	return HttpError{
		ErrStatus:      status,
		ErrError:       err,
		ErrDescription: description}
}

func NewBadRequestError(err any) HttpErr {
	if e, ok := err.(error); ok {
		err = e.Error()
	}

	return HttpError{
		ErrStatus:      http.StatusBadRequest,
		ErrError:       ErrBadRequest.Error(),
		ErrDescription: err,
	}
}

func NewNotFoundError(err any) HttpErr {
	if e, ok := err.(error); ok {
		err = e.Error()
	}

	return HttpError{
		ErrStatus:      http.StatusNotFound,
		ErrError:       ErrNotFound.Error(),
		ErrDescription: err,
	}
}

func NewBadQueryError(err any) HttpErr {
	if e, ok := err.(error); ok {
		err = e.Error()
	}
	return HttpError{
		ErrStatus:      http.StatusBadRequest,
		ErrError:       ErrBadQuery.Error(),
		ErrDescription: err,
	}
}

func NewInternalServerError(err any) HttpErr {

	if e, ok := err.(error); ok {
		err = e.Error()
	}

	return HttpError{
		ErrStatus:      http.StatusInternalServerError,
		ErrError:       ErrInternalServer.Error(),
		ErrDescription: err,
	}
}
func NewUnprocessableEntityError(err any) HttpErr {
	return HttpError{
		ErrStatus:      http.StatusUnprocessableEntity,
		ErrError:       ErrUnSupportedEntity.Error(),
		ErrDescription: err,
	}
}
func ParseErrors(err error) HttpErr {

	if httperr, ok := err.(HttpErr); ok {
		return httperr
	}

	var unmarshalTypeError *json.UnmarshalTypeError

	switch {

	case errors.As(err, new(*json.SyntaxError)):
		return NewHttpError(http.StatusBadRequest, ErrInvalidJsonFormat.Error(), err.Error())

	case errors.Is(err, io.ErrUnexpectedEOF):
		return NewHttpError(http.StatusBadRequest, ErrUnexpectedEOF.Error(), err.Error())

	case errors.As(err, &unmarshalTypeError):
		return NewHttpError(http.StatusBadRequest, ErrInvalidSyntax.Error(), fmt.Sprintf("invalid type specified for field %s at position %d", unmarshalTypeError.Field, unmarshalTypeError.Offset))

	case strings.HasPrefix(err.Error(), "json: unknown field"):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
		return NewHttpError(http.StatusBadRequest, fmt.Sprintf("Request body contains unknown field %s", fieldName), err.Error())

	case errors.Is(err, io.EOF):
		return NewHttpError(http.StatusBadRequest, "request body must not be empty", err.Error())

	case err.Error() == "http: request body too large":
		return NewHttpError(http.StatusRequestEntityTooLarge, "request body must not be larger than 1MB", err.Error())

	case errors.Is(err, context.DeadlineExceeded):
		return NewHttpError(http.StatusRequestTimeout, ErrRequestTimeout.Error(), err.Error())

	case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
		return NewHttpError(http.StatusConflict, ErrDuplicateEmail.Error(), err.Error())

	case strings.Contains(err.Error(), "strconv.ParseInt"):

		value := strings.Trim(strings.TrimSuffix(strings.TrimPrefix(err.Error(), "strconv.ParseInt: parsing "), ": invalid syntax"), `\"`)

		return NewHttpError(http.StatusUnprocessableEntity, ErrInvalidSyntax.Error(), fmt.Sprintf("this value ` %s ` is invalid", value))

	case strings.Contains(err.Error(), "violates check constraint"):
		return NewHttpError(http.StatusUnprocessableEntity, err.Error(), ErrUnSupportedEntity.Error())
	case strings.Contains(err.Error(), "edit conflict"):
		return NewHttpError(http.StatusConflict, ErrEditConflict.Error(), err)

	default:

		return NewInternalServerError(err.Error())
	}

}

func ParseValidationErrors(err error) HttpErr {

	errorsSlice := make(map[string]string, 0)
	var validationErrors validator.ValidationErrors

	switch {
	case errors.As(err, &validationErrors):
		for _, f := range validationErrors {
			msg := "Unkown error"

			switch f.ActualTag() {
			case "required":
				msg = fmt.Sprintf("The field %s Required", f.Field())
			case "email":
				msg = "Should be a valid email address"
			case "lte":
				msg = fmt.Sprintf("%s should be less than %s", f.Field(), f.Param())
			case "gte":
				msg = fmt.Sprintf("%s should be greater than %s", f.Field(), f.Param())
			case "alpha":
				msg = "Should be alpha characters only"
			case "unique":
				msg = fmt.Sprintf("%s should be unique", f.Field())
			case "numeric":
				msg = "Should be numbers only"
			case "oneof":
				msg = fmt.Sprintf("Should contain one of values %s", f.Param())
			case "url":
				msg = "Should be valid web address starting with http(s)://..."
			case "min":
				msg = fmt.Sprintf("%s should be minimum %s characters long", f.Field(), f.Param())
			case "max":
				msg = fmt.Sprintf("%s should be maximum %s characters long", f.Field(), f.Param())
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
	return NewHttpError(http.StatusBadRequest, ErrInvalidSyntax.Error(), errorsSlice)
}

func ErrorResponse(err error) (int, HttpErr) {
	e := ParseErrors(err)
	return e.Status(), e
}
