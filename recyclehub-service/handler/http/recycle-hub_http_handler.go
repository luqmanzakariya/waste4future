package handler

import (
	"fmt"
	"net/http"
	"reyclehub-service/middleware"
	"reyclehub-service/model"
	"reyclehub-service/usecase"

	"github.com/labstack/echo/v4"
)

type recycleHubHandler struct {
	RecycleHubUsecase usecase.IRecycleHubUsecase
}

func NewRecycleHubHandler(recycleHubUsecase usecase.IRecycleHubUsecase) recycleHubHandler {
	return recycleHubHandler{
		RecycleHubUsecase: recycleHubUsecase,
	}
}

func (r recycleHubHandler) InitRoutes(g *echo.Group) {
	g.POST("", r.Create)
	g.GET("", r.FindAll)
	g.GET("/:id", r.FindByID)
	g.PUT("/:id", r.Update, middleware.AuthMiddleware)
	g.DELETE("/:id", r.Delete, middleware.AuthMiddleware)
}

func (r recycleHubHandler) Create(c echo.Context) error {
	var payload model.PayloadCreateRecycleHub
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	created, err := r.RecycleHubUsecase.Create(c.Request().Context(), payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   created,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (r recycleHubHandler) FindAll(c echo.Context) error {
	recycleHubs, err := r.RecycleHubUsecase.FindAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   recycleHubs,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (r recycleHubHandler) FindByID(c echo.Context) error {
	id := c.Param("id")
	data, err := r.RecycleHubUsecase.FindByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (r recycleHubHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var payload model.PayloadUpdateRecycleHub
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedData, err := r.RecycleHubUsecase.Update(c.Request().Context(), id, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   updatedData,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (r recycleHubHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := r.RecycleHubUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := fmt.Sprintf("success deleted recycle hub with id: %s", id)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}

	return c.JSON(http.StatusOK, webResponse)
}
