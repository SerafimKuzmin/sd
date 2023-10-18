package delivery

import (
	ListUsecase "github.com/SerafimKuzmin/sd/backend/internal/List/usecase"
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
	ListUC ListUsecase.UsecaseI
}

func (del *Delivery) ownerOrAdminValidate(c echo.Context, List *models.List) error {
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

func (delivery *Delivery) CreateList(c echo.Context) error {

	var reqList dto.ReqCreateUpdateList
	err := c.Bind(&reqList)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqList); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
	}

	List := reqList.ToModelList()
	err = delivery.ListUC.CreateList(List)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respList := dto.GetResponseFromModelList(List)

	return c.JSON(http.StatusOK, pkg.Response{Body: *respList})
}

func (delivery *Delivery) GetList(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}
	List, err := delivery.ListUC.GetList(id)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	err = delivery.ownerOrAdminValidate(c, List)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respList := dto.GetResponseFromModelList(List)
	return c.JSON(http.StatusOK, pkg.Response{Body: *respList})
}

func (delivery *Delivery) UpdateList(c echo.Context) error {

	var reqList dto.ReqCreateUpdateList
	err := c.Bind(&reqList)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqList); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	List := reqList.ToModelList()
	err = delivery.ListUC.UpdateList(List)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	respList := dto.GetResponseFromModelList(List)

	return c.JSON(http.StatusOK, pkg.Response{Body: *respList})
}

func (delivery *Delivery) AddFilm(c echo.Context) error {

	var reqList dto.ReqAddFilm
	err := c.Bind(&reqList)

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := pkg.IsRequestValid(&reqList); !ok {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	List := reqList.ToModelList()
	err = delivery.ListUC.AddFilm(List.ID, List.FilmID)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	return c.JSON(http.StatusOK, pkg.Response{Body: nil})
}

func (delivery *Delivery) DeleteList(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, models.ErrNotFound.Error())
	}

	err = delivery.ListUC.DeleteList(id)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (delivery *Delivery) GetUserLists(c echo.Context) error {
	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	lists, err := delivery.ListUC.GetUserLists(userId)

	if err != nil {
		c.Logger().Error(err)
		return handleError(err)
	}

	resplists := dto.GetResponseFromModelLists(lists)

	return c.JSON(http.StatusOK, pkg.Response{Body: resplists})
}

func (delivery *Delivery) GetFilmsByList(c echo.Context) error {
	listID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
	}

	films, err := delivery.ListUC.GetFilmsByList(listID)

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

func NewDelivery(e *echo.Echo, eu ListUsecase.UsecaseI, aclM *middleware.AclMiddleware) {
	handler := &Delivery{
		ListUC: eu,
	}

	e.POST("/list/create", handler.CreateList)
	e.POST("/list/edit", handler.UpdateList) // acl: owner
	e.POST("/list/add", handler.AddFilm)
	e.GET("/list/:id", handler.GetList)
	e.GET("/user/:id/lists", handler.GetUserLists)
	e.GET("/list/:id/films", handler.GetFilmsByList)
	e.DELETE("/list/:id", handler.DeleteList) // acl: owner
}
