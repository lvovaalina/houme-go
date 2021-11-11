package models

import "gorm.io/gorm"

type ProjectJob struct {
	gorm.Model
	ProjectJobId        int `gorm:"primary_key;autoIncrement"`
	ConstructionWorkers int

	ConstructionCost     float32
	ConstructionDuration float32

	ConstructionDurationInHours int
	ConstructionDurationInDays  int

	ProjectRefer int

	JobId int
	Job   Job `gorm:"foreignKey:JobId;references:JobId"`
}
