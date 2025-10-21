package utils

import (
	"strings"
)

// Slugify converts a business name to a URL-friendly slug
func Slugify(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}

// CalculateAverage computes a new average rating
func CalculateAverage(currentAvg float64, totalReviews int, newRating int) float64 {
	total := currentAvg*float64(totalReviews) + float64(newRating)
	return total / float64(totalReviews+1)
}
