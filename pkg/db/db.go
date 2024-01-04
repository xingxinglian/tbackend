package db

import (
	"log"
	"tonx/pkg/config"
	"tonx/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := config.Config.DBDsn
	log.Println("db dsn : ", dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// atuo migrate
	err = DB.AutoMigrate(&models.CollectionMetadata{}, &models.NftItem{}, &models.CollectionTestnetMetadata{}, &models.NftTestnetItem{})
	if err != nil {
		log.Fatalf(" auto migrate : %v", err)
		return
	}
}
