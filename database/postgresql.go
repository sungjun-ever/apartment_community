package database

import (
	"apart_community/config"
	apartmentDomain "apart_community/internals/apartment/domain"
	attachmentDomain "apart_community/internals/attachment/domain"
	userDomain "apart_community/internals/user/domain"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setLogLevel(ginMode string) logger.LogLevel {
	var logLevel logger.LogLevel

	if ginMode == "debug" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Warn
	}

	return logLevel
}

func setConnOption(db *gorm.DB, env *config.Config) {
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(env.MaxIdleConns)
	sqlDB.SetMaxOpenConns(env.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(env.ConnMaxLifetime) * time.Minute)
}

func ConnectToPostgres(env *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		env.DBHost,
		env.DBUser,
		env.DBPass,
		env.DBName,
		env.DBPort,
	)

	logLevel := setLogLevel(env.GinMode)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logLevel),
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&userDomain.User{},
		&userDomain.Profile{},
		&userDomain.UserRole{},
		&userDomain.UserUnitRole{},
		&apartmentDomain.Apartment{},
		&apartmentDomain.Building{},
		&apartmentDomain.Unit{},
		&attachmentDomain.Attachment{},
	)

	if err != nil {
		fmt.Println("migration 에러", err)
	}

	setConnOption(db, env)

	return db
}
