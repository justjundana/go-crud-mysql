package handler

import (
	"net/http"
	"strconv"

	"github.com/justjundana/go-crud-mysql/helper"
	"github.com/justjundana/go-crud-mysql/service"
	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	bookService service.BookService
}

func NewBookHandler(BookService service.BookService) *bookHandler {
	return &bookHandler{BookService}
}

func (h *bookHandler) GetBooksHandler(c echo.Context) error {
	books, err := h.bookService.GetBooksService()
	if err != nil {
		response := helper.APIResponse("Failed Fetch Book Data", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusOK, response)
	}

	var data []helper.BookFormatter
	for i := 0; i < len(books); i++ {
		formatter := helper.FormatBook(books[i])
		data = append(data, formatter)
	}

	response := helper.APIResponse("Success Fetch Book Data", http.StatusOK, true, data)
	return c.JSON(http.StatusOK, response)
}

func (h *bookHandler) GetBookHandler(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		response := helper.APIResponse("Failed Get Book By ID", http.StatusInternalServerError, false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	book, err := h.bookService.GetBookService(id)
	if err != nil {
		response := helper.APIResponse("Failed Get Book By ID", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatBook(book)
	response := helper.APIResponse("Success Get Book By ID", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *bookHandler) CreateBookHandler(c echo.Context) error {
	var input helper.CreateBookRequest
	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Failed Create New Book", http.StatusUnprocessableEntity, false, nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	newBook, err := h.bookService.CreateBookService(input)
	if err != nil {
		response := helper.APIResponse("Failed Create New Book", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatBook(newBook)
	response := helper.APIResponse("Success Create New Book", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *bookHandler) UpdateBookHandler(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		response := helper.APIResponse("Failed Update Book", http.StatusInternalServerError, false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	var input helper.EditBookRequest
	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Failed Update Book", http.StatusUnprocessableEntity, false, nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	updateBook, err := h.bookService.UpdateBookService(id, input)
	if err != nil {
		response := helper.APIResponse("Failed Update Book", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatBook(updateBook)
	response := helper.APIResponse("Success Update Book", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *bookHandler) DeleteBookHandler(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		response := helper.APIResponse("Failed Delete Book", http.StatusInternalServerError, false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	book, err := h.bookService.DeleteBookService(id)
	if err != nil {
		response := helper.APIResponse("Failed Delete Book", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatBook(book)
	response := helper.APIResponse("Success Delete Book", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}
