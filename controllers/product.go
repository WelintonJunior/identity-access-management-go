package controllers

import (
	"github.com/WelintonJunior/identity-access-management-go/commons"
	"github.com/WelintonJunior/identity-access-management-go/types"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Create Product
// @Description  Creates a new product in the system
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        request body types.Product true "Product data"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/products [post]
func CreateProduct() fiber.Handler {
	return commons.CreateControllerRegister[types.Product]()
}

// @Summary      List Products
// @Description  Retrieves a list of all registered products
// @Tags         Product
// @Accept       json
// @Produce      json
// @Success      200 {object} types.ListProductResponse
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/products [get]
func ListProducts() fiber.Handler {
	return commons.ListControllerRegisters[types.Product]()
}

// @Summary      Get Product by ID
// @Description  Retrieves a product's information by its ID
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id path string true "Product ID"
// @Success      200 {object} types.ProductResponse
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/products/{id} [get]
func GetProductById() fiber.Handler {
	return commons.GetControllerRegisterById[types.Product]()
}

// @Summary      Update Product by ID
// @Description  Updates the details of a specific product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id path string true "Product ID"
// @Param        request body types.Product true "Updated product data"
// @Success      200 {object} types.ProductResponse
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/products/{id} [put]
func UpdateProductById() fiber.Handler {
	return commons.UpdateControllerRegisterById[types.Product]()
}

// @Summary      Delete Product by ID
// @Description  Removes a product by its ID
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id path string true "Product ID"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/products/{id} [delete]
func DeleteProductById() fiber.Handler {
	return commons.DeleteControllerRegisterById[types.Product]()
}
