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

// CreateFilm godoc
// @Summary      Create Film
// @Description  Create Film
// @Tags     	 Film
// @Accept	 application/json
// @Produce  application/json
// @Param    Film body dto.ReqCreateUpdateFilm true "Film info"
// @Success  200 {object} pkg.Response{body=dto.RespFilm} "success update Film"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 422 {object} echo.HTTPError "unprocessable entity"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 400 {object} echo.HTTPError "bad req"
// @Failure 403 {object} echo.HTTPError "invalid csrf or permission denied"
// @Router   /film/create [post]
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

// GetFilm godoc
// @Summary      Show a post
// @Description  Get Film by id. Acl: admin, owner
// @Tags     	 Film
// @Accept	 application/json
// @Produce  application/json
// @Param id  path int  true  "Film ID"
// @Success  200 {object} pkg.Response{body=dto.RespFilm} "success get Film"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /film/{id} [get]
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

// UpdateFilm godoc
// @Summary      Update a Film
// @Description  Update a Film. Acl: owner only
// @Tags     	 Film
// @Accept	 application/json
// @Produce  application/json
// @Param    Film body dto.ReqCreateUpdateFilm true "Film info"
// @Success  200 {object} pkg.Response{body=dto.RespFilm} "success update Film"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 422 {object} echo.HTTPError "unprocessable entity"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf or permission denied"
// @Router   /film/edit [post]
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

// DeleteFilm godoc
// @Summary      Delete a Film. Acl: owner only
// @Description  Delete a Film
// @Tags     	 Film
// @Accept	 application/json
// @Param id path int  true  "Film ID"
// @Success  204
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 404 {object} echo.HTTPError "can't find Film with such id"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /film/{id} [delete]
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

// GetFilmByPerson godoc
// @Summary      Get user lists
// @Description  Get user lists.
// @lists     tag
// @Produce  application/json
// @Param        day    query     string  false  "day for events"
// @Success  200 {object} pkg.Response{body=[]dto.RespTag} "success get lists"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /person/{person_id}/films [get]
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

// GetFilmByCountry godoc
// @Summary      Get user lists
// @Description  Get user lists.
// @lists     tag
// @Produce  application/json
// @Param        day    query     string  false  "day for events"
// @Success  200 {object} pkg.Response{body=[]dto.RespT} "success get lists"
// @Failure 405 {object} echo.HTTPError "Method Not Allowed"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /country/{country_id}/films [get]
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
	e.GET("/country/:country_id/films", handler.GetFilmByCountry)
	e.GET("/person/:person_id/films", handler.GetFilmByPerson)
}
