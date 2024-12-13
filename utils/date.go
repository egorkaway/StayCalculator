package utils

import (
	"fmt"
	"time"
)

// ParseDate converts a date string in YYYY-MM-DD format to time.Time
func ParseDate(dateStr string) (time.Time, error) {
	// Handle empty date string
	if dateStr == "" {
		return time.Time{}, fmt.Errorf("date cannot be empty")
	}

	// Parse the date
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format. Please use YYYY-MM-DD")
	}

	// Validate year range
	year := date.Year()
	if year < 2000 || year > 2100 {
		return time.Time{}, fmt.Errorf("year must be between 2000 and 2100")
	}

	// Don't allow dates in the past
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	if date.Before(today) {
		return time.Time{}, fmt.Errorf("date cannot be in the past")
	}

	return date, nil
}
