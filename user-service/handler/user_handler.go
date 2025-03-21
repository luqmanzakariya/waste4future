package handler

import (
	"net/http"
	"user-service/middleware"
	"user-service/model/web/request"
	"user-service/model/web/response"
	"user-service/usecase"
	"user-service/utils"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	UserUsecase usecase.IUserUsecase
}

func NewUserUsecase(userUsecase usecase.IUserUsecase) userHandler {
	return userHandler{
		UserUsecase: userUsecase,
	}
}

func (u *userHandler) InitRoutes(g *echo.Group) {
	g.POST("/register", u.Register)
	g.POST("/login", u.Login)
	g.GET("/me", u.GetMe, middleware.AuthMiddleware)
}

// Users Create
// @Summary Create a Users
// @Description Create new Users
// @Tags Users
// @Produce json
// @Param	request	body request.UserCreateRequest true "create user payload"
// @Success 200 {object} response.WebResponse{data=response.UserResponse} "Users created"
// @Failure 400 {object} response.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Router /api/users/register [post]
func (u *userHandler) Register(c echo.Context) error {
	var payload request.UserCreateRequest

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	newUser, err := u.UserUsecase.Create(c.Request().Context(), payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   newUser,
	}

	return c.JSON(http.StatusOK, webResponse)
}

// Users Login
// @Summary Login users
// @Description Login users and get token authentication
// @Tags Users
// @Produce json
// @Param	request	body request.UserLoginRequest true "user login payload"
// @Success 200 {object} response.WebResponse{data=response.TokenResponse} "Token generated"
// @Failure 400 {object} response.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Router /api/users/login [post]
func (u *userHandler) Login(c echo.Context) error {
	var payload request.UserLoginRequest
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tokenString, err := u.UserUsecase.Login(c.Request().Context(), payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tokenString,
	}

	return c.JSON(http.StatusOK, webResponse)

}

// Users GetMe
// @Summary Get detail user
// @Description Get detail user by token
// @Tags Users
// @Produce json
// @Success 200 {object} response.WebResponse{data=response.UserResponse} "Get detail users with token"
// @Failure 400 {object} response.WebResponse{code=int,data=interface{},status=string} "Bad Request"
// @Security BearerAuth
// @Router /api/users/me [get]
func (u *userHandler) GetMe(c echo.Context) error {
	// # Get userInfo from context after JWT Validation
	userInfo, ok := c.Get("userInfo").(*utils.UserInfo)
	if !ok || userInfo == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: user info not found")
	}

	// # Use the ID from userInfo to find the user
	user := response.UserResponse{
		ID:        userInfo.ID,
		Name:      userInfo.Name,
		Email:     userInfo.Email,
		AddressID: userInfo.AddressID,
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	}

	return c.JSON(http.StatusOK, webResponse)
}
