package models

import "github.com/jinzhu/gorm"

type ProjectProperty struct {
	gorm.Model
	ProjectPropertyId int `gorm:"primary_key;autoIncrement"`
	ProjectId         int
	PropertyValue     float32
	Project           Project `gorm:"foreignKey:ProjectId"`
}
