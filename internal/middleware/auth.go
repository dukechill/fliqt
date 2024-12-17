package middleware

import (
	"net/http"
	"slices"

	"fliqt/internal/model"
	"fliqt/internal/service"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService service.AuthServiceInterface, allowedRoles []model.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := authService.CurrentUser(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		if !slices.Contains(allowedRoles, user.Role) {
			c.AbortWithError(http.StatusForbidden, ErrForbidden)
			return
		}

		c.Next()
	}
}
