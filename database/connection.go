package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(dbHost string, dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		"host=%s port=5432 user=%s dbname=%s password=%s",
		dbHost, dbUser, dbName, dbPassword,
	)

	return gorm.Open(postgres.Open(connectionString), &gorm.Config{QueryFields: true})
}
