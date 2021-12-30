package services

import (
	"bitbucket.org/houmeteam/houme-go/dtos"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
)

func FindProjectJobs(repository *repositories.JobRepository, project models.Project) dtos.Response {
	var stairsNumber float32
	for _, prop := range project.ProjectProperties {
		if prop.PropertyCode == "SN" {
			stairsNumber = prop.PropertyValue
		}
	}
	operationResult := repository.FindProjectJobs(
		project.WallMaterial,
		project.FoundationMaterial,
		project.RoofingMaterial,
		project.FinishMaterial,
		stairsNumber > 0)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*[]models.Job)

	return dtos.Response{Success: true, Data: datas}
}

func FindAllJobs(repository *repositories.JobRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*[]models.Job)

	return dtos.Response{Success: true, Data: datas}
}
