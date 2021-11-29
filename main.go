package main

import (
	"log"
	"os"

	"bitbucket.org/houmeteam/houme-go/configs"
	"bitbucket.org/houmeteam/houme-go/database"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	env := os.Getenv("ENV")
	dbConfigs := configs.GetDBConfigs(env)
	db, err := database.ConnectToDB(*dbConfigs)

	//unable to connect to database
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
	db.AutoMigrate(&models.Property{})
	db.AutoMigrate(&models.Job{})
	db.AutoMigrate(&models.ConstructionJobProperty{})
	db.AutoMigrate(&models.ConstructionJobMaterial{})

	db.AutoMigrate(&models.ProjectJob{})
	db.AutoMigrate(&models.ProjectProperty{})
	db.AutoMigrate(&models.Project{})

	projectRepository := repositories.NewProjectRepository(db)
	propertiesRepository := repositories.NewPropertyRepository(db)
	jobsRepository := repositories.NewJobRepository(db)
	constructionJobPropertyRepository := repositories.NewConstructionJobPropertyRepository(db)
	constructionJobMaterialRepository := repositories.NewConstructionJobMaterialRepository(db)
	projectJobRepository := repositories.NewProjectJobRepository(db)
	projectPropertyRepository := repositories.NewProjectPropertyRepository(db)

	route := configs.SetupRoutes(
		projectRepository, propertiesRepository, jobsRepository,
		constructionJobPropertyRepository, constructionJobMaterialRepository,
		projectJobRepository, projectPropertyRepository)

	route.Run(":" + port)
}
