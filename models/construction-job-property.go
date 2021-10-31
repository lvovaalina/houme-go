package models

import (
	"github.com/jinzhu/gorm"
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
	JobCode                        string
	Job                            Job `gorm:"foreignKey:JobCode"`
	CompanyName                    string
}
