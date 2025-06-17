package commons

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateControllerRegister[T HasID]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request T

		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": err})
		}

		id, err := CreateRepoRegister(request)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("failed to find registers, %v", err)})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"id":      id,
			"message": "success",
			"success": true,
		})
	}
}

func ListControllerRegisters[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userFilter := getGenericFilters(c)

		registers, err := ListRepoRegisters[T](userFilter)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("failed to find registers, %v", err),
			})
		}

		if len(registers) == 0 {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Nenhum registro encontrado",
				"success": true,
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": registers,
			"success": true,
		})
	}
}

func GetControllerRegisterById[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {

		idParam := c.Params("id")

		strUUID, err := uuid.Parse(idParam)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("failed to parse uuid, %v", err)})
		}

		register, err := GetRepoRegisterById[T](strUUID)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("failed to find registers, %v", err)})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": register,
			"success": true,
		})
	}
}

func UpdateControllerRegisterById[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")

		strUUID, err := uuid.Parse(idParam)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("failed to parse uuid, %v", err)})
		}

		var request T
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": err})
		}

		register, err := UpdateRepoRegisterById(strUUID, request)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("failed to find registers, %v", err)})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": register,
			"success": true,
		})

	}
}

func DeleteControllerRegisterById[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")

		strUUID, err := uuid.Parse(idParam)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("failed to parse uuid, %v", err)})
		}

		err = DeleteRepoRegisterById[T](strUUID)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("failed to find registers, %v", err)})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "success",
			"success": true,
		})

	}
}
