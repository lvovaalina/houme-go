package services

import (
	"log"

	"bitbucket.org/houmeteam/houme-go/dtos"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
)

type ProjectJobValue struct {
	JobCode  string
	JobSpeed float32
	JobValue float32
}

type ProjectJobsVM struct {
	ProjectJobs      *[]models.ProjectJob
	ProjectJobValues *[]ProjectJobValue
}

func FindJobsByProjectId(
	projectId string,
	repository repositories.ProjectJobRepository,
	projectPropertiesRepository repositories.ProjectPropertyRepository,
	constructionJobPropertyRepository repositories.ConstructionJobPropertyRepository) dtos.Response {
	operationResult := repository.FindProjectJobsByProjectId(projectId)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var projectJobs = operationResult.Result.(*[]models.ProjectJob)
	log.Println(len(*projectJobs))

	constructionPropertiesRepositoryResult := constructionJobPropertyRepository.
		FindPropertiesByCompanyName("Construction")

	if constructionPropertiesRepositoryResult.Error != nil {
		return dtos.Response{Success: false, Message: constructionPropertiesRepositoryResult.Error.Error()}
	}

	constructionProperties := *constructionPropertiesRepositoryResult.Result.(*[]models.ConstructionJobProperty)
	jobPropMap := map[string]float32{}
	for _, prop := range constructionProperties {
		jobPropMap[prop.JobID] = prop.ConstructionSpeed
	}

	projectPropertiesRepositoryResult := projectPropertiesRepository.FindProjectPropertiesByProjectId(projectId)
	if projectPropertiesRepositoryResult.Error != nil {
		return dtos.Response{Success: false, Message: projectPropertiesRepositoryResult.Error.Error()}
	}

	projectProperties := *projectPropertiesRepositoryResult.Result.(*[]models.ProjectProperty)
	log.Println("PRPOS len", len(projectProperties))
	propertiesMap := map[string]float32{}
	for _, p := range projectProperties {
		propertiesMap[p.PropertyCode] = p.PropertyValue
	}

	projectJobValues := []ProjectJobValue{}
	for _, job := range *projectJobs {
		value := propertiesMap[job.Job.Property.PropertyCode]
		speed := jobPropMap[job.Job.JobCode]
		prop := &ProjectJobValue{JobValue: value, JobSpeed: speed, JobCode: job.Job.JobCode}
		projectJobValues = append(projectJobValues, *prop)
	}

	result := &ProjectJobsVM{ProjectJobs: projectJobs, ProjectJobValues: &projectJobValues}
	log.Println("LEN ", len(*result.ProjectJobs), " ", len(*result.ProjectJobValues))

	return dtos.Response{Success: true, Data: result}
}
