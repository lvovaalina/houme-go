package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Email          string
	Password       []byte
	PasswordString string `gorm:"-"`
	LastLoginDate  time.Time
	Role           string
}
