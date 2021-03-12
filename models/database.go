package models

import (
	AuthModels "main/auth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var User *AuthModels.User

func ConnectDB() {
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  "host=127.0.0.1 port=6432 user=postgres dbname=utmgo password=167a214b59 sslmode=disable",
				PreferSimpleProtocol: true,
			}),
		&gorm.Config{},
	)

	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	err = db.AutoMigrate(&Track{}, User)
	if err != nil {
		panic("Не удалось создать миграции")
	}
	DB = db
}
