package models

import (
	"time"

	"gorm.io/gorm"
)

type Foreman struct {
	gorm.Model
	Name           string
	Email          string
	Password       []byte
	PasswordString string `gorm:"-"`
	LastLoginDate  time.Time
	Role           string
	CompanyRefer   int
}
