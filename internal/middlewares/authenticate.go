package middlewares

import (
	"net/http"

	model "github.com/AbdulwahabNour/movies/internal/model/users"
	"github.com/AbdulwahabNour/movies/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHeader struct {
	IDToken string `header:"Authorization"`
}

func (m *MiddleWares) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Vary", "Authorization")

		h := AuthHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		idToken, err := utils.GetBearerToken(c.Request)
		if err != nil {
			c.Set("user", model.AnonymousUser)
			c.Next()
			return
		}

		claims, err := utils.ValidateIDToken(idToken, m.config.Server.PublicKeyToken)
		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user", &model.User{
			ID:    claims.ID,
			Name:  claims.Name,
			Email: claims.Email,
		})
		c.Next()
	}
}

func (m *MiddleWares) RequiredAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, found := c.Get("user")
		user, ok := users.(*model.User)
		if !found || !ok || user.IsAnonymous() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
