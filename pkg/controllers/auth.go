package controllers

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IAuthController interface {
	Login(e echo.Context) error
	Signup(e echo.Context) error
}

type AuthController struct {
	authSvc domain.IAuthService
}

func NewAuthController(authSvc domain.IAuthService) AuthController {
	return AuthController{
		authSvc: authSvc,
	}
}

func (authController *AuthController) Login(e echo.Context) error {
	loginRequest := &types.LoginRequest{}
	if err := e.Bind(loginRequest); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}

	if err := loginRequest.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	loginResponse, err := authController.authSvc.LoginUser(loginRequest)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, loginResponse)
}

func (authController *AuthController) Signup(e echo.Context) error {
	registerRequest := &types.SignupRequest{}
	if err := e.Bind(registerRequest); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid request body")
	}

	if err := registerRequest.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := authController.authSvc.SignupUser(registerRequest); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "user was created successfully")
}
