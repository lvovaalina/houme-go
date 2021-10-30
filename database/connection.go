package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func ConnectToDB(dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		"host=localhost port=5432 user=%s dbname=%s sslmode=disable password=%s",
		dbUser, dbName, dbPassword,
	)

	return gorm.Open("postgres", connectionString)
}
