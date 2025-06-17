package controllers

import (
	"github.com/WelintonJunior/identity-access-management-go/commons"
	"github.com/WelintonJunior/identity-access-management-go/types"
	"github.com/gofiber/fiber/v2"
)

// @Summary      List Users
// @Description  Retrieves a list of all registered users
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200 {object} types.ListUserResponse
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/users [get]
func ListUsers() fiber.Handler {
	return commons.ListControllerRegisters[types.User]()
}

// @Summary      Get User by ID
// @Description  Retrieves a specific user's details by their ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} types.UserResponse
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/users/{id} [get]
func GetUserById() fiber.Handler {
	return commons.GetControllerRegisterById[types.User]()
}

// @Summary      Update User by ID
// @Description  Updates a specific user's information
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        request body types.User true "Updated user data"
// @Success      200 {object} types.UserResponse
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/users/{id} [put]
func UpdateUserById() fiber.Handler {
	return commons.UpdateControllerRegisterById[types.User]()
}

// @Summary      Delete User by ID
// @Description  Deletes a user by their ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/users/{id} [delete]
func DeleteUserById() fiber.Handler {
	return commons.DeleteControllerRegisterById[types.User]()
}
