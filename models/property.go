package models

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	PropertyId     int    `gorm:"primary_key;autoIncrement"`
	PropertyCode   string `gorm:"unique"`
	PropertyName   string `gorm:"unique"`
	PropertyNamePL string
	PropertyUnit   string
}
