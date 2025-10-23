package routers

import (
	"github.com/AHKAYY007/Whisper-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterBusinessRoutes(router *gin.Engine) {
	business := router.Group("/business")
	{
		business.GET("", controllers.GetBusinesses)
		business.GET("/:id", controllers.GetBusinessByID)
		business.POST("", controllers.CreateBusiness)
		business.POST("/:id/upload", controllers.UploadBusinessImage)
	}
}
