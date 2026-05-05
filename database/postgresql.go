package database

import (
	"apart_community/internals/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgres() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Profile{},
		&model.Apartment{},
		&model.Attachment{},
		&model.Role{},
		&model.UserBelongApartment{},
	)

	if err != nil {
		fmt.Println("migration 에러", err)
	}

	return db
}
