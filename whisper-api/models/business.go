package models

import (
	"strings"
	"time"
	"encoding/json"
	"database/sql/driver"

	"gorm.io/gorm"
)


// ContactInfo represents nested contact information for a business.
type ContactInfo struct {
	WebsiteURL string `json:"website_url"`
	WhatsAppNo string `json:"whatsapp_no"`
	Facebook   string `json:"facebook"`
	TikTok     string `json:"tiktok"`
	Instagram  string `json:"instagram"`
}

// Implement the sql.Scanner interface for ContactInfo
func (c *ContactInfo) Scan(value interface{}) error {
	if value == nil {
		*c = ContactInfo{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, c)
}

// Implement the driver.Valuer interface for ContactInfo
func (c ContactInfo) Value() (driver.Value, error) {
	return json.Marshal(c)
}



type Business struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" binding:"required" gorm:"not null"`
	Slug         string    `json:"slug" gorm:"uniqueIndex"`
	Category     string    `json:"category" binding:"required"`
	City         string    `json:"city" binding:"required"`
	Address      string    `json:"address" binding:"required"`
	ImageURL     string    `json:"image_url" gorm:"type:text"`
	AvgRating    float64   `json:"avg_rating" gorm:"default:0"`
	ReviewCount  int       `json:"review_count" gorm:"default:0"`
	Verified     bool      `json:"verified" gorm:"default:false"`
	ContactInfo  ContactInfo  `json:"contact_info" gorm:"type:jsonb"`
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
