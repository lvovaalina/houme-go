package services

import (
	"log"

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

	calcJobMap := map[string]helpers.JobCalculations{}
	for _, prop := range projectJobsCalculated {
		calcJobMap[prop.JobCode] = prop
	}

	for index, job := range project.ProjectJobs {
		calJob := calcJobMap[job.JobCode]
		(&project.ProjectJobs[index]).ConstructionCost = calJob.ConstructionCost
		(&project.ProjectJobs[index]).ConstructionDurationInDays = calJob.ConstructionDurationInDays
		(&project.ProjectJobs[index]).ConstructionDurationInHours = calJob.ConstructionDurationInHours
		(&project.ProjectJobs[index]).ConstructionWorkers = calJob.ConstructionWorkers
		(&project.ProjectJobs[index]).ConstructionDuration = calJob.ConstructionDuration
	}

	var projectCost int
	var projectDuration int
	for _, j := range projectJobsCalculated {
		projectDuration += j.ConstructionDurationInDays
		projectCost += int(j.ConstructionCost)
	}

	project.ConstructionDuration = projectDuration
	project.ConstructionCost = projectCost

	operationResult := repository.Save(project)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Project)

	return dtos.Response{Success: true, Data: data}
}

func UpdateProjectById(
	id string,
	project *models.Project,
	repository repositories.ProjectRepository,
	constructionJobPropertyRepository repositories.ConstructionJobPropertyRepository,
	projectJobRepository repositories.ProjectJobRepository) dtos.Response {

	log.Println("Start update project with id: ", id)
	existingProjectResponse := GetProjectById(id, repository)

	if !existingProjectResponse.Success {
		return existingProjectResponse
	}
	log.Println("Succsessfully retrieved project with id: ", id)
	existingProject := existingProjectResponse.Data.(*models.Project)
	log.Println(existingProject)

	existingProject.Name = project.Name
	existingProject.BucketName = project.BucketName
	existingProject.LivingArea = project.LivingArea
	existingProject.ConstructionCompanyName = project.ConstructionCompanyName
	existingProject.ConstructionWorkersNumber = project.ConstructionWorkersNumber
	existingProject.WallMaterial = project.WallMaterial
	existingProject.FoundationMaterial = project.FoundationMaterial
	existingProject.FinishMaterial = project.FinishMaterial
	existingProject.RoofingMaterial = project.RoofingMaterial

	log.Println("Succsessfully updated properties for project with id: ", id)

	projectJobDeleteResult := projectJobRepository.DeleteProjectJobsByProjectId(existingProject.ProjectId)
	if projectJobDeleteResult.Error != nil {
		log.Println("Failed to remove properties for project with id: ", id)
		return dtos.Response{Success: false, Message: projectJobDeleteResult.Error.Error()}
	}

	log.Println("Succsessfully removed project jobs for project with id: ", id)

	constructionPropertiesRepositoryResult := constructionJobPropertyRepository.
		FindPropertiesByCompanyName(project.ConstructionCompanyName)

	if constructionPropertiesRepositoryResult.Error != nil {
		return dtos.Response{Success: false, Message: constructionPropertiesRepositoryResult.Error.Error()}
	}

	projectJobsCalculated := helpers.CalculateCostDurationForProjectJobs(
		*project,
		constructionPropertiesRepositoryResult.Result.(*[]models.ConstructionJobProperty))

	calcJobMap := map[string]helpers.JobCalculations{}
	for _, prop := range projectJobsCalculated {
		calcJobMap[prop.JobCode] = prop
	}

	for index, job := range project.ProjectJobs {
		calJob := calcJobMap[job.JobCode]
		(&project.ProjectJobs[index]).ConstructionCost = calJob.ConstructionCost
		(&project.ProjectJobs[index]).ConstructionDurationInDays = calJob.ConstructionDurationInDays
		(&project.ProjectJobs[index]).ConstructionDurationInHours = calJob.ConstructionDurationInHours
		(&project.ProjectJobs[index]).ConstructionWorkers = calJob.ConstructionWorkers
		(&project.ProjectJobs[index]).ConstructionDuration = calJob.ConstructionDuration
	}

	var projectCost int
	var projectDuration int
	for _, j := range projectJobsCalculated {
		projectDuration += j.ConstructionDurationInDays
		projectCost += int(j.ConstructionCost)
	}

	existingProject.ConstructionDuration = projectDuration
	existingProject.ConstructionCost = projectCost
	existingProject.ProjectJobs = project.ProjectJobs
	existingProject.ProjectProperties = project.ProjectProperties

	operationResult := repository.Save(existingProject)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
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

func GetProjectById(id string, repository repositories.ProjectRepository) dtos.Response {
	operationResult := repository.GetProjectById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}
