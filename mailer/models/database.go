package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ORM *gorm.DB

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

	ORM = db
}
