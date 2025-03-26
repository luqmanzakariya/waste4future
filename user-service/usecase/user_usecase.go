package usecase

import (
	"context"
	"errors"
	"user-service/helper"
	"user-service/model/domain"
	web "user-service/model/web/request"
	"user-service/model/web/response"
	"user-service/repository"
	"user-service/utils"

	"github.com/go-playground/validator/v10"
)

type IUserUsecase interface {
	FindAll(ctx context.Context) ([]response.UserResponse, error)
	FindById(ctx context.Context, id int) (response.UserResponse, error)
	Create(ctx context.Context, payload web.UserCreateRequest) (response.UserResponse, error)
	Update(ctx context.Context, payload web.UserUpdateRequest, id int) (response.UserResponse, error)
	Delete(ctx context.Context, id int) error
	Login(ctx context.Context, payload web.UserLoginRequest) (response.TokenResponse, error)
}

type userUsecase struct {
	UserRepo repository.IUserRepository
	Validate *validator.Validate
}

func NewUserUsecase(userRepo repository.IUserRepository, validate *validator.Validate) IUserUsecase {
	return &userUsecase{
		UserRepo: userRepo,
		Validate: validate,
	}
}

func (u *userUsecase) FindAll(ctx context.Context) ([]response.UserResponse, error) {
	users, err := u.UserRepo.FindAll(ctx)
	if err != nil {
		return []response.UserResponse{}, err
	}
	return helper.UserToResponses(users), err
}

func (u *userUsecase) FindById(ctx context.Context, id int) (response.UserResponse, error) {
	var user domain.User

	user, err := u.UserRepo.FindById(ctx, id)
	if err != nil {
		return response.UserResponse{}, err
	}

	return helper.UserToResponse(user), nil
}

func (u *userUsecase) Create(ctx context.Context, payload web.UserCreateRequest) (response.UserResponse, error) {
	err := u.Validate.Struct(payload)
	if err != nil {
		return response.UserResponse{}, err
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return response.UserResponse{}, err
	}

	newUser := domain.User{
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  hashedPassword,
		AddressID: payload.AddressID,
	}

	userCreated, err := u.UserRepo.Create(ctx, newUser)
	if err != nil {
		return response.UserResponse{}, err
	}

	err = utils.SendEmail(payload.Email, "Welcome to Waste4Future", "your account has been activated, you cant start contribute to society by taking care of your waste")
	if err != nil {
		return response.UserResponse{}, err
	}

	return helper.UserToResponse(userCreated), nil
}

func (u *userUsecase) Update(ctx context.Context, payload web.UserUpdateRequest, id int) (response.UserResponse, error) {
	err := u.Validate.Struct(payload)
	if err != nil {
		return response.UserResponse{}, err
	}

	user, err := u.UserRepo.FindById(ctx, id)
	if err != nil {
		return response.UserResponse{}, err
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return response.UserResponse{}, err
	}

	updatedUser := domain.User{
		ID:        user.ID,
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  hashedPassword,
		AddressID: payload.AddressID,
	}

	updatedUser, err = u.UserRepo.Update(ctx, updatedUser)
	if err != nil {
		return response.UserResponse{}, err
	}

	return helper.UserToResponse(updatedUser), nil
}

func (u *userUsecase) Delete(ctx context.Context, id int) error {
	user, err := u.UserRepo.FindById(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}

	return u.UserRepo.Delete(ctx, user)
}

func (u *userUsecase) Login(ctx context.Context, payload web.UserLoginRequest) (response.TokenResponse, error) {
	err := u.Validate.Struct(payload)
	if err != nil {
		return response.TokenResponse{}, err
	}

	user, err := u.UserRepo.FindByEmail(ctx, payload.Email)
	if err != nil {
		return response.TokenResponse{}, errors.New("user not found")
	}

	checkPassword := utils.CheckPassword(user.Password, payload.Password)
	if !checkPassword {
		return response.TokenResponse{}, errors.New("invalid email / password")
	}

	tokenString, err := utils.GenerateJWT(user.ID, user.Name, user.Email, user.AddressID)
	if err != nil {
		return response.TokenResponse{}, errors.New("failed generate jwt")
	}

	return response.TokenResponse{Token: tokenString}, nil
}
