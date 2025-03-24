package handler

import (
	"fmt"
	"net/http"
	"operation-service/middleware" // Assuming you have a similar middleware package
	"operation-service/model"
	pbUser "operation-service/pb/user"
	"operation-service/usecase"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
)

type orderDetailHandler struct {
	OrderDetailUsecase usecase.IOrderDetailUsecase
	UserServiceClient  pbUser.UserServiceClient
}

func NewOrderDetailHandler(orderDetailUsecase usecase.IOrderDetailUsecase, userServiceClient pbUser.UserServiceClient) orderDetailHandler {
	return orderDetailHandler{
		OrderDetailUsecase: orderDetailUsecase,
		UserServiceClient:  userServiceClient,
	}
}

func (h orderDetailHandler) InitRoutes(g *echo.Group) {
	g.POST("", h.Create, middleware.AuthMiddleware)
	g.GET("/:id", h.FindByID, middleware.AuthMiddleware)
	g.PUT("/:id", h.Update, middleware.AuthMiddleware)
	g.DELETE("/:id", h.Delete, middleware.AuthMiddleware)
}

// OrderDetail Create
// @Summary Create an Order Detail
// @Description Add and register new order detail
// @Tags OrderDetail
// @Produce json
// @Param request body model.PayloadCreateOrderDetail true "create order detail payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseOrderDetail} "Order detail created"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /order-details [post]
func (h orderDetailHandler) Create(c echo.Context) error {
	token := c.Get("token").(string)
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)

	// Validate token and get user info
	userResp, err := h.UserServiceClient.Validate(ctx, &pbUser.ValidateRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("User validation failed: %v", err))
	}

	var payload model.PayloadCreateOrderDetail
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Assuming userResp contains UserId (adjust based on your proto definition)
	userID := int(userResp.ID) // Adjust this based on your ValidateResponse proto

	created, err := h.OrderDetailUsecase.Create(ctx, userID, payload)
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

// OrderDetail FindByID
// @Summary Find an Order Detail by ID
// @Description Find order detail by ID
// @Tags OrderDetail
// @Produce json
// @Param id path string true "order detail id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=model.ResponseOrderDetail} "Order detail found"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /order-details/{id} [get]
func (h orderDetailHandler) FindByID(c echo.Context) error {
	token := c.Get("token").(string)
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)

	_, err := h.UserServiceClient.Validate(ctx, &pbUser.ValidateRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("User validation failed: %v", err))
	}

	id := c.Param("id")
	data, err := h.OrderDetailUsecase.FindByID(ctx, id)
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

// OrderDetail Update
// @Summary Update an Order Detail
// @Description Update order detail by ID
// @Tags OrderDetail
// @Produce json
// @Param id path string true "order detail id" example("67cdcb62a50a990a870d928f")
// @Param request body model.PayloadUpdateOrderDetail true "update order detail payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseOrderDetail} "Order detail updated"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /order-details/{id} [put]
func (h orderDetailHandler) Update(c echo.Context) error {
	token := c.Get("token").(string)
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)

	_, err := h.UserServiceClient.Validate(ctx, &pbUser.ValidateRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("User validation failed: %v", err))
	}

	id := c.Param("id")
	var payload model.PayloadUpdateOrderDetail
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updated, err := h.OrderDetailUsecase.Update(ctx, id, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   updated,
	}
	return c.JSON(http.StatusOK, webResponse)
}

// OrderDetail Delete
// @Summary Delete an Order Detail
// @Description Delete order detail by ID
// @Tags OrderDetail
// @Produce json
// @Param id path string true "order detail id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=string} "Order detail deleted"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /order-details/{id} [delete]
func (h orderDetailHandler) Delete(c echo.Context) error {
	token := c.Get("token").(string)
	ctx := metadata.AppendToOutgoingContext(c.Request().Context(), "authorization", token)

	_, err := h.UserServiceClient.Validate(ctx, &pbUser.ValidateRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("User validation failed: %v", err))
	}

	id := c.Param("id")
	if err := h.OrderDetailUsecase.Delete(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := fmt.Sprintf("success deleted order detail with id: %s", id)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}
	return c.JSON(http.StatusOK, webResponse)
}
