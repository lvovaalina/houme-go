package models

import (
	"gorm.io/gorm"
)

type ConstructionJobProperty struct {
	gorm.Model
	ConstructionJobPropertyId      int `gorm:"primary_key;autoIncrement"`
	ConstructionSpeed              float32
	ConstructionCost               float32
	ConstructionFixDurationInHours float32
	MaxWorkers                     int
	OptWorkers                     int
	MinWorkers                     int
	JobID                          string
	Job                            Job `gorm:"references:JobCode"`
	CompanyName                    string
}
