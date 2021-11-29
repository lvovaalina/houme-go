package database

import (
	"fmt"

	"bitbucket.org/houmeteam/houme-go/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(config configs.DBConfig) (*gorm.DB, error) {
	//sslmode=disable
	var connectionString = fmt.Sprintf(
		"host=%s port=5432 user=%s dbname=%s password=%s",
		config.Host, config.User, config.Name, config.Password,
	)

	return gorm.Open(postgres.Open(connectionString), &gorm.Config{QueryFields: true})
}
