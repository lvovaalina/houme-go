package models

import "github.com/jinzhu/gorm"

type Job struct {
	gorm.Model
	JobId              int    `gorm:"primary_key;autoIncrement"`
	JobName            string `gorm:"unique"`
	StageName          string
	WallMaterial       string
	FinishMaterial     string
	FoundationMaterial string
	RoofingMaterial    string
	InteriorMaterial   string
	Required           bool
	PropertyCode       string
	JobCode            string `gorm:"unique,index"`
}
