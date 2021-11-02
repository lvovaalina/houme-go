package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ProjectId                int    `gorm:"primary_key;autoIncrement"`
	Name                     string `gorm:"unique" binding:"required"`
	BucketName               string
	Filename                 string `gorm:"unique"`
	LivingArea               string
	RoomsNumber              int
	ConstructionDuraton      int
	ConstructionCost         int
	ConstructonWorkersNumber string
	FoundationMaterial       string
	WallMaterial             string
	FinishMaterial           string
	RoofingMaterial          string
	ConstructionCompanyName  string
	ProjectProperties        []ProjectProperty `gorm:"foreignKey:ProjectId"`
	ProjectJobs              []ProjectJob      `gorm:"foreignKey:ProjectJobId"`
}

type ProjectMin struct {
	ProjectId   int
	Name        string
	BucketName  string
	Filename    string
	LivingArea  string
	RoomsNumber int
}
