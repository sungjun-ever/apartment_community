package database

import (
	"apart_community/internal/model"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setDbLogLevel() logger.LogLevel {
	ginMode := os.Getenv("GIN_MODE")

	switch ginMode {
	case "release":
		return logger.Silent
	case "test":
		return logger.Error
	case "debug":
		return logger.Info
	}

	return logger.Silent
}

func ConnectMySQLDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(setDbLogLevel()),
	})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Profile{},
		&model.Apartment{},
		&model.Attachment{},
		&model.Board{},
		&model.BoardPermission{},
		&model.Post{},
		&model.PostComment{},
		&model.Role{},
		&model.UserBelongApartment{},
	)

	if err != nil {
		fmt.Println("migration 에러", err)
	}

	return db
}
