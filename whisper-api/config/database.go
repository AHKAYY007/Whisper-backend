package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/AHKAYY007/Whisper-backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getDatabaseURL() string {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using default database path")
	}
	dbPath := os.Getenv("DEV_DB_URL")
	if dbPath == "" {
		dbPath = "whisper.db"
	}
	return dbPath
}

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(getDatabaseURL()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("✅ Database connected successfully!")
	log.Println("📦 Running migrations...")

	if err := db.AutoMigrate(&models.Business{}, &models.Review{}); err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	log.Println("🎉 Migrations completed successfully")
	return db, nil
}
