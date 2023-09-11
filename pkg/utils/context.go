package utils

import (
	"fmt"

	model "github.com/AbdulwahabNour/movies/internal/model/users"
	"github.com/gin-gonic/gin"
)

func ContextGetUser(c *gin.Context) (*model.User, error) {

	users, found := c.Get("user")
	user, ok := users.(*model.User)

	if !found || !ok {
		return nil, fmt.Errorf("Unauthorized")
	}

	return user, nil
}
