package controllers

import (
	"net/http"

	"bitbucket.org/houmeteam/houme-go/forge"
	"bitbucket.org/houmeteam/houme-go/repositories"
	"bitbucket.org/houmeteam/houme-go/services"
	"github.com/gin-gonic/gin"
)

type CommonController struct {
	propertiesRepository *repositories.PropertyRepository
	jobsRepository       *repositories.JobRepository
}

func NewCommonController(propertiesRepository *repositories.PropertyRepository,
	jobsRepository *repositories.JobRepository) *CommonController {
	return &CommonController{
		propertiesRepository: propertiesRepository,
		jobsRepository:       jobsRepository,
	}
}

func (c *CommonController) ForgeGetHandler(context *gin.Context) {
	projects := forge.GetBucketObjects("houme")
	context.JSON(http.StatusOK, projects)
}

func (c *CommonController) GetProperties(context *gin.Context) {
	code := http.StatusOK

	response := services.FindAllProperties(c.propertiesRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *CommonController) GetJobsHandler(context *gin.Context) {
	code := http.StatusOK

	response := services.FindAllJobs(c.jobsRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}
