package models

import (
	"time"
)

type Review struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	BusinessID string    `json:"business_id" gorm:"index;not null"`
	Rating     int       `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Text       string    `json:"text"`
	Anonymous  bool      `json:"anonymous" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}
