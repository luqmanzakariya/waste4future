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

// RecycleHub Create
// @Summary Create a Recycle Hub
// @Description Add and register new Recycle Hub
// @Tags RecycleHubs
// @Accept json
// @Produce json
// @Param request body model.PayloadCreateRecycleHub true "create recycle hub payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseRecycleHub} "Recycle Hub created"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Router /api/recycle-hubs [post]
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

// RecycleHub FindAll
// @Summary FindAll Recycle Hubs
// @Description Retrieve all Recycle Hubs
// @Tags RecycleHubs
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.RecycleHub} "Recycle Hubs list"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Router /api/recycle-hubs [get]
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

// RecycleHub FindById
// @Summary Find Recycle Hub by ID
// @Description Retrieve a Recycle Hub by its ID
// @Tags RecycleHubs
// @Produce json
// @Param id path string true "recycle hub id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=model.ResponseRecycleHub} "Recycle Hub found"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Router /api/recycle-hubs/{id} [get]
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

// RecycleHub Update
// @Summary Update a Recycle Hub
// @Description Update Recycle Hub by ID
// @Tags RecycleHubs
// @Accept json
// @Produce json
// @Param id path string true "recycle hub id" example("67cdcb62a50a990a870d928f")
// @Param request body model.PayloadUpdateRecycleHub true "update recycle hub payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseRecycleHub} "Recycle Hub updated"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /api/recycle-hubs/{id} [put]
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

// RecycleHub Delete
// @Summary Delete a Recycle Hub
// @Description Delete Recycle Hub by ID
// @Tags RecycleHubs
// @Produce json
// @Param id path string true "recycle hub id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=string} "Recycle Hub deleted"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /api/recycle-hubs/{id} [delete]
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
