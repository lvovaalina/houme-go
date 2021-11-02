package models

import "gorm.io/gorm"

type ProjectProperty struct {
	gorm.Model
	ProjectPropertyId int `gorm:"primary_key;autoIncrement"`
	ProjectID         int
	PropertyId        int
	PropertyCode      string
	PropertyValue     float32
	Property          Property `gorm:"foreignKey:PropertyId;references:PropertyId"`
}
