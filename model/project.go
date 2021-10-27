package model

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	ProjectId                int    `gorm:"primary_key;autoIncrement"`
	Name                     string `gorm:"unique"`
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
}
