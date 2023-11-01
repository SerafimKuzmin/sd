package delivery

import (
	FilmUsecase "github.com/SerafimKuzmin/sd/backend/internal/Film/usecase"
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
	FilmUC FilmUsecase.UsecaseI
}

func (del *Delivery) ownerOrAdminValidate(c echo.Context) error {
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

func (delivery *Delivery) CreateFilm(c echo.Context) error {

	var reqFilm dto.ReqCreateUpdateFilm
	err := c.Bind(&reqFilm)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqFilm); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	err = delivery.ownerOrAdminValidate(c)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	Film := reqFilm.ToModelFilm()
	err = delivery.FilmUC.CreateFilm(Film)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respFilm := dto.GetResponseFromModelFilm(Film)

	return c.JSON(http.StatusOK, pkg.Response{Body: *respFilm})
}

func (delivery *Delivery) GetFilm(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}
	Film, err := delivery.FilmUC.GetFilm(id)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respFilm := dto.GetResponseFromModelFilm(Film)
	return c.JSON(http.StatusOK, pkg.Response{Body: *respFilm})
}

func (delivery *Delivery) UpdateFilm(c echo.Context) error {

	var reqFilm dto.ReqCreateUpdateFilm
	err := c.Bind(&reqFilm)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqFilm); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	_, ok := c.Get("user_id").(uint64)
	if !ok {
		c.Logger().Error("can't parse context user_id")
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	err = delivery.ownerOrAdminValidate(c)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	Film := reqFilm.ToModelFilm()
	err = delivery.FilmUC.UpdateFilm(Film)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respFilm := dto.GetResponseFromModelFilm(Film)

	return c.JSON(http.StatusOK, pkg.Response{Body: *respFilm})
}

func (delivery *Delivery) DeleteFilm(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
	}

	err = delivery.ownerOrAdminValidate(c)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	err = delivery.FilmUC.DeleteFilm(id)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (delivery *Delivery) GetFilmByPerson(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	films, err := delivery.FilmUC.GetFilmByPerson(id)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	resplists := dto.GetResponseFromModelFilms(films)

	return c.JSON(http.StatusOK, pkg.Response{Body: resplists})
}

func (delivery *Delivery) GetFilmByCountry(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	films, err := delivery.FilmUC.GetFilmByCountry(id)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	resplists := dto.GetResponseFromModelFilms(films)

	return c.JSON(http.StatusOK, pkg.Response{Body: resplists})
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

func NewDelivery(e *echo.Echo, eu FilmUsecase.UsecaseI, aclM *middleware.AclMiddleware) {
	handler := &Delivery{
		FilmUC: eu,
	}

	e.POST("/film/create", handler.CreateFilm)
	e.POST("/film/edit", handler.UpdateFilm)
	e.GET("/film/:id", handler.GetFilm)
	e.DELETE("/film/:id", handler.DeleteFilm)
	e.GET("/country/:id/films", handler.GetFilmByCountry)
	e.GET("/person/:person_id/films", handler.GetFilmByPerson)
}
