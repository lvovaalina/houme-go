package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	JobId              int    `gorm:"primary_key;autoIncrement"`
	JobName            string `gorm:"unique"`
	StageName          string
	SubStageName       string
	WallMaterial       string
	FinishMaterial     string
	FoundationMaterial string
	RoofingMaterial    string
	InteriorMaterial   string
	Required           bool
	PropertyCode       string
	InParallel         bool
	ParallelGroupCode  string
	JobCode            string `gorm:"unique,index"`
}
