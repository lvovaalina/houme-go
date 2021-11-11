package models

import (
	"gorm.io/gorm"
)

type ConstructionJobMaterial struct {
	gorm.Model
	ConstructionJobMaterialId int `gorm:"primary_key;autoIncrement"`
	MaterialName              string
	MaterialCost              float32
	JobID                     string
	Job                       Job `gorm:"references:JobCode"`
	PropertyID                string
	Property                  Property `gorm:"references:PropertyCode"`
	CompanyName               string
}
