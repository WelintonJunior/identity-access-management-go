package cmd

import (
	"log"
	"os"

	infraestructure "github.com/WelintonJunior/identity-access-management-go/infraestructure/postgres"
	"github.com/WelintonJunior/identity-access-management-go/utils"
	"github.com/spf13/cobra"
)

var option string

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply or rollback database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.LoadEnvMem(); err != nil {
			log.Fatalf("Failed to load environment variables: %v", err)
		}

		sqlConf := infraestructure.SqlConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			DbName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		}

		db, err := infraestructure.NewSqlDbConnection(sqlConf)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		migrationService, err := infraestructure.NewPostgresMigrateService(db)
		if err != nil {
			log.Fatalf("Failed to initialize migration service: %v", err)
		}

		switch option {
		case "up":
			if err := migrationService.MigrateApply(); err != nil {
				log.Fatalf("Failed to apply migrations: %v", err)
			}

			if err := infraestructure.Seed(db); err != nil {
				log.Fatal("Failed to apply Seed:", err)
			}

			log.Println("Migrations applied successfully")
		case "down":
			if err := migrationService.MigrateRevert(); err != nil {
				log.Fatalf("Failed to revert migrations: %v", err)
			}
			log.Println("Migrations reverted successfully")
		default:
			log.Println("Invalid operation. Use --operation=up or --operation=down")
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().StringVarP(&option, "operation", "o", "", "Migration operation: up or down")
}
