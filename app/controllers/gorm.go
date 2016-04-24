package controllers

import (
	"github.com/boxp/pso2-bot/app/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"log"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("mysql", dbInfoString())
	if err != nil {
		log.Panicf("Failed to connect to database: %v\n", err)
	}

	db.DB()
	db.AutoMigrate(&models.Arks{})
	DB = db
}

func dbInfoString() string {
	s, b := revel.Config.String("db.info")
	if !b {
		log.Panicf("database info not found")
	}

	return s
}
