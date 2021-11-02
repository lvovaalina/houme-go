package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		"host=localhost port=5432 user=%s dbname=%s sslmode=disable password=%s",
		dbUser, dbName, dbPassword,
	)

	return gorm.Open(postgres.Open(connectionString), &gorm.Config{QueryFields: true})
}
