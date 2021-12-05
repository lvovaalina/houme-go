package models

import "gorm.io/gorm"

type ProjectJobMaterial struct {
	gorm.Model
	ProjectJobMaterialId int `gorm:"primary_key;autoIncrement"`
	MaterialCost         float32

	ConstructionJobMaterialId int
	ConstructionJobMaterial   ConstructionJobMaterial `gorm:"foreignKey:ConstructionJobMaterialId;references:ConstructionJobMaterialId"`

	ProjectRefer int
}
