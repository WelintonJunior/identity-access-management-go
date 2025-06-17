package routes

import (
	"github.com/WelintonJunior/identity-access-management-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(route fiber.Router) {
	productGroup := route.Group("/products")
	productGroup.Post("/", controllers.CreateProduct())
	productGroup.Get("/", controllers.ListProducts())
	productGroup.Get("/:id", controllers.GetProductById())
	productGroup.Put("/:id", controllers.UpdateProductById())
	productGroup.Delete("/:id", controllers.DeleteProductById())
}
