package controllers

import (
	"log"
	"net/http"

	"bitbucket.org/houmeteam/houme-go/helpers"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
	"bitbucket.org/houmeteam/houme-go/services"
	"github.com/gin-gonic/gin"
)

type ConstructionPropertiesController struct {
	projectRepository                   *repositories.ProjectRepository
	constructionJobPropertiesRepository *repositories.ConstructionJobPropertyRepository
	projectJobRepository                *repositories.ProjectJobRepository
	projectPropertyRepository           *repositories.ProjectPropertyRepository
	jobsRepository                      *repositories.JobRepository
	constructionJobMaterialsRepository  *repositories.ConstructionJobMaterialRepository
	projectMaterialRepository           *repositories.ProjectMaterialRepository
}

func NewConstructionPropertiesController(
	projectRepository *repositories.ProjectRepository,
	constructionJobPropertiesRepository *repositories.ConstructionJobPropertyRepository,
	projectJobRepository *repositories.ProjectJobRepository,
	projectPropertyRepository *repositories.ProjectPropertyRepository,
	jobsRepository *repositories.JobRepository,
	constructionJobMaterialsRepository *repositories.ConstructionJobMaterialRepository,
	projectMaterialRepository *repositories.ProjectMaterialRepository) *ConstructionPropertiesController {
	return &ConstructionPropertiesController{
		projectRepository:                   projectRepository,
		constructionJobPropertiesRepository: constructionJobPropertiesRepository,
		projectJobRepository:                projectJobRepository,
		projectPropertyRepository:           projectPropertyRepository,
		jobsRepository:                      jobsRepository,
		constructionJobMaterialsRepository:  constructionJobMaterialsRepository,
		projectMaterialRepository:           projectMaterialRepository,
	}
}

func (c *ConstructionPropertiesController) GetJobPropertiesHandler(context *gin.Context) {
	code := http.StatusOK

	response := services.FindJobProperties(c.constructionJobPropertiesRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ConstructionPropertiesController) CreateMaterialHandler(context *gin.Context) {
	var material models.ConstructionJobMaterial

	// validate json
	err := context.ShouldBindJSON(&material)

	// validation errors
	if err != nil {
		log.Println("Cannot unmarshal project, error: ", err.Error())
		// generate validation errors response
		response := helpers.GenerateValidationResponse(err)

		context.JSON(http.StatusBadRequest, response)

		return
	}

	// default http status code = 200
	code := http.StatusOK

	// save project & get it's response
	response := services.CreateMaterial(&material, c.constructionJobMaterialsRepository)

	// save contact failed
	if !response.Success {
		// change http status code to 400
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ConstructionPropertiesController) DeleteJobMaterialByIdHandler(context *gin.Context) {
	id := context.Param("id")

	code := http.StatusOK

	response := services.DeleteJobMaterialById(id, c.constructionJobMaterialsRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ConstructionPropertiesController) UpdateJobPropertyByIdHandler(context *gin.Context) {
	id := context.Param("id")

	var jobProperty models.ConstructionJobProperty

	err := context.ShouldBindJSON(&jobProperty)

	// validation errors
	if err != nil {
		log.Println(err.Error())
		response := helpers.GenerateValidationResponse(err)

		context.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	response := services.UpdateJobPropertyById(
		id, jobProperty, c.constructionJobPropertiesRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ConstructionPropertiesController) UpdateJobMaterialByIdHandler(context *gin.Context) {
	id := context.Param("id")

	var jobMaterial models.ConstructionJobMaterial

	err := context.ShouldBindJSON(&jobMaterial)

	// validation errors
	if err != nil {
		response := helpers.GenerateValidationResponse(err)

		context.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	response := services.UpdateJobMaterialById(
		id, jobMaterial, c.constructionJobMaterialsRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ConstructionPropertiesController) GetMaterialsHandler(context *gin.Context) {
	code := http.StatusOK

	response := services.FindJobMaterials(c.constructionJobMaterialsRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}
