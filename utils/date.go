package utils

import (
	"fmt"
	"time"
)

// ParseDate converts a date string in YYYY-MM-DD format to time.Time
func ParseDate(dateStr string) (time.Time, error) {
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
	if date.Before(time.Now().AddDate(0, 0, -1)) {
		return time.Time{}, fmt.Errorf("date cannot be in the past")
	}

	return date, nil
}
