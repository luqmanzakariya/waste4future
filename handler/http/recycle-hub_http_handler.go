package http

import (
	"net/http"
	"recyclehub-service/model/web/request"
	"recyclehub-service/model/web/response"
	"recyclehub-service/usecase"

	"github.com/labstack/echo/v4"
)

type RecycleHubHandler struct {
	useCase usecase.IRecycleHubUsecase
}

func NewRecycleHubHandler(useCase usecase.IRecycleHubUsecase) *RecycleHubHandler {
	return &RecycleHubHandler{useCase: useCase}
}

// CreateRecycleHub handles the HTTP POST request to create a new recycle hub.
func (h *RecycleHubHandler) CreateRecycleHub(c echo.Context) error {
	var req request.RecycleHubCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid request payload"})
	}

	recycleHub, err := h.useCase.Create(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, recycleHub)
}

// GetAllRecycleHubs handles the HTTP GET request to retrieve all recycle hubs.
func (h *RecycleHubHandler) GetAllRecycleHubs(c echo.Context) error {
	recycleHubs, err := h.useCase.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, recycleHubs)
}

// GetRecycleHubByID handles the HTTP GET request to retrieve a specific recycle hub by ID.
func (h *RecycleHubHandler) GetRecycleHubByID(c echo.Context) error {
	id := c.Param("id")
	recycleHub, err := h.useCase.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{Message: "Recycle hub not found"})
	}
	return c.JSON(http.StatusOK, recycleHub)
}

// UpdateRecycleHub handles the HTTP PUT request to update an existing recycle hub.
func (h *RecycleHubHandler) UpdateRecycleHub(c echo.Context) error {
	id := c.Param("id")
	var req request.RecycleHubUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid request payload"})
	}

	recycleHub, err := h.useCase.Update(c.Request().Context(), req, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, recycleHub)
}

// DeleteRecycleHub handles the HTTP DELETE request to delete a recycle hub by ID.
func (h *RecycleHubHandler) DeleteRecycleHub(c echo.Context) error {
	id := c.Param("id")
	if err := h.useCase.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{Message: "Recycle hub deleted successfully"})
}

// RegisterRoutes registers the recycle hub routes with the Echo instance.
func (h *RecycleHubHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/recycle-hubs", h.CreateRecycleHub)
	e.GET("/recycle-hubs", h.GetAllRecycleHubs)
	e.GET("/recycle-hubs/:id", h.GetRecycleHubByID)
	e.PUT("/recycle-hubs/:id", h.UpdateRecycleHub)
	e.DELETE("/recycle-hubs/:id", h.DeleteRecycleHub)
}
