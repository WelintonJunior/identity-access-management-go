package controllers

import (
	"github.com/WelintonJunior/identity-access-management-go/commons"
	"github.com/WelintonJunior/identity-access-management-go/types"
	"github.com/gofiber/fiber/v2"
)

// @Summary     Criar Produto
// @Description Cria um novo Produto
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       request body types.Product true "Dados do Produto"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/Product [post]
func CreateProduct() fiber.Handler {
	return commons.CreateControllerRegister[types.Product]()
}

// @Summary     Listar Produtos
// @Description Lista todos os Produtos cadastrados
// @Tags        Product
// @Accept      json
// @Produce     json
// @Success     200 {object} types.ListProductResponse
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/Product [get]
func ListProducts() fiber.Handler {
	return commons.ListControllerRegisters[types.Product]()
}

// @Summary     Buscar Produto por ID
// @Description Retorna os dados de um Produto específico
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do Produto"
// @Success     200 {object} types.ProductResponse
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/Product/{id} [get]
func GetProductById() fiber.Handler {
	return commons.GetControllerRegisterById[types.Product]()
}

// @Summary     Atualizar Produto por ID
// @Description Atualiza os dados de um Produto específico
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do Produto"
// @Param       request body types.Product true "Dados atualizados do Produto"
// @Success     200 {object} types.ProductResponse
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/Product/{id} [put]
func UpdateProductById() fiber.Handler {
	return commons.UpdateControllerRegisterById[types.Product]()
}

// @Summary     Remover Produto por ID
// @Description Remove um Produto específico
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do Produto"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/Product/{id} [delete]
func DeleteProductById() fiber.Handler {
	return commons.DeleteControllerRegisterById[types.Product]()
}
