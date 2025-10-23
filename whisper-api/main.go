package main
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/AHKAYY007/Whisper-backend/config"
	"github.com/AHKAYY007/Whisper-backend/routers"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		fmt.Println("⚠️  No .env file found, using system environment variables")
	}

	// Connect to database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("Welcome to the Whisper API powered by Applift Labs 🚀")

	if err := os.MkdirAll("uploads/business", os.ModePerm); err != nil {
        log.Fatalf("failed to create uploads directory: %v", err)
    }

	// Initialize Gin router
	router := gin.Default()

	// Health check route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Whisper API",
			"status":  "API is running successfully ✅",
		})
	})

	router.Static("/uploads", "./uploads")

	routers.RegisterBusinessRoutes(router)
	routers.RegisterReviewRoutes(router, db)

	// Get port from env or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("🚀 Server is running on port %s...\n", port)
	router.Run(":" + port)
}
