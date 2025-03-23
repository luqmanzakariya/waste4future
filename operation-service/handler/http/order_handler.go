package handler

import (
	"fmt"
	"net/http"
	"operation-service/middleware"
	"operation-service/model"
	"operation-service/usecase"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	OrderUsecase usecase.IOrderUsecase
}

func NewOrderHandler(orderUsecase usecase.IOrderUsecase) orderHandler {
	return orderHandler{
		OrderUsecase: orderUsecase,
	}
}

func (o orderHandler) InitRoutes(g *echo.Group) {
	g.POST("", o.Create)
	g.GET("", o.FindAll)
	g.GET("/:id", o.FindByID)
	g.PUT("/:id", o.Update, middleware.AuthMiddleware)
	g.DELETE("/:id", o.Delete, middleware.AuthMiddleware)
}

// Order Create
// @Summary Create an Order
// @Description Add and register new Order
// @Tags Orders
// @Produce json
// @Param	request	body model.PayloadCreateOrder true "create order payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseAddress} "Order created"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /orders [post]
func (o orderHandler) Create(c echo.Context) error {
	var payloadCreateOrder model.PayloadCreateOrder
	err := c.Bind(&payloadCreateOrder)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	created, err := o.OrderUsecase.Create(c.Request().Context(), payloadCreateOrder)
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
// @Success 200 {object} model.WebResponse{data=[]model.Address} "Address list"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Router /orders [get]
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
// @Router /order/{id} [get]
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
// @Router /orders/{id} [PUT]
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
// @Router /orders/{id} [DELETE]
func (o orderHandler) Delete(c echo.Context) error {
	// # Get ID Params
	id := c.Param("id")

	err := o.OrderUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := fmt.Sprintf("success deleted order with id: %s", id)

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}

	return c.JSON(http.StatusOK, webResponse)
}
