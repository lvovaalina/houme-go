package services

import (
	"bitbucket.org/houmeteam/houme-go/dtos"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
)

func FindJobsByProjectId(projectId string, repository repositories.ProjectJobRepository) dtos.Response {
	operationResult := repository.FindProjectJobsByProjectId(projectId)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*[]models.ProjectJob)

	return dtos.Response{Success: true, Data: datas}
}
