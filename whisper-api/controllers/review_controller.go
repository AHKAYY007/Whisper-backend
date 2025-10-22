package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/AHKAYY007/Whisper-backend/models"
)

// CreateReview handles POST /reviews
func CreateReview(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			BusinessID string `json:"business_id"`
			Rating     int    `json:"rating"`
			Text       string `json:"text"`
			Anonymous  bool   `json:"anonymous"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		// Validate UUID format
		if _, err := uuid.Parse(req.BusinessID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid business ID"})
			return
		}

		review := models.Review{
			ID:         uuid.New().String(),
			BusinessID: req.BusinessID,
			Rating:     req.Rating,
			Text:       req.Text,
			Anonymous:  req.Anonymous,
			CreatedAt:  time.Now().UTC(),
		}

		if err := db.Create(&review).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save review"})
			return
		}

		// Recalculate avg_rating and reviews_count
		var stats struct {
			Avg   float64
			Count int64
		}

		db.Model(&models.Review{}).
			Select("AVG(rating) as avg, COUNT(*) as count").
			Where("business_id = ?", req.BusinessID).
			Scan(&stats)

		db.Model(&models.Business{}).
			Where("id = ?", req.BusinessID).
			Updates(map[string]interface{}{
				"avg_rating":    stats.Avg,
				"review_count": stats.Count,
			})

		c.JSON(http.StatusCreated, review)
	}
}

// GetBusinessReviews handles GET /reviews/business/:id
func GetBusinessReviews(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		businessID := c.Param("id")

		var business models.Business
		if err := db.Preload("Reviews").First(&business, "id = ?", businessID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Business not found"})
			return
		}

		var reviews []gin.H
		for _, r := range business.Reviews {
			reviews = append(reviews, gin.H{
				"rating":    r.Rating,
				"text":      r.Text,
				"anonymous": r.Anonymous,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"id":            business.ID,
			"name":          business.Name,
			"avg_rating":    business.AvgRating,
			"reviews_count": business.ReviewCount,
			"reviews":       reviews,
		})
	}
}
