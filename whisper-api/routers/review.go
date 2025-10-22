package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/AHKAYY007/Whisper-backend/controllers"
)

func RegisterReviewRoutes(router *gin.Engine, db *gorm.DB) {
	reviewGroup := router.Group("/review")
	{
		reviewGroup.POST("", controllers.CreateReview(db))
		reviewGroup.GET("/business/:id", controllers.GetBusinessReviews(db))
	}
}
