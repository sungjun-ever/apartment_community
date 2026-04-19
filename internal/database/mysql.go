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

	db.Migrator().CreateIndex(&model.User{}, "idx_user_uuid")
	db.Migrator().CreateIndex(&model.User{}, "idx_user_email")

	db.Migrator().CreateIndex(&model.Apartment{}, "idx_apart_uuid")

	db.Migrator().CreateIndex(&model.UserBelongApartment{}, "idx_upa_userId")
	db.Migrator().CreateIndex(&model.UserBelongApartment{}, "idx_upa_apartmentId")

	db.Migrator().CreateIndex(&model.Profile{}, "idx_profile_userId")
	db.Migrator().CreateIndex(&model.Profile{}, "idx_profile_imageId")

	db.Migrator().CreateIndex(&model.Post{}, "idx_post_uuid")
	db.Migrator().CreateIndex(&model.Post{}, "idx_post_boardId")
	db.Migrator().CreateIndex(&model.Post{}, "idx_post_apartmentId")
	db.Migrator().CreateIndex(&model.Post{}, "idx_post_userId")

	db.Migrator().CreateIndex(&model.PostComment{}, "idx_comment_postId")
	db.Migrator().CreateIndex(&model.PostComment{}, "idx_comment_userId")

	db.Migrator().CreateIndex(&model.Board{}, "idx_board_apartmentId")
	return db
}
