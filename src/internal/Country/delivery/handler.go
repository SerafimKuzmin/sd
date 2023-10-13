package delivery

import (
	CountryUsecase "github.com/SerafimKuzmin/sd/src/internal/Country/usecase"
	"github.com/SerafimKuzmin/sd/src/internal/middleware"
	"github.com/SerafimKuzmin/sd/src/models"
	"github.com/SerafimKuzmin/sd/src/models/dto"
	"github.com/SerafimKuzmin/sd/src/pkg"
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

// CreateCountry godoc
// @Summary      Create Country
// @Description  Create Country
// @Tags     	 Country
// @Accept	 application/json
// @Produce  application/json
// @Param    Country body dto.ReqCreateUpdateCountry true "Country info"
// @Success  200 {object} pkg.Response{body=dto.RespCountry} "success update Country"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 422 {object} echo.HTTPError "unprocessable entity"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 400 {object} echo.HTTPError "bad req"
// @Failure 403 {object} echo.HTTPError "invalid csrf or permission denied"
// @Router   /country/create [post]
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

// GetCountry godoc
// @Summary      Show a post
// @Description  Get Country by id. Acl: admin, owner
// @Tags     	 Country
// @Accept	 application/json
// @Produce  application/json
// @Param id  path int  true  "Country ID"
// @Success  200 {object} pkg.Response{body=dto.RespCountry} "success get Country"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /country/{id} [get]
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

// UpdateCountry godoc
// @Summary      Update a Country
// @Description  Update a Country. Acl: owner
// @Tags     	 Country
// @Accept	 application/json
// @Produce  application/json
// @Param    Country body dto.ReqCreateUpdateCountry true "Country info"
// @Success  200 {object} pkg.Response{body=dto.RespCountry} "success update Country"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 422 {object} echo.HTTPError "unprocessable entity"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf or permission denied"
// @Router   /country/edit [post]
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

// DeleteCountry godoc
// @Summary      Delete an Country
// @Description  Delete an Country. Acl: owner
// @Tags     	 Country
// @Accept	 application/json
// @Param id path int  true  "Country ID"
// @Success  204
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 404 {object} echo.HTTPError "can't find Country with such id"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /country/{id} [delete]
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
