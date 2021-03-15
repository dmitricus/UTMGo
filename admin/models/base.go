package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
	"main/mailer/models"
)

var (
	ORM *gorm.DB
	err error
)

func Init(c db.Connection) {
	ORM, err = gorm.Open("postgres", c.GetDB("default"))
	if err != nil {
		panic("initialize orm failed")
	}
	ORM.AutoMigrate(
		&models.EmailServers{},
		&models.Domains{},
		&models.Bounced{},
		&models.Emails{},
	)
}
