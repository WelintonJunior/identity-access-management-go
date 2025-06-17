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
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid request body",
			})
		}

		id, err := CreateRepoRegister(request)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("Failed to create record: %v", err),
			})
		}

		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"id":      id,
			"message": "Record created successfully",
			"success": true,
		})
	}
}

func ListControllerRegisters[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		filters := getGenericFilters(c)

		records, err := ListRepoRegisters[T](filters)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("Failed to list records: %v", err),
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"data":    records,
			"success": true,
		})
	}
}

func GetControllerRegisterById[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid UUID format",
			})
		}

		record, err := GetRepoRegisterById[T](id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("Failed to retrieve record: %v", err),
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"data":    record,
			"success": true,
		})
	}
}

func UpdateControllerRegisterById[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid UUID format",
			})
		}

		var request T
		if err := c.BodyParser(&request); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid request body",
			})
		}

		updated, err := UpdateRepoRegisterById(id, request)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("Failed to update record: %v", err),
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"data":    updated,
			"message": "Record updated successfully",
			"success": true,
		})
	}
}

func DeleteControllerRegisterById[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid UUID format",
			})
		}

		if err := DeleteRepoRegisterById[T](id); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("Failed to delete record: %v", err),
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Record deleted successfully",
			"success": true,
		})
	}
}
