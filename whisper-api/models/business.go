package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Business struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	Slug         string    `json:"slug" gorm:"uniqueIndex"`
	Category     string    `json:"category"`
	City         string    `json:"city"`
	Address      string    `json:"address"`
	AvgRating    float64   `json:"avg_rating" gorm:"default:0"`
	ReviewCount  int       `json:"review_count" gorm:"default:0"`
	Verified     bool      `json:"verified" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	Reviews      []Review  `json:"reviews,omitempty" gorm:"foreignKey:BusinessID;constraint:OnDelete:CASCADE"`
}

// Generate slug before saving
func (b *Business) BeforeCreate(tx *gorm.DB) (err error) {
	if b.Slug == "" {
		b.Slug = strings.ToLower(strings.ReplaceAll(b.Name, " ", "-"))
	}
	return
}
