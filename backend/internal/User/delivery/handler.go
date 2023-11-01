package delivery

import (
	echo "github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	userUsecase "github.com/SerafimKuzmin/sd/backend/internal/User/usecase"
	"github.com/SerafimKuzmin/sd/backend/internal/middleware"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/SerafimKuzmin/sd/backend/models/dto"
	"github.com/SerafimKuzmin/sd/backend/pkg"
)

type Delivery struct {
	UserUC userUsecase.UsecaseI
}

func (del *Delivery) GetUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}
	user, err := del.UserUC.GetUser(id)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respUser := dto.GetResponseFromModelUser(user)

	return c.JSON(http.StatusOK, pkg.Response{Body: respUser})
}

func (del *Delivery) GetUsers(c echo.Context) error {
	users, err := del.UserUC.GetUsers()

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respUsers := dto.GetResponseFromModelUsers(users)

	return c.JSON(http.StatusOK, pkg.Response{Body: respUsers})
}

func (del *Delivery) GetMe(c echo.Context) error {
	user, ok := c.Get("user").(*models.User)
	if !ok {
		c.Logger().Error("can't get user from context")
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	respUser := dto.GetResponseFromModelUser(user)
	return c.JSON(http.StatusOK, pkg.Response{Body: respUser})
}

func (del *Delivery) UpdateUser(c echo.Context) error {
	var reqUser dto.ReqUpdateUser
	err := c.Bind(&reqUser)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqUser); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	user, ok := c.Get("user").(*models.User)
	if !ok {
		c.Logger().Error("can't get user from context")
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	modelUser := reqUser.ToModelUser()
	modelUser.ID = user.ID

	err = del.UserUC.UpdateUser(modelUser)
	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}
	return c.NoContent(http.StatusNoContent)
}

func handleError(err error) *echo.HTTPError {
	causeErr := errors.Cause(err)
	switch {
	case errors.Is(causeErr, models.ErrNotFound):
		return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
	case errors.Is(causeErr, models.ErrBadRequest):
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	case errors.Is(causeErr, models.ErrPermissionDenied):
		return echo.NewHTTPError(http.StatusForbidden, models.ErrPermissionDenied.Error())
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, causeErr.Error())
	}
}

func NewDelivery(e *echo.Echo, uc userUsecase.UsecaseI, aclM *middleware.AclMiddleware) {
	handler := &Delivery{
		UserUC: uc,
	}

	e.GET("/users/:user_id", handler.GetUser, aclM.AdminOnly)
	e.GET("/me", handler.GetMe)
	e.GET("/users", handler.GetUsers, aclM.AdminOnly)
	e.PUT("/me/edit", handler.UpdateUser)
}
