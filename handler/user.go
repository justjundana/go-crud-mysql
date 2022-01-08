package handler

import (
	"net/http"
	"strconv"

	"github.com/justjundana/go-crud-mysql/helper"
	"github.com/justjundana/go-crud-mysql/middleware"
	"github.com/justjundana/go-crud-mysql/service"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	authService middleware.JWTService
	userService service.UserService
}

func NewUserHandler(authService middleware.JWTService, userService service.UserService) *userHandler {
	return &userHandler{authService, userService}
}

func (h *userHandler) AuthUserHandler(c echo.Context) error {
	var input helper.LoginUserRequest
	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Failed Login User", http.StatusUnprocessableEntity, false, nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	loginUser, err := h.userService.LoginUserService(input)
	if err != nil {
		response := helper.APIResponse("Failed Login User", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	token, err := h.authService.GenerateToken(loginUser.ID)
	if err != nil {
		response := helper.APIResponse("Failed Login User", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatAuth(loginUser, token)
	response := helper.APIResponse("Success Login User", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUsersHandler(c echo.Context) error {
	users, err := h.userService.GetUsersService()
	if err != nil {
		response := helper.APIResponse("Failed Fetch User Data", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusOK, response)
	}

	var data []helper.UserFormatter
	for i := 0; i < len(users); i++ {
		formatter := helper.FormatUser(users[i])
		data = append(data, formatter)
	}

	response := helper.APIResponse("Success Fetch User Data", http.StatusOK, true, data)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUserHandler(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		response := helper.APIResponse("Failed Get User By ID", http.StatusInternalServerError, false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	user, err := h.userService.GetUserService(id)
	if err != nil {
		response := helper.APIResponse("Failed Get User By ID", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatUser(user)
	response := helper.APIResponse("Success Get User By ID", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) CreateUserHandler(c echo.Context) error {
	var input helper.CreateUserRequest
	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Failed Create New User", http.StatusUnprocessableEntity, false, nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	newUser, err := h.userService.CreateUserService(input)
	if err != nil {
		response := helper.APIResponse("Failed Create New User", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatUser(newUser)
	response := helper.APIResponse("Success Create New User", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUserHandler(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		response := helper.APIResponse("Failed Update User", http.StatusInternalServerError, false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	var input helper.EditUserRequest
	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Failed Update User", http.StatusUnprocessableEntity, false, nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	updateUser, err := h.userService.UpdateUserService(id, input)
	if err != nil {
		response := helper.APIResponse("Failed Update User", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatUser(updateUser)
	response := helper.APIResponse("Success Update User", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUserHandler(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		response := helper.APIResponse("Failed Delete User", http.StatusInternalServerError, false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	user, err := h.userService.DeleteUserService(id)
	if err != nil {
		response := helper.APIResponse("Failed Delete User", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatUser(user)
	response := helper.APIResponse("Success Delete User", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}
