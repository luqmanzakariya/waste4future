package handler

import (
	"net/http"
	"operation-service/middleware"
	"operation-service/model"
	pbUser "operation-service/pb/user"
	userPB "operation-service/pb/user"
	"operation-service/usecase"
	"operation-service/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
)

type TransactionHandler struct {
	transactionUsecase usecase.TransactionUsecase
	validator          *validator.Validate
	UserGrpcClient     userPB.UserServiceClient
}

func NewTransactionHandler(tu usecase.TransactionUsecase, v *validator.Validate, userServiceClient userPB.UserServiceClient) *TransactionHandler {
	return &TransactionHandler{
		transactionUsecase: tu,
		validator:          v,
		UserGrpcClient:     userServiceClient,
	}
}

func (h *TransactionHandler) InitRoutes(e *echo.Group) {
	e.POST("", h.CreateTransaction, middleware.AuthMiddleware)
	e.GET("/:id", h.GetTransaction, middleware.AuthMiddleware)
	e.GET("/orders/:orderId/transactions", h.GetTransactionsByOrder, middleware.AuthMiddleware)
	e.PUT("/:id/status", h.UpdateTransactionStatus, middleware.AuthMiddleware)
	e.DELETE("/:id", h.DeleteTransaction, middleware.AuthMiddleware)
	e.GET("", h.GetTransactionsByDateRange, middleware.AuthMiddleware)
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Create a new payment transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body model.CreateTransactionRequest true "Transaction data"
// @Success 201 {object} model.Transaction
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Security BearerAuth
// @Router /transactions [post]
func (h *TransactionHandler) CreateTransaction(c echo.Context) error {

	token := c.Get("token").(string)
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)

	// Validate user and get user ID
	userInfo, err := h.UserGrpcClient.Validate(ctx, &pbUser.ValidateRequest{})
	if err != nil {
		utils.LogEntry(c).Error("User validation failed: ", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "User validation failed")
	}

	var req model.CreateTransactionRequest
	if err := c.Bind(&req); err != nil {
		utils.LogEntry(c).Error("Invalid request payload: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Validate request
	if err := h.validator.Struct(req); err != nil {
		utils.LogEntry(c).Error("Validation failed: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, utils.ValidationError(err))
	}

	createdTransaction, err := h.transactionUsecase.CreateTransaction(ctx, &req, userInfo.ID)
	if err != nil {
		utils.LogEntry(c).Error("Failed to create transaction: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create transaction")
	}

	return c.JSON(http.StatusCreated, createdTransaction)
}

// GetTransaction godoc
// @Summary Get a transaction
// @Description Get a transaction by ID
// @Tags transactions
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} model.Transaction
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Security BearerAuth
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetTransaction(c echo.Context) error {
	token := c.Get("token").(string)
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)

	if _, err := h.UserGrpcClient.Validate(ctx, &pbUser.ValidateRequest{}); err != nil {
		utils.LogEntry(c).Error("User validation failed: ", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "User validation failed")
	}

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.LogEntry(c).Error("Invalid transaction ID format: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid transaction ID format")
	}

	transaction, err := h.transactionUsecase.GetTransaction(ctx, objID)
	if err != nil {
		utils.LogEntry(c).Error("Failed to get transaction: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get transaction")
	}

	if transaction == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Transaction not found")
	}

	return c.JSON(http.StatusOK, transaction)
}

// GetTransactionsByOrder godoc
// @Summary Get transactions by order
// @Description Get all transactions for a specific order
// @Tags transactions
// @Produce json
// @Param orderId path string true "Order ID"
// @Success 200 {array} model.Transaction
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {object} utils.HTTPError
// @Security BearerAuth
// @Router /orders/{orderId}/transactions [get]
func (h *TransactionHandler) GetTransactionsByOrder(c echo.Context) error {
	token := c.Get("token").(string)
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)

	if _, err := h.UserGrpcClient.Validate(ctx, &pbUser.ValidateRequest{}); err != nil {
		utils.LogEntry(c).Error("User validation failed: ", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "User validation failed")
	}

	orderID := c.Param("orderId")

	transactions, err := h.transactionUsecase.GetTransactionsByOrder(ctx, orderID)
	if err != nil {
		utils.LogEntry(c).Error("Failed to get transactions by order: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get transactions by order")
	}

	return c.JSON(http.StatusOK, transactions)
}

// UpdateTransactionStatus godoc
// @Summary Update transaction status
// @Description Update payment status of a transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Param status body model.UpdateTransactionStatusRequest true "Status update"
// @Success 200 {object} model.Transaction
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {object} utils.HTTPError
// @Failure 403 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Security BearerAuth
// @Router /transactions/{id}/status [put]
func (h *TransactionHandler) UpdateTransactionStatus(c echo.Context) error {
	token := c.Get("token").(string)
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)

	userResp, err := h.UserGrpcClient.Validate(ctx, &pbUser.ValidateRequest{})
	if err != nil {
		utils.LogEntry(c).Error("User validation failed: ", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "User validation failed")
	}

	if userResp.ID == 0 {
		return echo.NewHTTPError(http.StatusForbidden, "User not authorized")
	}
	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.LogEntry(c).Error("Invalid transaction ID format: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid transaction ID format")
	}

	var req model.UpdateTransactionStatusRequest
	if err := c.Bind(&req); err != nil {
		utils.LogEntry(c).Error("Failed to bind request: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := h.transactionUsecase.UpdateTransactionStatus(ctx, objID, req.Status); err != nil {
		utils.LogEntry(c).Error("Failed to update status: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update status")
	}

	// Return the updated transaction
	transaction, err := h.transactionUsecase.GetTransaction(ctx, objID)
	if err != nil {
		utils.LogEntry(c).Error("Failed to get transaction: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get transaction")
	}

	return c.JSON(http.StatusOK, transaction)
}

// DeleteTransaction godoc
// @Summary Delete a transaction
// @Description Delete a transaction by ID
// @Tags transactions
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 204
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {object} utils.HTTPError
// @Failure 403 {object} utils.HTTPError
// @Security BearerAuth
// @Router /transactions/{id} [delete]
func (h *TransactionHandler) DeleteTransaction(c echo.Context) error {
	token := c.Get("token").(string)
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)

	userResp, err := h.UserGrpcClient.Validate(ctx, &pbUser.ValidateRequest{})
	if err != nil {
		utils.LogEntry(c).Error("User validation failed: ", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "User validation failed")
	}

	if userResp.ID == 0 {
		return echo.NewHTTPError(http.StatusForbidden, "User not authorized")
	}

	id := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.LogEntry(c).Error("Invalid transaction ID format: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid transaction ID format")
	}

	if err := h.transactionUsecase.DeleteTransaction(ctx, objID); err != nil {
		utils.LogEntry(c).Error("Failed to delete transaction: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete transaction")
	}

	return c.NoContent(http.StatusNoContent)
}

// GetTransactionsByDateRange godoc
// @Summary Get transactions by date range
// @Description Get transactions within a specific date range
// @Tags transactions
// @Produce json
// @Param start query string true "Start date (RFC3339 format)"
// @Param end query string true "End date (RFC3339 format)"
// @Success 200 {array} model.Transaction
// @Failure 400 {object} utils.HTTPError
// @Failure 401 {object} utils.HTTPError
// @Failure 403 {object} utils.HTTPError
// @Security BearerAuth
// @Router /transactions [get]
func (h *TransactionHandler) GetTransactionsByDateRange(c echo.Context) error {
	token := c.Get("token").(string)
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)

	userResp, err := h.UserGrpcClient.Validate(ctx, &pbUser.ValidateRequest{})
	if err != nil {
		utils.LogEntry(c).Error("User validation failed: ", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "User validation failed")
	}

	if userResp.ID == 0 {
		return echo.NewHTTPError(http.StatusForbidden, "User not authorized")
	}

	start, err := time.Parse(time.RFC3339, c.QueryParam("start"))
	if err != nil {
		utils.LogEntry(c).Error("Invalid start date format: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid start date format, use RFC3339 format")
	}

	end, err := time.Parse(time.RFC3339, c.QueryParam("end"))
	if err != nil {
		utils.LogEntry(c).Error("Invalid end date format: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid end date format, use RFC3339 format")
	}

	transactions, err := h.transactionUsecase.GetTransactionsByDateRange(ctx, start, end)
	if err != nil {
		utils.LogEntry(c).Error("Failed to get transactions by date range: ", err)
		return err
	}

	return c.JSON(http.StatusOK, transactions)
}
