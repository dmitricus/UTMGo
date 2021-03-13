package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	MailerModels "main/mailer/models"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  "host=127.0.0.1 port=5432 user=postgres dbname=utmgo password=root sslmode=disable",
				PreferSimpleProtocol: true,
			}),
		&gorm.Config{},
	)

	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	err = db.AutoMigrate(
		//&AuthModels.User{},
		&MailerModels.EmailServers{},
		&MailerModels.Emails{},
		&MailerModels.Bounced{},
		&MailerModels.Domains{},
	)
	if err != nil {
		panic("Не удалось создать миграции")
	}
	DB = db
}
