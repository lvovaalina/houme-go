package models

import "gorm.io/gorm"

type CompanyJob struct {
	gorm.Model
	CompanyJobId                   int `gorm:"primary_key;autoIncrement"`
	JobName                        string
	StageName                      string
	SubStageName                   string
	WallMaterial                   string
	FinishMaterial                 string
	FoundationMaterial             string
	RoofingMaterial                string
	InteriorMaterial               string
	Required                       bool
	InParallel                     bool
	ParallelGroupCode              string
	JobCode                        string
	JobId                          int
	ConstructionSpeed              float32
	ConstructionCost               float32
	ConstructionFixDurationInHours float32

	CompanyRefer int
}
