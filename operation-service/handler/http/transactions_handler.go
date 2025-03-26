package handler

import (
	"fmt"
	"net/http"
	"operation-service/middleware"
	"operation-service/model"
	"operation-service/usecase"

	userPB "operation-service/pb/user"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
)

type transactionHandler struct {
	TransactionUsecase usecase.ITransactionUsecase
	UserGrpcClient     userPB.UserServiceClient
}

func NewTransactionHandler(transactionUsecase usecase.ITransactionUsecase, userServiceClient userPB.UserServiceClient) transactionHandler {
	return transactionHandler{
		TransactionUsecase: transactionUsecase,
		UserGrpcClient:     userServiceClient,
	}
}

func (t transactionHandler) InitRoutes(g *echo.Group) {
	g.POST("", t.Create, middleware.AuthMiddleware)
	g.GET("", t.FindAll, middleware.AuthMiddleware)
	g.GET("/:id", t.FindByID, middleware.AuthMiddleware)
	g.PUT("/:id", t.Update, middleware.AuthMiddleware)
	g.DELETE("/:id", t.Delete, middleware.AuthMiddleware)
	g.PUT("/approved/:id", t.ApprovePayment, middleware.AuthMiddleware)
	g.PUT("/reject/:id", t.RejectPayment, middleware.AuthMiddleware)
}

// Transaction Create
// @Summary Create a Transaction
// @Description Add and register new Transaction
// @Tags Transactions
// @Produce json
// @Param request body model.PayloadCreateTransaction true "create transaction payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseTransaction} "Transaction created"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /api/transactions [post]
func (t transactionHandler) Create(c echo.Context) error {
	token := c.Get("token").(string)
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)
	_, err := t.UserGrpcClient.Validate(ctx, &userPB.ValidateRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("User validation failed: %v", err))
	}

	var payload model.PayloadCreateTransaction
	err = c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	created, err := t.TransactionUsecase.Create(ctx, payload)
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

// Transaction FindAll
// @Summary FindAll Transactions
// @Description Find All Transactions
// @Tags Transactions
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.Transaction} "Transaction list"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /api/transactions [get]
func (t transactionHandler) FindAll(c echo.Context) error {
	transactions, err := t.TransactionUsecase.FindAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   transactions,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// Transaction FindById
// @Summary FindById a Transaction
// @Description Find Transaction By ID
// @Tags Transactions
// @Produce json
// @Param id path string true "transaction id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=model.ResponseTransaction} "Transaction found"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /api/transactions/{id} [get]
func (t transactionHandler) FindByID(c echo.Context) error {
	id := c.Param("id")
	data, err := t.TransactionUsecase.FindByID(c.Request().Context(), id)
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

// Transaction UpdateById
// @Summary Update a Transaction
// @Description Update Transaction By ID
// @Tags Transactions
// @Produce json
// @Param id path string true "transaction id" example("67cdcb62a50a990a870d928f")
// @Param request body model.PayloadUpdateTransaction true "update transaction payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseTransaction} "Transaction updated"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /api/transactions/{id} [PUT]
func (t transactionHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var payload model.PayloadUpdateTransaction
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedData, err := t.TransactionUsecase.Update(c.Request().Context(), id, payload)
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

// Transaction DeleteById
// @Summary DeleteById a Transaction
// @Description Delete Transaction By ID
// @Tags Transactions
// @Produce json
// @Param id path string true "transaction id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=string} "Transaction deleted"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /api/transactions/{id} [DELETE]
func (t transactionHandler) Delete(c echo.Context) error {
	token := c.Get("token").(string)
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)
	_, err := t.UserGrpcClient.Validate(ctx, &userPB.ValidateRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("User validation failed: %v", err))
	}

	id := c.Param("id")
	err = t.TransactionUsecase.Delete(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   fmt.Sprintf("success deleted transaction with id: %s", id),
	}

	return c.JSON(http.StatusOK, webResponse)
}

// Transaction Approved Payment
// @Summary Update a Payment Status to Completed
// @Description Update Status Payment to Completed By ID
// @Tags Transactions
// @Produce json
// @Param id path string true "transaction id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=model.ResponseTransaction} "Transaction updated"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /api/transactions/approved/{id} [PUT]
func (t transactionHandler) ApprovePayment(c echo.Context) error {
	id := c.Param("id")
	var payload model.PayloadUpdateTransaction
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedData, err := t.TransactionUsecase.ApprovePayment(c.Request().Context(), id)
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

// Transaction Reject Payment
// @Summary Update a Payment Status to Reject
// @Description Update Status Payment to Reject By ID
// @Tags Transactions
// @Produce json
// @Param id path string true "transaction id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=model.ResponseTransaction} "Transaction updated"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /api/transactions/reject/{id} [PUT]
func (t transactionHandler) RejectPayment(c echo.Context) error {
	id := c.Param("id")
	var payload model.PayloadUpdateTransaction
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedData, err := t.TransactionUsecase.RejectPayment(c.Request().Context(), id)
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
