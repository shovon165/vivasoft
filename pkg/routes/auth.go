package routes

import (
	"book-crud/pkg/controllers"

	"github.com/labstack/echo/v4"
)

type AuthRoutes struct {
	echo    *echo.Echo
	authCtr controllers.AuthController
}

func NewAuthRoutes(echo *echo.Echo, authCtr controllers.AuthController) *AuthRoutes {
	return &AuthRoutes{
		echo:    echo,
		authCtr: authCtr,
	}
}

func (routes *AuthRoutes) InitAuthRoutes() {
	e := routes.echo

	e.POST("/login", routes.authCtr.Login)
	e.POST("/signup", routes.authCtr.Signup)
}
