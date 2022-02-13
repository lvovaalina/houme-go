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

type CompanyJobController struct {
	companyJobRepository *repositories.CompanyJobRepository
}

func NewCompanyJobController(companyJobRepository *repositories.CompanyJobRepository) *CompanyJobController {
	return &CompanyJobController{
		companyJobRepository: companyJobRepository,
	}
}

func (c *CompanyJobController) GetCompanyJobsHandler(context *gin.Context) {
	id := context.Param("companyId")
	code := http.StatusOK

	response := services.GetCompanyJobs(id, c.companyJobRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *CompanyJobController) DeleteCompanyJobHandler(context *gin.Context) {
	id := context.Param("id")

	code := http.StatusOK

	response := services.DeleteCompanyJob(id, c.companyJobRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *CompanyJobController) CreateCompanyJobHandler(context *gin.Context) {
	var companyJob models.CompanyJob

	// validate json
	err := context.ShouldBindJSON(&companyJob)

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
	response := services.CreateCompanyJob(companyJob, c.companyJobRepository)

	// save contact failed
	if !response.Success {
		// change http status code to 400
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *CompanyJobController) UpdateCompanyJobHandler(context *gin.Context) {
	id := context.Param("id")

	var companyJob models.CompanyJob

	err := context.ShouldBindJSON(&companyJob)

	// validation errors
	if err != nil {
		log.Println(err.Error())
		response := helpers.GenerateValidationResponse(err)

		context.JSON(http.StatusBadRequest, response)

		return
	}

	code := http.StatusOK

	response := services.UpdateCompanyJob(
		id, companyJob, c.companyJobRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}
