package delivery

import (
	PersonUsecase "github.com/SerafimKuzmin/sd/backend/internal/Person/usecase"
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
	PersonUC PersonUsecase.UsecaseI
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

// CreatePerson godoc
// @Summary      Create Person
// @Description  Create Person
// @Tags     	 Person
// @Accept	 application/json
// @Produce  application/json
// @Param    Person body dto.ReqCreateUpdatePerson true "Person info"
// @Success  200 {object} pkg.Response{body=dto.RespPerson} "success update Person"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 422 {object} echo.HTTPError "unprocessable entity"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 400 {object} echo.HTTPError "bad req"
// @Failure 403 {object} echo.HTTPError "invalid csrf or permission denied"
// @Router   /person/create [post]
func (delivery *Delivery) CreatePerson(c echo.Context) error {

	var reqPerson dto.ReqCreateUpdatePerson
	err := c.Bind(&reqPerson)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqPerson); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	Person := reqPerson.ToModelPerson()

	err = delivery.adminValidate(c)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	err = delivery.PersonUC.CreatePerson(Person)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respPerson := dto.GetResponseFromModelPerson(Person)

	return c.JSON(http.StatusOK, pkg.Response{Body: *respPerson})
}

// GetPerson godoc
// @Summary      Show a post
// @Description  Get Person by id. Acl: admin, owner
// @Tags     	 Person
// @Accept	 application/json
// @Produce  application/json
// @Param id  path int  true  "Person ID"
// @Success  200 {object} pkg.Response{body=dto.RespPerson} "success get Person"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /person/{id} [get]
func (delivery *Delivery) GetPerson(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}
	Person, err := delivery.PersonUC.GetPerson(id)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	err = delivery.adminValidate(c)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respPerson := dto.GetResponseFromModelPerson(Person)
	return c.JSON(http.StatusOK, pkg.Response{Body: *respPerson})
}

// UpdatePerson godoc
// @Summary      Update a Person
// @Description  Update a Person. Acl: owner
// @Tags     	 Person
// @Accept	 application/json
// @Produce  application/json
// @Param    Person body dto.ReqCreateUpdatePerson true "Person info"
// @Success  200 {object} pkg.Response{body=dto.RespPerson} "success update Person"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 422 {object} echo.HTTPError "unprocessable entity"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf or permission denied"
// @Router   /person/edit [post]
func (delivery *Delivery) UpdatePerson(c echo.Context) error {

	var reqPerson dto.ReqCreateUpdatePerson
	err := c.Bind(&reqPerson)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqPerson); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	Person := reqPerson.ToModelPerson()

	err = delivery.adminValidate(c)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	err = delivery.PersonUC.UpdatePerson(Person)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respPerson := dto.GetResponseFromModelPerson(Person)

	return c.JSON(http.StatusOK, pkg.Response{Body: *respPerson})
}

// DeletePerson godoc
// @Summary      Delete an Person
// @Description  Delete an Person. Acl: owner
// @Tags     	 Person
// @Accept	 application/json
// @Param id path int  true  "Person ID"
// @Success  204
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 404 {object} echo.HTTPError "can't find Person with such id"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /person/{id} [delete]
func (delivery *Delivery) DeletePerson(c echo.Context) error {

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

	err = delivery.PersonUC.DeletePerson(id)

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

func NewDelivery(e *echo.Echo, pu PersonUsecase.UsecaseI, aclM *middleware.AclMiddleware) {
	handler := &Delivery{
		PersonUC: pu,
	}

	e.POST("/person/create", handler.CreatePerson)
	e.POST("/person/edit", handler.UpdatePerson)  //acl: owner
	e.GET("/person/:id", handler.GetPerson)       //acl: owner, admin
	e.DELETE("/person/:id", handler.DeletePerson) //acl: owner
}
