package handler

import (
	"operation-service/middleware"
	"operation-service/model"
	"operation-service/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type addressHandler struct {
	AddressUsecase usecase.IAddressUsecase
}

func NewAddressHandler(addressUsecase usecase.IAddressUsecase) addressHandler {
	return addressHandler{
		AddressUsecase: addressUsecase,
	}
}

func (a addressHandler) InitRoutes(g *echo.Group) {
	g.POST("", a.Create)
	g.GET("", a.FindAll)
	g.GET("/:id", a.FindByID)
	g.PUT("/:id", a.Update, middleware.AuthMiddleware)
	g.DELETE("/:id", a.Delete, middleware.AuthMiddleware)
}

// Address Create
// @Summary Create an Addressa
// @Description Add and register new Address
// @Tags Address
// @Produce json
// @Param	request	body model.PayloadCreateAddress true "create address payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseAddress} "Address created"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /addresses [post]
func (a addressHandler) Create(c echo.Context) error {
	var payloadCreateAddress model.PayloadCreateAddress
	err := c.Bind(&payloadCreateAddress)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	created, err := a.AddressUsecase.Create(c.Request().Context(), payloadCreateAddress)
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

// Address FindAll
// @Summary FindAll Address
// @Description Find All Addresses
// @Tags Address
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.Address} "Address list"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Router /addresses [get]
func (a addressHandler) FindAll(c echo.Context) error {
	addresses, err := a.AddressUsecase.FindAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   addresses,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// Address FindById
// @Summary FindById an Address
// @Description Find Address By ID
// @Tags Address
// @Produce json
// @Param id path string true "address id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=model.ResponseAddress} "Address found"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Router /addresses/{id} [get]
func (a addressHandler) FindByID(c echo.Context) error {
	// # Get ID Params
	id := c.Param("id")

	data, err := a.AddressUsecase.FindByID(c.Request().Context(), id)
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

// Address UpdateById
// @Summary Update an Address
// @Description Update Address By ID
// @Tags Address
// @Produce json
// @Param id path string true "address id" example("67cdcb62a50a990a870d928f")
// @Param	request	body model.PayloadUpdateAddress true "update address payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseAddress} "Address updated"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /addresses/{id} [PUT]
func (a addressHandler) Update(c echo.Context) error {
	// # Get ID Params
	id := c.Param("id")

	var payload model.PayloadUpdateAddress
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedData, err := a.AddressUsecase.Update(c.Request().Context(), id, payload)
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

// Address DeleteById
// @Summary DeleteById an Address
// @Description Delete Address By ID
// @Tags Address
// @Produce json
// @Param id path string true "address id" example("67cdcb62a50a990a870d928f")
// @Success 200 {object} model.WebResponse{data=string} "Address deleted"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Failure 500 {object} model.WebResponse{code=int,data=interface{},status=string} "Internal Server Error"
// @Security BearerAuth
// @Router /addresses/{id} [DELETE]
func (a addressHandler) Delete(c echo.Context) error {
	// # Get ID Params
	id := c.Param("id")

	err := a.AddressUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := fmt.Sprintf("success deleted address with id: %s", id)

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}

	return c.JSON(http.StatusOK, webResponse)
}
