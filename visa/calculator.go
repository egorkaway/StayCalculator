package visa

import (
	"fmt"
	"time"
	"visa-calculator/utils"
)

type StayResult struct {
	EntryDate      time.Time
	ExitDate       time.Time
	TotalDays      int
	MaxAllowedDays int
	RemainingDays  int
}

type Calculator struct{}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) CalculateStay(visaType, entryDateStr, exitDateStr string) (*StayResult, error) {
	// Parse dates
	entryDate, err := utils.ParseDate(entryDateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid entry date: %v", err)
	}

	exitDate, err := utils.ParseDate(exitDateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid exit date: %v", err)
	}

	// Validate dates
	if exitDate.Before(entryDate) {
		return nil, fmt.Errorf("exit date cannot be before entry date")
	}

	// Validate visa type
	if !IsValidVisaType(visaType) {
		return nil, fmt.Errorf("invalid visa type: %s", visaType)
	}

	// Get maximum allowed days for visa type
	maxDays := getMaxAllowedDays(visaType)
	
	// Validate if visa type has allowed days
	if maxDays == 0 {
		return nil, fmt.Errorf("no stay duration defined for visa type: %s", visaType)
	}

	// Calculate total stay duration
	totalDays := int(exitDate.Sub(entryDate).Hours() / 24)

	// Calculate remaining days
	remainingDays := maxDays - totalDays

	return &StayResult{
		EntryDate:      entryDate,
		ExitDate:       exitDate,
		TotalDays:      totalDays,
		MaxAllowedDays: maxDays,
		RemainingDays:  remainingDays,
	}, nil
}

func getMaxAllowedDays(visaType string) int {
	switch visaType {
	case "tourist":
		return 90
	case "business":
		return 180
	case "student":
		return 365
	default:
		return 0
	}
}
