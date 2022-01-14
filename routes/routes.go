package routes

import (
	"database/sql"

	"github.com/justjundana/go-crud-mysql/handler"
	"github.com/justjundana/go-crud-mysql/middleware"
	"github.com/justjundana/go-crud-mysql/repository"
	"github.com/justjundana/go-crud-mysql/service"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Router(db *sql.DB) *echo.Echo {
	// create a new echo instance
	e := echo.New()
	e.Pre(echoMiddleware.RemoveTrailingSlash())
	authService := middleware.AuthService()
	// Route User
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(authService, userService)

	e.POST("/auth", userHandler.AuthUserHandler)
	e.GET("/users", middleware.AuthMiddleware(authService, userService, userHandler.GetUsersHandler))
	e.GET("/users/:id", middleware.AuthMiddleware(authService, userService, userHandler.GetUserHandler))
	e.POST("/users", userHandler.CreateUserHandler)
	e.PUT("/users/:id", middleware.AuthMiddleware(authService, userService, userHandler.UpdateUserHandler))
	e.DELETE("/users/:id", middleware.AuthMiddleware(authService, userService, userHandler.DeleteUserHandler))

	// Route Book
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	e.GET("/books", bookHandler.GetBooksHandler)
	e.POST("/books", middleware.AuthMiddleware(authService, userService, bookHandler.CreateBookHandler))
	e.GET("/books/:id", bookHandler.GetBookHandler)
	e.PUT("/books/:id", middleware.AuthMiddleware(authService, userService, bookHandler.UpdateBookHandler))
	e.DELETE("/books/:id", middleware.AuthMiddleware(authService, userService, bookHandler.DeleteBookHandler))

	// Route Product
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	e.GET("/products", productHandler.GetProductsHandler)
	e.POST("/products", middleware.AuthMiddleware(authService, userService, productHandler.CreateProductHandler))
	e.GET("/products/:id", productHandler.GetProductHandler)
	e.PUT("/products/:id", middleware.AuthMiddleware(authService, userService, productHandler.UpdateProductHandler))
	e.DELETE("/products/:id", middleware.AuthMiddleware(authService, userService, productHandler.DeleteProductHandler))

	return e
}
