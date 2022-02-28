package services

import (
	"bitbucket.org/houmeteam/houme-go/dtos"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
)

func FindJobMaterials(repository *repositories.ConstructionJobMaterialRepository) dtos.Response {
	operationResult := repository.FindMaterials()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*[]models.ConstructionJobMaterial)

	return dtos.Response{Success: true, Data: datas}
}

func CreateMaterial(material *models.ConstructionJobMaterial, repository *repositories.ConstructionJobMaterialRepository) dtos.Response {
	operationResult := repository.Save(material)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.ConstructionJobMaterial)

	return dtos.Response{Success: true, Data: data}
}

func FindJobMaterialById(id string, repository *repositories.ConstructionJobMaterialRepository) dtos.Response {
	operationResult := repository.FindJobMaterialById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.ConstructionJobMaterial)

	return dtos.Response{Success: true, Data: data}
}

func UpdateJobMaterialById(id string, jobMaterial models.ConstructionJobMaterial,
	repository *repositories.ConstructionJobMaterialRepository) dtos.Response {

	existingPropertyResponse := FindJobMaterialById(id, repository)

	if !existingPropertyResponse.Success {
		return existingPropertyResponse
	}

	existingJobMaterial := existingPropertyResponse.Data.(*models.ConstructionJobMaterial)

	existingJobMaterial.CompanyName = jobMaterial.CompanyName
	existingJobMaterial.MaterialName = jobMaterial.MaterialName
	existingJobMaterial.MaterialNamePL = jobMaterial.MaterialNamePL
	existingJobMaterial.MaterialCost = jobMaterial.MaterialCost

	operationResult := repository.Save(existingJobMaterial)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteJobMaterialById(id string, repository *repositories.ConstructionJobMaterialRepository) dtos.Response {
	operationResult := repository.DeleteJobMaterialById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
