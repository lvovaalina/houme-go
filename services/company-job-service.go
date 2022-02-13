package services

import (
	"bitbucket.org/houmeteam/houme-go/dtos"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
)

func CreateCompanyJob(companyJob models.CompanyJob, repository *repositories.CompanyJobRepository) dtos.Response {
	operationResult := repository.Save(companyJob)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.CompanyJob)

	return dtos.Response{Success: true, Data: data}
}

func DeleteCompanyJob(id string, repository *repositories.CompanyJobRepository) dtos.Response {
	operationResult := repository.Delete(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.([]*models.CompanyJob)

	return dtos.Response{Success: true, Data: datas}
}

func GetCompanyJobs(companyId string, repository *repositories.CompanyJobRepository) dtos.Response {
	operationResult := repository.GetByCompanyId(companyId)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func UpdateCompanyJob(id string, companyJob models.CompanyJob,
	repository *repositories.CompanyJobRepository) dtos.Response {

	existingCompanyJobResponse := repository.GetById(id)

	if existingCompanyJobResponse.Error != nil {
		return dtos.Response{Success: false, Message: existingCompanyJobResponse.Error.Error()}
	}

	existingCompanyJob := existingCompanyJobResponse.Result.(*models.CompanyJob)

	existingCompanyJob.ConstructionSpeed = companyJob.ConstructionSpeed
	existingCompanyJob.ConstructionCost = companyJob.ConstructionCost
	existingCompanyJob.ConstructionFixDurationInHours = companyJob.ConstructionFixDurationInHours
	operationResult := repository.Save(*existingCompanyJob)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}
