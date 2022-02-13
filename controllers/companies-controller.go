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

type CompanyController struct {
	companyRepository *repositories.CompanyRepository
	jobRepository     *repositories.JobRepository
}

func NewCompanyController(companyRepository *repositories.CompanyRepository,
	jobRepository *repositories.JobRepository) *CompanyController {
	return &CompanyController{
		companyRepository: companyRepository,
		jobRepository:     jobRepository,
	}
}

func (c *CompanyController) AddCompanyHandler(context *gin.Context) {
	// initialization project model
	var company models.Company

	// validate json
	err := context.ShouldBindJSON(&company)

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
	response := services.CreateCompany(c.companyRepository, c.jobRepository, &company)

	// save contact failed
	if !response.Success {
		// change http status code to 400
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *CompanyController) GetCompaniesHandler(context *gin.Context) {
	code := http.StatusOK

	response := services.GetCompanies(*c.companyRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func (c *CompanyController) DeleteCompanyHandler(context *gin.Context) {
	id := context.Param("id")

	code := http.StatusOK

	response := services.DeleteCompany(*c.companyRepository, id)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}
