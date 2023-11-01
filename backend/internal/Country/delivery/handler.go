package delivery

import (
	CountryUsecase "github.com/SerafimKuzmin/sd/backend/internal/Country/usecase"
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
	CountryUC CountryUsecase.UsecaseI
}

func (del *Delivery) adminValidate(c echo.Context) error {
	user, ok := c.Get("user").(*models.User)

	if !ok {
		c.Logger().Error("can't get user from context")
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	if user.Role == 2 {
		return nil
	}

	return models.ErrPermissionDenied
}

func (delivery *Delivery) CreateCountry(c echo.Context) error {

	var reqCountry dto.ReqCreateUpdateCountry
	err := c.Bind(&reqCountry)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqCountry); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	Country := reqCountry.ToModelCountry()

	err = delivery.adminValidate(c)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	err = delivery.CountryUC.CreateCountry(Country)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respCountry := dto.GetResponseFromModelCountry(Country)

	return c.JSON(http.StatusOK, pkg.Response{Body: *respCountry})
}

func (delivery *Delivery) GetCountry(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}
	Country, err := delivery.CountryUC.GetCountry(id)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	err = delivery.adminValidate(c)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respCountry := dto.GetResponseFromModelCountry(Country)
	return c.JSON(http.StatusOK, pkg.Response{Body: *respCountry})
}

func (delivery *Delivery) UpdateCountry(c echo.Context) error {

	var reqCountry dto.ReqCreateUpdateCountry
	err := c.Bind(&reqCountry)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqCountry); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	Country := reqCountry.ToModelCountry()

	err = delivery.adminValidate(c)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	err = delivery.CountryUC.UpdateCountry(Country)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respCountry := dto.GetResponseFromModelCountry(Country)

	return c.JSON(http.StatusOK, pkg.Response{Body: *respCountry})
}

func (delivery *Delivery) DeleteCountry(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
	}

	err = delivery.adminValidate(c)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	err = delivery.CountryUC.DeleteCountry(id)

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

func NewDelivery(e *echo.Echo, pu CountryUsecase.UsecaseI, aclM *middleware.AclMiddleware) {
	handler := &Delivery{
		CountryUC: pu,
	}

	e.POST("/country/create", handler.CreateCountry)
	e.POST("/country/edit", handler.UpdateCountry)  //acl: owner
	e.GET("/country/:id", handler.GetCountry)       //acl: owner, admin
	e.DELETE("/country/:id", handler.DeleteCountry) //acl: owner
}
