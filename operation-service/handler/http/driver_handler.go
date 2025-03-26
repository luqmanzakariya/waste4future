package handler

import (
	"fmt"
	"net/http"
	"operation-service/middleware"
	"operation-service/model"
	"operation-service/usecase"

	"github.com/labstack/echo/v4"
)

type driverHandler struct {
	DriverUsecase usecase.IDriverUsecase
}

func NewDriverUsecase(driverUsecase usecase.IDriverUsecase) driverHandler {
	return driverHandler{
		DriverUsecase: driverUsecase,
	}
}

func (d *driverHandler) InitRoutes(g *echo.Group) {
	g.GET("", d.FindAllActive, middleware.AuthMiddleware)
	g.GET("/all", d.FindAll, middleware.AuthMiddleware)
	g.POST("", d.Register, middleware.AuthMiddleware)
	g.PUT("/:id", d.Update, middleware.AuthMiddleware)
	g.DELETE("/:id", d.Delete, middleware.AuthMiddleware)
}

// Driver Create
// @Summary Create a Driver
// @Description Create new Driver
// @Tags Drivers
// @Produce json
// @Param	request	body model.PayloadCreateDriver true "create driver payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseDriver} "Driver created"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /api/drivers [post]
func (d *driverHandler) Register(c echo.Context) error {
	var payload model.PayloadCreateDriver

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newDriver, err := d.DriverUsecase.Create(c.Request().Context(), payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   newDriver,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// Driver Update
// @Summary Update a Driver
// @Description Update Driver by ID
// @Tags Drivers
// @Produce json
// @Param id path string true "driver id" example("67e1327bc481a422f0293ff9")
// @Param request body model.PayloadUpdateDriver true "update driver payload"
// @Success 200 {object} model.WebResponse{data=model.ResponseDriver} "Driver updated"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /api/drivers/{id} [put]
func (d *driverHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var payload model.PayloadUpdateDriver
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updated, err := d.DriverUsecase.Update(c.Request().Context(), id, payload)
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

// Driver Delete
// @Summary Delete a Driver
// @Description Delete a Driver
// @Tags Drivers
// @Produce json
// @Param id path string true "order detail id" example("67e1327bc481a422f0293ff9")
// @Success 200 {object} model.WebResponse{data=model.ResponseDriver} "Driver created"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /api/drivers/{id} [delete]
func (d *driverHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := d.DriverUsecase.Delete(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := fmt.Sprintf("success deleted driver detail with id: %s", id)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// Driver FindAll
// @Summary Find All Driver
// @Description Find All Driver
// @Tags Drivers
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.Driver} "List Driver"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /api/drivers/all [get]
func (d *driverHandler) FindAll(c echo.Context) error {
	drivers, err := d.DriverUsecase.FindAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   drivers,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// Driver FindAll Active
// @Summary Find All Active Driver
// @Description Find All Active Driver
// @Tags Drivers
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.Driver} "List Active Driver"
// @Failure 400 {object} model.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /api/drivers [get]
func (d *driverHandler) FindAllActive(c echo.Context) error {
	drivers, err := d.DriverUsecase.FindAllActive(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   drivers,
	}

	return c.JSON(http.StatusOK, webResponse)
}
