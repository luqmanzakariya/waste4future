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

// WasteType FindAll
// @Summary FindAll Waste Types
// @Description Retrieve all waste types
// @Tags WasteTypes
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.WasteType} "Waste types list"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Router /api/waste-types [get]
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

// WasteType FindById
// @Summary Find Waste Type by ID
// @Description Retrieve a waste type by its ID
// @Tags WasteTypes
// @Produce json
// @Param id path string true "waste type id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=model.WasteType} "Waste type found"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Router /api/waste-types/{id} [get]
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
