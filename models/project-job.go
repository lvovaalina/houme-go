package models

import "github.com/jinzhu/gorm"

type ProjectJob struct {
	gorm.Model
	ProjectJobId        int `gorm:"primary_key;autoIncrement"`
	ConstructionWorkers int

	ConstructionCost     float32
	ConstructionDuration float32

	ConstructionJobPropertyId int
	JobId                     int
	ProjectId                 int
	ConstructionJobProperty   ConstructionJobProperty `gorm:"foreignKey:ConstructionJobPropertyId"`
	Project                   Project                 `gorm:"foreignKey:ProjectId"`
	Job                       Job                     `gorm:"foreignKey:JobId"`
}
