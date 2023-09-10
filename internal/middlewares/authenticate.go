package middlewares

import (
	"context"
	"net/http"

	permissionModel "github.com/AbdulwahabNour/movies/internal/model/permission"
	Usermodel "github.com/AbdulwahabNour/movies/internal/model/users"
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
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		idToken, err := utils.GetBearerToken(c.Request)
		if err != nil {
			c.Set("user", Usermodel.AnonymousUser)
			c.Next()
			return
		}

		claims, err := utils.ValidateIDToken(idToken, m.config.Server.PublicKeyToken)
		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user", &Usermodel.User{
			ID:    claims.ID,
			Name:  claims.Name,
			Email: claims.Email,
		})
		c.Next()
	}
}

func (m *MiddleWares) RequiredAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := utils.ContextGetUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if user.IsAnonymous() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
func (m *MiddleWares) RequirePermission(code string) gin.HandlerFunc {

	return func(c *gin.Context) {

		user, err := utils.ContextGetUser(c)
		if err != nil || user.IsAnonymous() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		ctx, cancle := context.WithTimeout(context.Background(), m.config.Server.CtxDefaultTimeout)
		defer cancle()

		permissions, err := m.permissionServ.GetUserPermissions(ctx, user.ID)
		if err != nil {
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}

		if !permissionModel.HasCode(permissions, code) {

			c.JSON(http.StatusForbidden, gin.H{"error": "your user account doesn't have permissions to access this resource"})
			c.Abort()
			return

		}
		c.Next()
	}
}
