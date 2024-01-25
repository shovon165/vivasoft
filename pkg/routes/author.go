package routes

import (
	"book-crud/pkg/controllers"
	"book-crud/pkg/middlewares"

	"github.com/labstack/echo/v4"
)

type AuthorRoutes struct {
	echo       *echo.Echo
	controller controllers.AuthorController
}

func NewAuthorRoutes(echo *echo.Echo, controller controllers.AuthorController) *AuthorRoutes {
	return &AuthorRoutes{
		echo:       echo,
		controller: controller,
	}
}

func (authorRoutes *AuthorRoutes) InitAuthorRoutes() {
	e := authorRoutes.echo

	author := e.Group("/bookstore")
	author.GET("/authors", authorRoutes.controller.GetAuthors)
	author.GET("/authors/:id", authorRoutes.controller.GetAuthor)

	author.Use(middlewares.Auth)

	author.POST("/authors", authorRoutes.controller.CreateAuthor)
	author.PUT("/authors/:id", authorRoutes.controller.UpdateAuthor)
	author.DELETE("/authors/:id", authorRoutes.controller.DeleteAuthor)
}
