package handler

import (
	"crud-database/helper"
	"crud-database/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUsersHandler(c echo.Context) error {
	users, err := h.userService.GetUsersService()
	if err != nil {
		response := helper.APIResponse("Failed Fetch User Data", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusOK, response)
	}

	response := helper.APIResponse("Success Fetch User Data", http.StatusOK, true, users)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetUserHandler(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		response := helper.APIResponse("Failed Get Book By ID", http.StatusInternalServerError, false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	book, err := h.userService.GetUserService(id)
	if err != nil {
		response := helper.APIResponse("Failed Get Book By ID", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Success Get Book By ID", http.StatusOK, true, book)
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

	response := helper.APIResponse("Success Create New User", http.StatusOK, true, newUser)
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

	response := helper.APIResponse("Success Update User", http.StatusOK, true, updateUser)
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

	response := helper.APIResponse("Success Delete User", http.StatusOK, true, user)
	return c.JSON(http.StatusOK, response)
}