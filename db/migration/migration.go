package main

import (
	"iot-logging/configs"
	"iot-logging/db"
	"iot-logging/models"

	"gorm.io/gorm"
)

func main() {
	conf := configs.GetConfig()

	db := db.NewpostgresDatabase(conf.Database)
	tx := db.Connect().Begin()

	temperatureMigration(tx)
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func temperatureMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&models.Temperature{})
}
