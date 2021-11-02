package main

import (
	"log"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"

	"bitbucket.org/houmeteam/houme-go/configs"
	"bitbucket.org/houmeteam/houme-go/database"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	// database configs
	//dbUser, dbPassword, dbName := "postgres", "l8397040", "houmly"

	dbHost, dbUser, dbPassword, dbName :=
		"ec2-52-22-81-147.compute-1.amazonaws.com", "soxoxijvmbhqiv", "0ab277b623defd4ca7a72cba84bc60f06d7cabb6a8b311bc7580250bcef78b69", "ddnmu64tjqh9ju"

	db, err := database.ConnectToDB(dbHost, dbUser, dbPassword, dbName)

	// unable to connect to database
	if err != nil {
		log.Fatalln(err)
	}

	// ping to database
	//err = db..Ping()

	// error ping to database
	if err != nil {
		log.Fatalln(err)
	}

	// migration
	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Property{})
	db.AutoMigrate(&models.ProjectProperty{})
	db.AutoMigrate(&models.Job{})
	db.AutoMigrate(&models.ConstructionJobProperty{})
	db.AutoMigrate(&models.ProjectJob{})

	//defer db.Close()

	projectRepository := repositories.NewProjectRepository(db)
	propertiesRepository := repositories.NewPropertyRepository(db)
	jobsRepository := repositories.NewJobRepository(db)
	constructionJobPropertyRepository := repositories.NewConstructionJobPropertyRepository(db)

	route := configs.SetupRoutes(
		projectRepository, propertiesRepository, jobsRepository, constructionJobPropertyRepository)

	route.Run(":" + port)
}
