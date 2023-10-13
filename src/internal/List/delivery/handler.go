package delivery

import (
	ListUsecase "github.com/SerafimKuzmin/sd/src/internal/List/usecase"
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

// CreateList godoc
// @Summary      Create List
// @Description  Create List
// @Lists     	 List
// @Accept	 application/json
// @Produce  application/json
// @Param    List body dto.ReqCreateUpdateList true "List info"
// @Success  200 {object} pkg.Response{body=dto.RespList} "success update List"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 422 {object} echo.HTTPError "unprocessable entity"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 400 {object} echo.HTTPError "bad req"
// @Failure 403 {object} echo.HTTPError "invalid csrf or permission denied"
// @Router   /list/create [post]
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

// GetList godoc
// @Summary      Show a post
// @Description  Get List by id
// @Lists     	 List
// @Accept	 application/json
// @Produce  application/json
// @Param id  path int  true  "List ID"
// @Success  200 {object} pkg.Response{body=dto.RespList} "success get List"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Router   /list/{id} [get]
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

// UpdateList godoc
// @Summary      Update an List
// @Description  Update an List. Acl: owner
// @Lists     	 List
// @Accept	 application/json
// @Produce  application/json
// @Param    List body dto.ReqCreateUpdateList true "List info"
// @Success  200 {object} pkg.Response{body=dto.RespList} "success update List"
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 400 {object} echo.HTTPError "bad request"
// @Failure 422 {object} echo.HTTPError "unprocessable entity"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 403 {object} echo.HTTPError "invalid csrf or permission denied"
// @Router   /list/edit [post]
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

// DeleteList godoc
// @Summary      Delete an List
// @Description  Delete an List. Acl: owner
// @Lists     	 List
// @Accept	 application/json
// @Param id path int  true  "List ID"
// @Success  204
// @Failure 405 {object} echo.HTTPError "invalid http method"
// @Failure 500 {object} echo.HTTPError "internal server error"
// @Failure 401 {object} echo.HTTPError "no cookie"
// @Failure 404 {object} echo.HTTPError "can't find List with such id"
// @Failure 403 {object} echo.HTTPError "invalid csrf"
// @Router   /list/{id} [delete]
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

// GetUserLists godoc
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
// @Router   /user/{user_id}/lists [get]
func (delivery *Delivery) GetUserLists(c echo.Context) error {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)

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

// GetUserFilms godoc
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
// @Router   /list/{id}/films [get]
func (delivery *Delivery) GetUserFilms(c echo.Context) error {
	listID, err := strconv.ParseUint(c.Param("list_id"), 10, 64)

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
	e.GET("/list/:id", handler.GetList)
	e.GET("/user/:user_id/lists", handler.GetUserLists)
	e.GET("/list/:id/films", handler.GetUserLists)
	e.DELETE("/list/:id", handler.DeleteList) // acl: owner
}
