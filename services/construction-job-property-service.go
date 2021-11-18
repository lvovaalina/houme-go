package services

import (
	"bitbucket.org/houmeteam/houme-go/dtos"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
)

func FindJobProperties(repository repositories.ConstructionJobPropertyRepository) dtos.Response {
	operationResult := repository.FindPropertiesByCompanyName("Construction")

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*[]models.ConstructionJobProperty)

	return dtos.Response{Success: true, Data: datas}
}

func FindJobPropertyById(id string, repository repositories.ConstructionJobPropertyRepository) dtos.Response {
	operationResult := repository.FindJobPropertyById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.ConstructionJobProperty)

	return dtos.Response{Success: true, Data: data}
}

func UpdateJobPropertyById(id string, jobProperty models.ConstructionJobProperty,
	repository repositories.ConstructionJobPropertyRepository,
	projectRepository repositories.ProjectRepository,
	projectJobRepository repositories.ProjectJobRepository,
	projectPropertyRepository repositories.ProjectPropertyRepository) dtos.Response {

	existingPropertyResponse := FindJobPropertyById(id, repository)

	if !existingPropertyResponse.Success {
		return existingPropertyResponse
	}

	existingJobProperty := existingPropertyResponse.Data.(*models.ConstructionJobProperty)

	existingJobProperty.CompanyName = jobProperty.CompanyName
	existingJobProperty.ConstructionCost = jobProperty.ConstructionCost
	existingJobProperty.ConstructionFixDurationInHours = jobProperty.ConstructionFixDurationInHours
	existingJobProperty.ConstructionSpeed = jobProperty.ConstructionSpeed
	existingJobProperty.MaxWorkers = jobProperty.MaxWorkers
	existingJobProperty.MinWorkers = jobProperty.MinWorkers
	existingJobProperty.OptWorkers = jobProperty.OptWorkers

	operationResult := repository.Save(existingJobProperty)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	go UpdateProjectsJobs(projectRepository, repository, projectJobRepository, projectPropertyRepository)

	return dtos.Response{Success: true, Data: operationResult.Result}
}
