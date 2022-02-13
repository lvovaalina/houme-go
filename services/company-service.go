package services

import (
	"log"

	"bitbucket.org/houmeteam/houme-go/converters"
	"bitbucket.org/houmeteam/houme-go/dtos"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
	"golang.org/x/crypto/bcrypt"
)

func CreateCompany(
	repository *repositories.CompanyRepository,
	jobRepository *repositories.JobRepository,
	company *models.Company) dtos.Response {
	if len(company.Foremen) > 0 {
		for index, foreman := range company.Foremen {
			log.Println(foreman.Name)
			password, _ := bcrypt.GenerateFromPassword([]byte(foreman.PasswordString), 14)
			foreman.Password = password

			company.Foremen[index] = foreman
		}
	}

	findJobsOperationResult := FindAllJobs(jobRepository)

	if !findJobsOperationResult.Success {
		return dtos.Response{Success: false, Message: findJobsOperationResult.Message}
	}

	jobs := findJobsOperationResult.Data.(*[]models.Job)
	for _, job := range *jobs {
		company.CompanyJobs = append(company.CompanyJobs, converters.ConvertJobToCompanyJob(job))
	}

	operationResult := repository.Save(*company)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Company)

	return dtos.Response{Success: true, Data: data}
}

func AddForeman(repository *repositories.CompanyRepository, company *models.Company) dtos.Response {
	return dtos.Response{}
}

func GetCompanies(repository repositories.CompanyRepository) dtos.Response {
	operationResult := repository.Get()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*[]models.Company)

	return dtos.Response{Success: true, Data: datas}
}

func DeleteCompany(repository repositories.CompanyRepository, companyId string) dtos.Response {
	operationResult := repository.Delete(companyId)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
