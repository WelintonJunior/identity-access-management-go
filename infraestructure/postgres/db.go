package infraestructure

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type SqlConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

func GetSqlConfig() SqlConfig {
	sqlConfig := SqlConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DbName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	return sqlConfig
}

func NewSqlDbConnection(sqlConfig SqlConfig) (*gorm.DB, error) {
	var err error

	var (
		db_host = os.Getenv("DB_HOST")
		db_user = os.Getenv("DB_USER")
		db_pass = os.Getenv("DB_PASS")
		db_name = os.Getenv("DB_NAME")
		db_port = os.Getenv("DB_PORT")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", db_host, db_user, db_pass, db_name, db_port)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return Db, nil
}
