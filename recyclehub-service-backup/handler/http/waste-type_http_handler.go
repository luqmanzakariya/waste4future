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

// GetAllWasteTypes godoc
// @Summary Get all waste types
// @Description Retrieve a list of all waste types
// @Tags waste-types
// @Accept json
// @Produce json
// @Success 200 {array} response.WasteTypeResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /waste-types [get]
func (h *WasteTypeHandler) GetAllWasteTypes(c echo.Context) error {
	wasteTypes, err := h.useCase.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wasteTypes)
}

// GetWasteTypeByID handles the HTTP GET request to retrieve a specific waste type by ID.

// GetWasteTypeByID godoc
// @Summary Get a waste type by ID
// @Description Retrieve a specific waste type by its ID
// @Tags waste-types
// @Accept json
// @Produce json
// @Param id path string true "Waste Type ID"
// @Success 200 {object} response.WasteTypeResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /waste-types/{id} [get]
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
