package controllers

import (
	"net/http"

	"github.com/AHKAYY007/Whisper-backend/config"
	"github.com/AHKAYY007/Whisper-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

// GET /businesses
func GetBusinesses(c *gin.Context) {
	var businesses []models.Business
	query := config.DB

	// Optional search by ?name= or ?city=
	name := c.Query("name")
	city := c.Query("city")

	if name != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%")
	}
	if city != "" {
		query = query.Where("LOWER(city) LIKE LOWER(?)", "%"+city+"%")
	}

	if err := query.Find(&businesses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, businesses)
}

// GET /businesses/:id
func GetBusinessByID(c *gin.Context) {
	id := c.Param("id")
	var business models.Business

	if err := config.DB.Preload("Reviews").First(&business, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Business not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, business)
}

// POST /businesses
func CreateBusiness(c *gin.Context) {
    var input models.Business

    // Bind JSON from request
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Generate UUID for primary key
    input.ID = uuid.New().String()

    // Auto-slug handled by BeforeCreate in the model
    if err := config.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return minimal response
    c.JSON(http.StatusCreated, gin.H{
        "id":            input.ID,
        "name":          input.Name,
        "avg_rating":    input.AvgRating,
        "reviews_count": input.ReviewCount,
    })
}