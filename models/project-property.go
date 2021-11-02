package models

import "gorm.io/gorm"

type ProjectProperty struct {
	gorm.Model
	ProjectPropertyId int `gorm:"primary_key;autoIncrement"`
	ProjectId         int
	PropertyId        int
	PropertyCode      string
	PropertyValue     float32
	Project           Project `gorm:"foreignKey:ProjectId"`
}
