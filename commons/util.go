package commons

import "github.com/gofiber/fiber/v2"

func getGenericFilters(c *fiber.Ctx) map[string]interface{} {
	filters := make(map[string]interface{})

	if userID := c.QueryInt("user_id"); userID != 0 {
		filters["user_id"] = userID
	}

	// Adicione outros filtros aqui, se necessário:
	// if status := c.Query("status"); status != "" {
	//     filters["status"] = status
	// }

	return filters
}
