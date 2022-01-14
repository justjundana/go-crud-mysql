package handler

import (
	"net/http"
	"strconv"

	"github.com/justjundana/go-crud-mysql/helper"
	"github.com/justjundana/go-crud-mysql/models"
	"github.com/justjundana/go-crud-mysql/service"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(ProductService service.ProductService) *ProductHandler {
	return &ProductHandler{ProductService}
}

func (h *ProductHandler) GetProductsHandler(c echo.Context) error {
	products, err := h.productService.GetProductsService()
	if err != nil {
		response := helper.APIResponse("Failed Fetch Product Data", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusOK, response)
	}

	var data []helper.ProductFormatter
	for i := 0; i < len(products); i++ {
		formatter := helper.FormatProduct(products[i])
		data = append(data, formatter)
	}

	response := helper.APIResponse("Success Fetch Product Data", http.StatusOK, true, data)
	return c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) GetProductHandler(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		response := helper.APIResponse("Failed Get Product By ID", http.StatusInternalServerError, false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	product, err := h.productService.GetProductService(id)
	if err != nil {
		response := helper.APIResponse("Failed Get Product By ID", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatProduct(product)
	response := helper.APIResponse("Success Get Product By ID", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) CreateProductHandler(c echo.Context) error {
	var input helper.CreateProductRequest
	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Failed Create New Product", http.StatusUnprocessableEntity, false, nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	userId := c.Get("currentUser").(models.User)
	input.User = userId

	newProduct, err := h.productService.CreateProductService(input)
	if err != nil {
		response := helper.APIResponse("Failed Create New Product", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatProduct(newProduct)
	response := helper.APIResponse("Success Create New Product", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) UpdateProductHandler(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		response := helper.APIResponse("Failed Update Product", http.StatusInternalServerError, false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	var input helper.EditProductRequest
	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Failed Update Product", http.StatusUnprocessableEntity, false, nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	userId := c.Get("currentUser").(models.User)
	input.User = userId

	updateProduct, err := h.productService.UpdateProductService(id, input)
	if err != nil {
		response := helper.APIResponse("Failed Update Product", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatProduct(updateProduct)
	response := helper.APIResponse("Success Update Product", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) DeleteProductHandler(c echo.Context) error {
	id, errId := strconv.Atoi(c.Param("id"))
	if errId != nil {
		response := helper.APIResponse("Failed Delete Product", http.StatusInternalServerError, false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	userId := c.Get("currentUser").(models.User)
	currentUser := userId.ID

	product, err := h.productService.DeleteProductService(id, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed Delete Product", http.StatusBadRequest, false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := helper.FormatProduct(product)
	response := helper.APIResponse("Success Delete Product", http.StatusOK, true, formatter)
	return c.JSON(http.StatusOK, response)
}
