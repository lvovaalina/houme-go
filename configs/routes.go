package configs

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"bitbucket.org/houmeteam/houme-go/helpers"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
	"bitbucket.org/houmeteam/houme-go/services"
)

func SetupRoutes(
	projectRepository *repositories.ProjectRepository,
	propertiesRepository *repositories.PropertyRepository,
	jobsRepository *repositories.JobRepository,
	constructionJobPropertiesRepository *repositories.ConstructionJobPropertyRepository) *gin.Engine {
	route := gin.Default()

	route.Use(gin.Logger())

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Length", "Content-Type", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	route.POST("/create", func(context *gin.Context) {
		// initialization contact model
		var project models.Project

		// validate json
		err := context.ShouldBindJSON(&project)

		// validation errors
		if err != nil {
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save contact & get it's response
		response := services.CreateProject(&project, *projectRepository, *constructionJobPropertiesRepository)

		// save contact failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/getProperties", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllProperties(*propertiesRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/getJobs", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllJobs(*jobsRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/getProjects", func(context *gin.Context) {
		code := http.StatusOK

		response := services.GetAllProjects(*projectRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	return route
}
