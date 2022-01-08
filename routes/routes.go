package routes

import (
	"github.com/justjundana/go-crud-mysql/handler"
	"github.com/justjundana/go-crud-mysql/helper"
	"github.com/justjundana/go-crud-mysql/repository"
	"github.com/justjundana/go-crud-mysql/service"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route User
	userRepository := repository.NewUserRepository(helper.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	e.GET("/users", userHandler.GetUsersHandler)
	e.GET("/users/:id", userHandler.GetUserHandler)
	e.POST("/users", userHandler.CreateUserHandler)
	e.PUT("/users/:id", userHandler.UpdateUserHandler)
	e.DELETE("/users/:id", userHandler.DeleteUserHandler)

	// Route Book
	bookRepository := repository.NewBookRepository(helper.DB)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	e.GET("/books", bookHandler.GetBooksHandler)
	e.POST("/books", bookHandler.CreateBookHandler)
	e.GET("/books/:id", bookHandler.GetBookHandler)
	e.PUT("/books/:id", bookHandler.UpdateBookHandler)
	e.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	return e
}
