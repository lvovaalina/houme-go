package converters

import "bitbucket.org/houmeteam/houme-go/models"

func ConvertMaterialToProjectMaterial(material models.ConstructionJobMaterial) models.ProjectJobMaterial {
	return models.ProjectJobMaterial{
		ConstructionJobMaterialId: material.ConstructionJobMaterialId,
		ConstructionJobMaterial:   material,
	}
}
