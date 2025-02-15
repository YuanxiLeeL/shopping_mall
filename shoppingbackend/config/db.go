package config

import (
	"Democratic_shopping_mall/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := Appconfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error reading database config file: %v", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(Appconfig.Database.maxIdleConns)
	sqlDB.SetMaxOpenConns(Appconfig.Database.maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("Fail to configure database,got error: %v", err)
	}

	global.DB = db

}
