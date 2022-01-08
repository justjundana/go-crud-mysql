package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/justjundana/go-crud-mysql/helper"
	"github.com/justjundana/go-crud-mysql/service"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(authService JWTService, userService service.UserService, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, false, nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, false, nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, false, nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		userID := int(payload["id"].(float64))
		user, err := userService.GetUserService(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, false, nil)
			return c.JSON(http.StatusUnauthorized, response)
		}
		c.Set("currentUser", user)
		return next(c)
	}
}
