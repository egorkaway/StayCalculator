package visa

// Valid visa types
var validVisaTypes = map[string]bool{
	"tourist":  true,
	"business": true,
	"student":  true,
}

// IsValidVisaType checks if the provided visa type is valid
func IsValidVisaType(visaType string) bool {
	_, exists := validVisaTypes[visaType]
	return exists
}
