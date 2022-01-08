package routes

import (
	"github.com/justjundana/go-crud-mysql/handler"
	"github.com/justjundana/go-crud-mysql/helper"
	"github.com/justjundana/go-crud-mysql/middleware"
	"github.com/justjundana/go-crud-mysql/repository"
	"github.com/justjundana/go-crud-mysql/service"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	e.Pre(echoMiddleware.RemoveTrailingSlash())
	authService := middleware.AuthService()
	// Route User
	userRepository := repository.NewUserRepository(helper.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(authService, userService)

	e.POST("/auth", userHandler.AuthUserHandler)
	e.GET("/users", middleware.AuthMiddleware(authService, userService, userHandler.GetUsersHandler))
	e.GET("/users/:id", middleware.AuthMiddleware(authService, userService, userHandler.GetUserHandler))
	e.POST("/users", userHandler.CreateUserHandler)
	e.PUT("/users/:id", middleware.AuthMiddleware(authService, userService, userHandler.UpdateUserHandler))
	e.DELETE("/users/:id", middleware.AuthMiddleware(authService, userService, userHandler.DeleteUserHandler))

	// Route Book
	bookRepository := repository.NewBookRepository(helper.DB)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	e.GET("/books", bookHandler.GetBooksHandler)
	e.POST("/books", middleware.AuthMiddleware(authService, userService, bookHandler.CreateBookHandler))
	e.GET("/books/:id", bookHandler.GetBookHandler)
	e.PUT("/books/:id", middleware.AuthMiddleware(authService, userService, bookHandler.UpdateBookHandler))
	e.DELETE("/books/:id", middleware.AuthMiddleware(authService, userService, bookHandler.DeleteBookHandler))

	return e
}
