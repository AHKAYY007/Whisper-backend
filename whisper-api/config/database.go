package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/AHKAYY007/Whisper-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDatabaseURL() string {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL not set — please add it to your Railway variables or .env file")
	}
	return dsn
}

func ConnectDatabase() (*gorm.DB, error) {
	dsn := getDatabaseURL()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db

	log.Println("✅ Connected to PostgreSQL successfully!")
	log.Println("📦 Running migrations...")

	if err := DB.AutoMigrate(&models.Business{}, &models.Review{}); err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	log.Println("🎉 Migrations completed successfully")
	return DB, nil
}
