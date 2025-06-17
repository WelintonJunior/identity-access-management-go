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
	Short: "Comando para aplicar/remover atualizações do banco",
	Long:  ">",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.LoadEnvMem(); err != nil {
			panic(err)
		}

		sqlConf := infraestructure.SqlConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			DbName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		}

		gormDb, err := infraestructure.NewSqlDbConnection(sqlConf)

		if err != nil {
			panic(err)
		}

		sqlService, err := infraestructure.NewPostgresMigrateService(gormDb)

		if err != nil {
			log.Fatalf("Erro ao criar serviço para execução das migrations: %v", err)
		}

		if option == "up" {
			if err := sqlService.MigrateApply(); err != nil {
				log.Println("Erro ao executar migrate apply")
			}

			return
		}

		if option == "down" {
			if err := sqlService.MigrateRevert(); err != nil {
				log.Println("Erro ao executar migrate revert")
			}

			return
		}

	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().StringVarP(&option, "operation", "o", "", "Migration operation: up or down")
}
