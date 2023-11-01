package delivery

import (
	PersonalRatingUsecase "github.com/SerafimKuzmin/sd/backend/internal/PersonalRating/usecase"
	"github.com/SerafimKuzmin/sd/backend/internal/middleware"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/SerafimKuzmin/sd/backend/models/dto"
	"github.com/SerafimKuzmin/sd/backend/pkg"
	echo "github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type Delivery struct {
	PersonalRatingUC PersonalRatingUsecase.UsecaseI
}

func (del *Delivery) ownerOrAdminValidate(c echo.Context, PersonalRating *models.PersonalRating) error {
	user, ok := c.Get("user").(*models.User)

	if !ok {
		c.Logger().Error("can't get user from context")
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	if user.Role == 0 || PersonalRating.UserID == user.ID {
		return nil
	}

	return models.ErrPermissionDenied
}

func (delivery *Delivery) CreatePersonalRating(c echo.Context) error {

	var reqPersonalRating dto.ReqCreateUpdatePersonalRating
	err := c.Bind(&reqPersonalRating)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqPersonalRating); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error("can't parse context user_id")
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	PersonalRating := reqPersonalRating.ToModelPersonalRating()
	PersonalRating.UserID = userId
	err = delivery.PersonalRatingUC.CreatePersonalRating(PersonalRating)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respPersonalRating := dto.GetResponseFromModelPersonalRating(PersonalRating)

	return c.JSON(http.StatusOK, pkg.Response{Body: *respPersonalRating})
}

func (delivery *Delivery) GetPersonalRating(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}
	PersonalRating, err := delivery.PersonalRatingUC.GetPersonalRating(id)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	err = delivery.ownerOrAdminValidate(c, PersonalRating)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respPersonalRating := dto.GetResponseFromModelPersonalRating(PersonalRating)
	return c.JSON(http.StatusOK, pkg.Response{Body: *respPersonalRating})
}

func (delivery *Delivery) UpdatePersonalRating(c echo.Context) error {

	var reqPersonalRating dto.ReqCreateUpdatePersonalRating
	err := c.Bind(&reqPersonalRating)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqPersonalRating); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error("can't parse context user_id")
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	PersonalRating := reqPersonalRating.ToModelPersonalRating()
	PersonalRating.UserID = userId
	err = delivery.PersonalRatingUC.UpdatePersonalRating(PersonalRating)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respPersonalRating := dto.GetResponseFromModelPersonalRating(PersonalRating)

	return c.JSON(http.StatusOK, pkg.Response{Body: *respPersonalRating})
}

func (delivery *Delivery) DeletePersonalRating(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
	}

	userId, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error("can't parse context user_id")
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	err = delivery.PersonalRatingUC.DeletePersonalRating(id, userId)

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

func NewDelivery(e *echo.Echo, eu PersonalRatingUsecase.UsecaseI, aclM *middleware.AclMiddleware) {
	handler := &Delivery{
		PersonalRatingUC: eu,
	}

	e.POST("/personal_rating/create", handler.CreatePersonalRating)
	e.POST("/personal_rating/edit", handler.UpdatePersonalRating)  // acl: owner
	e.GET("/personal_rating/:id", handler.GetPersonalRating)       // acl: owner, admin
	e.DELETE("/personal_rating/:id", handler.DeletePersonalRating) // acl: owner
}
