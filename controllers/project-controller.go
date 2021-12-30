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

type ProjectsController struct {
	projectRepository                   *repositories.ProjectRepository
	constructionJobPropertiesRepository *repositories.ConstructionJobPropertyRepository
	projectJobRepository                *repositories.ProjectJobRepository
	projectPropertyRepository           *repositories.ProjectPropertyRepository
	jobsRepository                      *repositories.JobRepository
	constructionJobMaterialsRepository  *repositories.ConstructionJobMaterialRepository
	projectMaterialRepository           *repositories.ProjectMaterialRepository
}

func NewProjectsController(
	projectRepository *repositories.ProjectRepository,
	constructionJobPropertiesRepository *repositories.ConstructionJobPropertyRepository,
	projectJobRepository *repositories.ProjectJobRepository,
	projectPropertyRepository *repositories.ProjectPropertyRepository,
	jobsRepository *repositories.JobRepository,
	constructionJobMaterialsRepository *repositories.ConstructionJobMaterialRepository,
	projectMaterialRepository *repositories.ProjectMaterialRepository) *ProjectsController {
	return &ProjectsController{
		projectRepository:                   projectRepository,
		constructionJobPropertiesRepository: constructionJobPropertiesRepository,
		projectJobRepository:                projectJobRepository,
		projectPropertyRepository:           projectPropertyRepository,
		jobsRepository:                      jobsRepository,
		constructionJobMaterialsRepository:  constructionJobMaterialsRepository,
		projectMaterialRepository:           projectMaterialRepository,
	}
}

func (c *ProjectsController) GetProjectsHandler(context *gin.Context) {
	code := http.StatusOK

	response := services.GetAllProjects(c.projectRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ProjectsController) UpdateProjectByIdHandler(context *gin.Context) {
	id := context.Param("id")

	var project models.Project

	err := context.ShouldBindJSON(&project)

	// validation errors
	if err != nil {
		log.Println("ERROR: ", err.Error())
		response := err.Error()

		context.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	response := services.UpdateProjectById(
		id, &project, c.projectRepository,
		c.constructionJobPropertiesRepository,
		c.projectJobRepository, c.projectPropertyRepository,
		c.jobsRepository, c.constructionJobMaterialsRepository,
		c.projectMaterialRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ProjectsController) UpdateProjectPropertiesByIdHandler(context *gin.Context) {
	id := context.Param("id")

	var project models.Project

	err := context.ShouldBindJSON(&project)

	// validation errors
	if err != nil {
		log.Println("ERROR: ", err.Error())
		response := err.Error()

		context.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	response := services.UpdateProjectProperties(
		id, &project, c.projectRepository,
		c.jobsRepository, c.projectJobRepository,
		c.constructionJobPropertiesRepository,
		c.constructionJobMaterialsRepository, c.projectMaterialRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ProjectsController) CreateProjectHandler(context *gin.Context) {
	// initialization project model
	var project models.Project

	// validate json
	err := context.ShouldBindJSON(&project)

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
	response := services.CreateProject(
		&project, c.projectRepository, c.constructionJobPropertiesRepository,
		c.jobsRepository, c.constructionJobMaterialsRepository)

	// save contact failed
	if !response.Success {
		// change http status code to 400
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ProjectsController) DeleteProjectByIdHandler(context *gin.Context) {
	id := context.Param("id")

	code := http.StatusOK

	response := services.DeleteProjectById(
		id, c.projectRepository, c.projectPropertyRepository, c.projectJobRepository, c.projectMaterialRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ProjectsController) GetProjectJobsByProjectIdHandler(context *gin.Context) {
	projectId := context.Param("projectId")

	code := http.StatusOK

	response := services.FindJobsByProjectId(
		projectId,
		c.projectJobRepository,
		c.projectPropertyRepository,
		c.constructionJobPropertiesRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ProjectsController) UpdateProjectsHandler(context *gin.Context) {
	code := http.StatusOK

	response := services.UpdateProjectsJobs(
		c.projectRepository,
		c.constructionJobPropertiesRepository,
		c.projectJobRepository,
		c.jobsRepository,
		c.constructionJobMaterialsRepository,
		c.projectMaterialRepository)
	if !response.Success {
		// change http status code to 400
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *ProjectsController) GetProjectByIdHandler(context *gin.Context) {
	id := context.Param("id")

	code := http.StatusOK

	response := services.GetProjectById(id, c.projectRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}
