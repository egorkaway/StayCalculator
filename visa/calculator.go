package visa

import (
	"fmt"
	"time"
	"visa-calculator/utils"
)

type StayPeriod struct {
	EntryDate time.Time
	ExitDate  time.Time
	Duration  int
}

type StayResult struct {
	Periods        []StayPeriod
	TotalDays      int
	MaxAllowedDays int
	RemainingDays  int
}

type Calculator struct{}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) CalculateStay(visaType string, periods []Period) (*StayResult, error) {
	if !IsValidVisaType(visaType) {
		return nil, fmt.Errorf("invalid visa type: %s", visaType)
	}

	maxDays := getMaxAllowedDays(visaType)
	if maxDays == 0 {
		return nil, fmt.Errorf("no stay duration defined for visa type: %s", visaType)
	}

	var stayPeriods []StayPeriod
	totalDays := 0

	for _, period := range periods {
		entryDate, err := utils.ParseDate(period.EntryDate)
		if err != nil {
			return nil, fmt.Errorf("invalid entry date: %v", err)
		}

		exitDate, err := utils.ParseDate(period.ExitDate)
		if err != nil {
			return nil, fmt.Errorf("invalid exit date: %v", err)
		}

		if exitDate.Before(entryDate) {
			return nil, fmt.Errorf("exit date cannot be before entry date")
		}

		duration := int(exitDate.Sub(entryDate).Hours() / 24)
		totalDays += duration

		stayPeriods = append(stayPeriods, StayPeriod{
			EntryDate: entryDate,
			ExitDate:  exitDate,
			Duration:  duration,
		})
	}

	remainingDays := maxDays - totalDays

	return &StayResult{
		Periods:        stayPeriods,
		TotalDays:      totalDays,
		MaxAllowedDays: maxDays,
		RemainingDays:  remainingDays,
	}, nil
}

func getMaxAllowedDays(visaType string) int {
	switch visaType {
	case "tourist30":
		return 30
	case "tourist90":
		return 90
	case "tourist120":
		return 120
	case "business":
		return 180
	case "student":
		return 365
	default:
		return 0
	}
}
