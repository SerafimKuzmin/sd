package middleware

import (
	friendUsecase "github.com/SerafimKuzmin/sd/backend/internal/List/usecase"
	"github.com/SerafimKuzmin/sd/backend/models"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

type AclMiddleware struct {
	friendUC friendUsecase.UsecaseI
}

func NewAclMiddleware(friendUC friendUsecase.UsecaseI) *AclMiddleware {
	return &AclMiddleware{friendUC: friendUC}
}

func (am *AclMiddleware) AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*models.User)
		if user.Role != 2 {
			return echo.NewHTTPError(http.StatusForbidden, models.ErrPermissionDenied.Error())
		}
		return next(c)
	}
}
