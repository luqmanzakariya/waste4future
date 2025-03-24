package handler

import (
	"net/http"
	"reyclehub-service/model"
	"reyclehub-service/usecase"

	"github.com/labstack/echo/v4"
)

type wasteTypeHandler struct {
	WasteTypeUsecase usecase.IWasteTypeUsecase
}

func NewWasteTypeHandler(wasteTypeUsecase usecase.IWasteTypeUsecase) wasteTypeHandler {
	return wasteTypeHandler{
		WasteTypeUsecase: wasteTypeUsecase,
	}
}

func (w wasteTypeHandler) InitRoutes(g *echo.Group) {
	g.GET("", w.FindAll)
	g.GET("/:id", w.FindByID)
}

func (w wasteTypeHandler) FindAll(c echo.Context) error {
	wasteTypes, err := w.WasteTypeUsecase.FindAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   wasteTypes,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (w wasteTypeHandler) FindByID(c echo.Context) error {
	id := c.Param("id")
	data, err := w.WasteTypeUsecase.FindByID(c.Request().Context(), id)
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
