package routes

import (
	"book-crud/pkg/controllers"
	"book-crud/pkg/middlewares"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookRoutes struct {
	echo       *echo.Echo
	controller controllers.BookController
}

func NewBookRoutes(echo *echo.Echo, controller controllers.BookController) *BookRoutes {
	return &BookRoutes{
		echo:       echo,
		controller: controller,
	}
}

func (bookRoutes *BookRoutes) InitBookRoutes() {
	e := bookRoutes.echo

	e.GET("/ping", Pong)

	book := e.Group("/bookstore")
	book.GET("/books", bookRoutes.controller.GetBook)
	book.GET("/books/:id", bookRoutes.controller.GetBook)

	book.Use(middlewares.Auth)

	book.POST("/books", bookRoutes.controller.CreateBook)
	book.PUT("/books/:id", bookRoutes.controller.UpdateBook)
	book.DELETE("/books/:id", bookRoutes.controller.DeleteBook)
}

// func (bc *bookRoutes) InitBookRoutes(e *echo.Echo) {
// 	//grouping route endpoints
// 	book := e.Group("/bookstore")

// 	book.GET("/ping", Pong)
// 	//book.Use(middlewares.Demo)

// 	//initializing http methods - routing endpoints and their handlers
// 	//book.POST("/book", bc.bookCtr.CreateBook, middlewares.Auth) // middleware can be added like this
// 	book.POST("/book", bc.bookCtr.CreateBook)
// 	book.GET("/book", bc.bookCtr.GetBook)
// 	book.PUT("/book/:bookID", bc.bookCtr.UpdateBook)
// 	book.DELETE("/book/:bookID", bc.bookCtr.DeleteBook)

// 	//login
// 	book.POST("/auth/login", bc.bookCtr.Login)
// }

func Pong(ctx echo.Context) error {
	fmt.Println("Pong")
	return ctx.JSON(http.StatusOK, "Pong")
}
