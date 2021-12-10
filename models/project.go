package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ProjectId                 int    `gorm:"unique;primary_key;autoIncrement"`
	Name                      string `gorm:"unique" binding:"required"`
	BucketName                string
	Filename                  string `gorm:"unique"`
	LivingArea                string
	Margin                    string
	Workers                   string
	ProjectCoverBase64        string `gorm:"-"`
	ProjectCover              []byte
	ConstructionDuration      int
	ConstructionCost          int
	ConstructionMaterialCost  int
	ConstructionJobCost       int
	ConstructionWorkersNumber string
	FoundationMaterial        string
	WallMaterial              string
	FinishMaterial            string
	RoofingMaterial           string
	ConstructionCompanyName   string

	ProjectJobs       []ProjectJob         `gorm:"foreignKey:ProjectRefer;references:ProjectId;"`
	ProjectProperties []ProjectProperty    `gorm:"foreignKey:ProjectRefer;references:ProjectId;"`
	ProjectMaterials  []ProjectJobMaterial `gorm:"foreignKey:ProjectRefer;references:ProjectId;"`
}

type ProjectMin struct {
	ProjectId                int
	Name                     string
	BucketName               string
	Filename                 string
	LivingArea               string
	Margin                   string
	Workers                  string
	RoomsNumber              int
	ConstructionCost         int
	ConstructionMaterialCost int
	ConstructionJobCost      int
	ConstructionDuration     int
	ProjectCover             []byte
	ProjectCoverBase64       string `gorm:"-"`
}
