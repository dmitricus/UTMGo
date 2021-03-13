package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
	"main/mailer/models"
)

var (
	orm *gorm.DB
	err error
)

func Init(c db.Connection) {
	orm, err = gorm.Open("postgres", c.GetDB("default"))

	if err != nil {
		panic("initialize orm failed")
	}
	orm.AutoMigrate(
		&models.EmailServers{},
	)
}
