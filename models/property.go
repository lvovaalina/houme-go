package models

import "github.com/jinzhu/gorm"

type Property struct {
	gorm.Model
	PropertyId   int    `gorm:"primary_key;autoIncrement"`
	PropertyCode string `gorm:"index"`
	PropertyName string `gorm:"unique"`
	PropertyUnit string
}
