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

type orderHandler struct {
	OrderUsecase   usecase.IOrderUsecase
	UserGrpcClient userPB.UserServiceClient
}

func NewOrderHandler(orderUsecase usecase.IOrderUsecase, userServiceClient userPB.UserServiceClient) orderHandler {
	return orderHandler{
		OrderUsecase:   orderUsecase,
		UserGrpcClient: userServiceClient,
	}
}

func (o orderHandler) InitRoutes(g *echo.Group) {
	g.GET("", o.FindAll, middleware.AuthMiddleware)
	g.POST("/checkout", o.CheckoutOrder, middleware.AuthMiddleware)
	g.GET("/scheduler-shipping", o.SchedulerShipping)
}

// Order FindAll
// @Summary FindAll My Orders
// @Description Find All Orders by user token
// @Tags Orders
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.Order} "Order list"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /api/orders [get]
func (o orderHandler) FindAll(c echo.Context) error {
	token := c.Get("token").(string)
	// Create a context with the Authorization header for gRPC
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)
	userInfo, err := o.UserGrpcClient.Validate(ctx, &userPB.ValidateRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("User validation failed: %v", err))
	}
	userId := userInfo.ID

	orders, err := o.OrderUsecase.FindAllByUserID(ctx, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   orders,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// Order Checkout Order
// @Summary Checkout Order
// @Description Checkout Order and Change Status
// @Tags Orders
// @Produce json
// @Success 200 {object} model.WebResponse{data=model.ResponseOrder} "Order created"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /api/orders/checkout [post]
func (o orderHandler) CheckoutOrder(c echo.Context) error {
	token := c.Get("token").(string)
	// Create a context with the Authorization header for gRPC
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)
	userInfo, err := o.UserGrpcClient.Validate(ctx, &userPB.ValidateRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("User validation failed: %v", err))
	}
	userId := userInfo.ID

	err = o.OrderUsecase.CheckoutOrder(ctx, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Error save order id: %v", err))
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	return c.JSON(http.StatusOK, webResponse)
}

// Order SchedulerShipping
// @Summary SchedulerShipping Orders
// @Description SchedulerShipping Update Shipping And Orders
// @Tags Orders
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.Order} "Order list"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Router /api/orders/scheduler-shipping [get]
func (o orderHandler) SchedulerShipping(c echo.Context) error {
	err := o.OrderUsecase.SchedulerUpdateOrderAndShipping(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	return c.JSON(http.StatusOK, webResponse)
}
