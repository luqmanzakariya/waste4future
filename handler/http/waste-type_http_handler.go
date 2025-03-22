package http

import (
	"net/http"
	"recyclehub-service/model/web/response"
	"recyclehub-service/usecase"

	"github.com/labstack/echo/v4"
)

type WasteTypeHandler struct {
	useCase usecase.IWasteTypeUsecase
}

func NewWasteTypeHandler(useCase usecase.IWasteTypeUsecase) *WasteTypeHandler {
	return &WasteTypeHandler{useCase: useCase}
}

// GetAllWasteTypes handles the HTTP GET request to retrieve all waste types.
func (h *WasteTypeHandler) GetAllWasteTypes(c echo.Context) error {
	wasteTypes, err := h.useCase.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wasteTypes)
}

// GetWasteTypeByID handles the HTTP GET request to retrieve a specific waste type by ID.
func (h *WasteTypeHandler) GetWasteTypeByID(c echo.Context) error {
	id := c.Param("id")
	wasteType, err := h.useCase.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{Message: "Waste type not found"})
	}
	return c.JSON(http.StatusOK, wasteType)
}

// RegisterRoutes registers the waste type routes with the Echo instance.
func (h *WasteTypeHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/waste-types", h.GetAllWasteTypes)
	e.GET("/waste-types/:id", h.GetWasteTypeByID)
}
