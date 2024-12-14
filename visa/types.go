package visa

// Valid visa types
var validVisaTypes = map[string]bool{
	"tourist30":  true,
	"tourist90":  true,
	"tourist120": true,
	"business":   true,
	"student":    true,
}

// Period represents a single stay period with entry and exit dates
type Period struct {
	EntryDate string `json:"entryDate"`
	ExitDate  string `json:"exitDate"`
}
// IsValidVisaType checks if the provided visa type is valid
func IsValidVisaType(visaType string) bool {
	_, exists := validVisaTypes[visaType]
	return exists
}
