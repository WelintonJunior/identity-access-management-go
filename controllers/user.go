package controllers

import (
	"github.com/WelintonJunior/identity-access-management-go/commons"
	"github.com/WelintonJunior/identity-access-management-go/types"
	"github.com/gofiber/fiber/v2"
)

// @Summary     Listar Usuários
// @Description Lista todos os usuários cadastrados
// @Tags        User
// @Accept      json
// @Produce     json
// @Success     200 {object} types.ListUserResponse
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/user [get]
func ListUsers() fiber.Handler {
	return commons.ListControllerRegisters[types.User]()
}

// @Summary     Buscar Usuário por ID
// @Description Retorna os dados de um usuário específico
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do usuário"
// @Success     200 {object} types.UserResponse
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/user/{id} [get]
func GetUserById() fiber.Handler {
	return commons.GetControllerRegisterById[types.User]()
}

// @Summary     Atualizar Usuário por ID
// @Description Atualiza os dados de um usuário específico
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do usuário"
// @Param       request body types.User true "Dados atualizados do usuário"
// @Success     200 {object} types.UserResponse
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/user/{id} [put]
func UpdateUserById() fiber.Handler {
	return commons.UpdateControllerRegisterById[types.User]()
}

// @Summary     Remover Usuário por ID
// @Description Remove um usuário específico
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       id path string true "ID do usuário"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/user/{id} [delete]
func DeleteUserById() fiber.Handler {
	return commons.DeleteControllerRegisterById[types.User]()
}
