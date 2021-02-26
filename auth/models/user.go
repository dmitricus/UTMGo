package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Password     string
	Email        string
	FirstName    string
	LastName     string
	IsSuperusers bool
	IsStaff      bool
	IsActive     bool
}
