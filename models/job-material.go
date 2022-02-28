package models

import (
	"gorm.io/gorm"
)

type ConstructionJobMaterial struct {
	gorm.Model
	ConstructionJobMaterialId int `gorm:"primary_key;autoIncrement"`
	MaterialName              string
	MaterialNamePL            string
	MaterialCost              float32
	JobID                     string
	Job                       Job `gorm:"references:JobCode"`
	CompanyName               string
}
