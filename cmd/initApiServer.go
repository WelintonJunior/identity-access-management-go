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
	postgres "github.com/WelintonJunior/identity-access-management-go/infraestructure/postgres"
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
	Short: "Inicializa o serviço web de API",
	Long:  "Comando para realizar a execução da aplicação",
	Run: func(cmd *cobra.Command, args []string) {
		app := SetupApp(context.Background())

		portStr := os.Getenv("API_PORT")

		apiPort, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatalf("Variavel de ambiente API_PORT inválido: %v", err)
		}

		go func() {
			runPort := fmt.Sprintf(":%d", apiPort)
			if err := app.Listen(runPort); err != nil {
				log.Panic(err)
			}
		}()

		cancelChan := make(chan os.Signal, 1)
		signal.Notify(cancelChan, os.Interrupt, syscall.SIGTERM)

		<-cancelChan
		fmt.Println("Gracefully shutdown")
		_ = app.Shutdown()
	},
}

func init() {
	rootCmd.AddCommand(initApiServerCmd)
	initApiServerCmd.Flags().IntVarP(&apiPort, "port", "p", apiPort, "Http server execution port")
}

func SetupApp(ctx context.Context) *fiber.App {
	err := utils.LoadEnvMem()
	if err != nil {
		log.Fatal("No .env file found")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	_, err = postgres.NewSqlDbConnection(infraestructure.GetSqlConfig())

	if err != nil {
		fmt.Println("Erro ao conectar no Redis:", err)
	}

	rdb := redis.InitRedis()

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Erro ao conectar no Redis:", err)
	} else {
		fmt.Println("Conectado ao Redis com sucesso!")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint",
		})
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	authRoutes := app.Group("api/v1")

	routes.AuthRoutes(authRoutes)

	apiV1 := app.Group("api/v1", middlewares.JwtAuth())
	routes.UserRoutes(apiV1)
	routes.ProductRoutes(apiV1)

	return app
}
