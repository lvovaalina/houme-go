package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CompanyId   int `gorm:"primary_key;autoIncrement"`
	CompanyName string

	CompanyJobs []CompanyJob `gorm:"foreignKey:CompanyRefer;references:CompanyId;"`

	Foremen []Foreman `gorm:"foreignKey:CompanyRefer;references:CompanyId;"`
}
