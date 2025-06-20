package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/WelintonJunior/identity-access-management-go/cmd/middlewares"
	_ "github.com/WelintonJunior/identity-access-management-go/docs"
	infraestructure "github.com/WelintonJunior/identity-access-management-go/infraestructure/postgres"
	redis "github.com/WelintonJunior/identity-access-management-go/infraestructure/redis"
	"github.com/WelintonJunior/identity-access-management-go/routes"
	"github.com/WelintonJunior/identity-access-management-go/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/spf13/cobra"
)

var apiPort int

var initApiServerCmd = &cobra.Command{
	Use:   "initApiServer",
	Short: "Starts the API web server",
	Run: func(cmd *cobra.Command, args []string) {
		app := SetupApp(context.Background())

		portStr := os.Getenv("API_PORT")
		port, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatalf("Invalid API_PORT environment variable: %v", err)
		}

		go func() {
			addr := fmt.Sprintf(":%d", port)
			if err := app.Listen(addr); err != nil {
				log.Panicf("Server failed to start: %v", err)
			}
		}()

		// Wait for termination signal
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig

		log.Println("Shutting down gracefully...")
		if err := app.Shutdown(); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initApiServerCmd)
	initApiServerCmd.Flags().IntVarP(&apiPort, "port", "p", apiPort, "HTTP server port")
}

func SetupApp(ctx context.Context) *fiber.App {
	if err := utils.LoadEnvMem(); err != nil {
		log.Fatal("Environment variables could not be loaded")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	if _, err := infraestructure.NewSqlDbConnection(infraestructure.GetSqlConfig()); err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	rdb := redis.InitRedis()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	} else {
		log.Println("Connected to Redis successfully")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "API is running",
		})
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	// auth.SetupSSO()

	// Rotas públicas de autenticação
	auth := app.Group("/api/v1")
	routes.AuthRoutes(auth)

	// Rotas protegidas, exigem JWT e autenticação
	protected := app.Group("/api/v1")
	protected.Use(middlewares.RequireAuth())
	routes.ProductRoutes(protected)

	// Rotas administrativas, requerem permissão 'admin'
	adminGroup := protected.Group("/admin")
	adminGroup.Use(middlewares.RequirePermission("admin"))
	routes.UserRoutes(adminGroup)

	return app
}
