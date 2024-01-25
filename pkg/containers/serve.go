package containers

import (
	"book-crud/pkg/config"
	"book-crud/pkg/connection"
	"book-crud/pkg/controllers"
	"book-crud/pkg/repositories"
	"book-crud/pkg/routes"
	"book-crud/pkg/services"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {

	//config initialization
	config.SetConfig()

	//database initializations
	db := connection.GetDB()

	// repository initialization
	bookRepo := repositories.BookDBInstance(db)
	authorRepo := repositories.AuthorDBInstance(db)
	userRepo := repositories.UserDBInstance(db)

	//service initialization
	bookService := services.BookServiceInstance(bookRepo, authorRepo)
	authorService := services.AuthorServiceInstance(authorRepo, bookRepo)
	authService := services.AuthServiceInstance(userRepo)

	//controller initialization
	bookCtr := controllers.NewBookController(bookService)
	authorCtr := controllers.NewAuthorController(authorService)
	authCtr := controllers.NewAuthController(authService)

	//route initialization
	bookRoutes := routes.NewBookRoutes(e, bookCtr)
	authorRoutes := routes.NewAuthorRoutes(e, authorCtr)
	authRoutes := routes.NewAuthRoutes(e, authCtr)

	//route binding
	bookRoutes.InitBookRoutes()
	authorRoutes.InitAuthorRoutes()
	authRoutes.InitAuthRoutes()

	// starting server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))

}
