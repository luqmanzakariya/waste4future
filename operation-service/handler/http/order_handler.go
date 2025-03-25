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
	g.POST("", o.Create, middleware.AuthMiddleware)
	g.GET("", o.FindAll, middleware.AuthMiddleware)
	g.GET("/:id", o.FindByID, middleware.AuthMiddleware)
	g.PUT("/:id", o.Update, middleware.AuthMiddleware)
	g.DELETE("/:id", o.Delete, middleware.AuthMiddleware)
	g.POST("/save", o.SaveOrderID, middleware.AuthMiddleware)
	g.POST("/checkout", o.CheckoutOrder, middleware.AuthMiddleware)
	g.GET("/scheduler-shipping", o.SchedulerShipping)
}

// Order Create
// @Summary Create an Order
// @Description Add and register new Order
// @Tags Orders
// @Produce json
// @Param	request	body model.PayloadCreateOrder true "create order payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseOrder} "Order created"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /api/orders [post]
func (o orderHandler) Create(c echo.Context) error {
	token := c.Get("token").(string)
	// Create a context with the Authorization header for gRPC
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)
	userInfo, err := o.UserGrpcClient.Validate(ctx, &userPB.ValidateRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("User validation failed: %v", err))
	}
	userId := userInfo.ID

	var payloadCreateOrder model.PayloadCreateOrder
	err = c.Bind(&payloadCreateOrder)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	created, err := o.OrderUsecase.Create(ctx, payloadCreateOrder, userId)
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

// Order FindAll
// @Summary FindAll Orders
// @Description Find All Orders
// @Tags Orders
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.Order} "Order list"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Router /api/orders [get]
func (o orderHandler) FindAll(c echo.Context) error {
	orders, err := o.OrderUsecase.FindAll(c.Request().Context())
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

// Order FindById
// @Summary FindById an Order
// @Description Find Order By ID
// @Tags Orders
// @Produce json
// @Param id path string true "address id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=model.ResponseOrder} "Address found"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Router /api/orders/{id} [get]
func (o orderHandler) FindByID(c echo.Context) error {
	// # Get ID Params
	id := c.Param("id")

	data, err := o.OrderUsecase.FindByID(c.Request().Context(), id)
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

// Order UpdateById
// @Summary Update an Order
// @Description Update Order By ID
// @Tags Orders
// @Produce json
// @Param id path string true "order id" example("67cdcb62a50a990a870d928f")
// @Param	request	body model.PayloadUpdateOrder true "update order payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseOrder} "Order updated"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /api/orders/{id} [PUT]
func (o orderHandler) Update(c echo.Context) error {
	// # Get ID Params
	id := c.Param("id")

	var payload model.PayloadUpdateOrder
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedData, err := o.OrderUsecase.Update(c.Request().Context(), id, payload)
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

// Order DeleteById
// @Summary DeleteById an Order
// @Description Delete Order By ID
// @Tags Orders
// @Produce json
// @Param id path string true "order id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=string} "Order deleted"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /api/orders/{id} [DELETE]
func (o orderHandler) Delete(c echo.Context) error {
	token := c.Get("token").(string)
	fmt.Println("disini ada token", token)

	// Create a context with the Authorization header for gRPC
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)

	userInfo, err := o.UserGrpcClient.Validate(ctx, &userPB.ValidateRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("User validation failed: %v", err))
	}

	fmt.Println("----", userInfo.ID)
	fmt.Println("----", userInfo.Email)
	fmt.Println("----", userInfo.Name)
	fmt.Println("----", userInfo.AddressID)

	// # USER Info disini yang nantinya bisa dipakai

	// // # Get ID Params
	// id := c.Param("id")

	// err = o.OrderUsecase.Delete(ctx, id)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }

	// res := fmt.Sprintf("success deleted order with id: %s", id)

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		// Data:   res,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// Order Save Order Detail ID
// @Summary Save Order Detail ID
// @Description Save Order Detail ID
// @Tags Orders
// @Produce json
// @Success 200 {object} model.WebResponse{data=model.ResponseOrder} "Order created"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /api/orders/save [post]
func (o orderHandler) SaveOrderID(c echo.Context) error {
	token := c.Get("token").(string)
	// Create a context with the Authorization header for gRPC
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)
	userInfo, err := o.UserGrpcClient.Validate(ctx, &userPB.ValidateRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("User validation failed: %v", err))
	}
	userId := userInfo.ID

	err = o.OrderUsecase.SaveOrderDetail(ctx, "randomgenerated", userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Error save order id: %v", err))
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
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
