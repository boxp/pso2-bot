package services

import (
	"log"
	"time"

	"github.com/boxp/pso2-bot/app/controllers"
	"github.com/boxp/pso2-bot/app/models"
)

func CreateArks(arks models.Arks) {
	tx := controllers.DB.Begin()

	e := tx.Create(&arks).Error
	if e != nil {
		tx.Rollback()
		log.Fatalf("Failed to create arks %v\n", e)
	}

	tx.Commit()
}

func DeleteExpiredArks() {
	tx := controllers.DB.Begin()
	now := time.Now()
	expiredDate := now.AddDate(0, 0, -3)

	e := tx.Where("created_at <= ?", expiredDate).Delete(&models.Arks{}).Error
	if e != nil {
		log.Fatalf("Failed to delete arks %v\n", e)
		tx.Rollback()
	}

	tx.Commit()
}

func SearchArksByShip(ship int) []models.Arks {
	arkses := []models.Arks{}

	controllers.DB.Find(&arkses, "ship = ?", ship)

	return arkses
}
