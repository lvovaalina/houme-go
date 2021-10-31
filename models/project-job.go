package models

import "github.com/jinzhu/gorm"

type ProjectJob struct {
	gorm.Model
	ProjectJobId        int `gorm:"primary_key;autoIncrement"`
	ConstructionWorkers int

	ConstructionCost     float32
	ConstructionDuration float32

	ConstructionCostInHours int
	ConstructionCostInDays  int

	JobCode      string
	PropertyCode string

	JobId     int
	ProjectId int
	Project   Project `gorm:"foreignKey:ProjectId"`
	Job       Job     `gorm:"foreignKey:JobId"`
}
