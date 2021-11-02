package services

import (
	"bitbucket.org/houmeteam/houme-go/dtos"
	"bitbucket.org/houmeteam/houme-go/helpers"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
)

func CreateProject(
	project *models.Project,
	repository repositories.ProjectRepository,
	constructionJobPropertyRepository repositories.ConstructionJobPropertyRepository) dtos.Response {

	constructionPropertiesRepositoryResult := constructionJobPropertyRepository.
		FindPropertiesByCompanyName(project.ConstructionCompanyName)

	if constructionPropertiesRepositoryResult.Error != nil {
		return dtos.Response{Success: false, Message: constructionPropertiesRepositoryResult.Error.Error()}
	}

	projectJobsCalculated := helpers.CalculateCostDurationForProjectJobs(
		*project,
		constructionPropertiesRepositoryResult.Result.(*[]models.ConstructionJobProperty))

	project.ProjectJobs = projectJobsCalculated
	operationResult := repository.Save(project)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Project)

	return dtos.Response{Success: true, Data: data}
}

func GetAllProjects(repository repositories.ProjectRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*[]models.ProjectMin)

	return dtos.Response{Success: true, Data: datas}
}

func DeleteProjectById(id string, repository repositories.ProjectRepository) dtos.Response {
	operationResult := repository.DeleteProjectById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
