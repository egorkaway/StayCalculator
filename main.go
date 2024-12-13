package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"visa-calculator/visa"
)

func getInput(reader *bufio.Reader, prompt string) (string, error) {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				continue // Allow retry on EOF
			}
			return "", err
		}
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput != "" {
			return trimmedInput, nil
		}
		fmt.Println("Input cannot be empty. Please try again.")
	}
}

func getValidVisaType(reader *bufio.Reader) (string, error) {
	for {
		visaType, err := getInput(reader, "Enter visa type (tourist/business/student): ")
		if err != nil {
			return "", err
		}

		visaType = strings.ToLower(visaType)
		if visa.IsValidVisaType(visaType) {
			return visaType, nil
		}

		fmt.Println("Error: Invalid visa type. Please enter tourist, business, or student.")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Visa Stay Calculator")
	fmt.Println("-------------------")

	// Get valid visa type with retry
	visaType, err := getValidVisaType(reader)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	// Get dates with validation
	var entryDate, exitDate string
	for {
		var err error
		entryDate, err = getInput(reader, "Enter entry date (YYYY-MM-DD): ")
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}

		exitDate, err = getInput(reader, "Enter exit date (YYYY-MM-DD): ")
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}

		// Try to calculate stay duration
		calc := visa.NewCalculator()
		result, err := calc.CalculateStay(visaType, entryDate, exitDate)
		if err != nil {
			fmt.Printf("Error: %v\nPlease try again.\n\n", err)
			continue
		}

		// Display results
		fmt.Println("\nStay Duration Results")
		fmt.Println("--------------------")
		fmt.Printf("Visa Type: %s\n", strings.Title(visaType))
		fmt.Printf("Entry Date: %s\n", result.EntryDate.Format("2006-01-02"))
		fmt.Printf("Exit Date: %s\n", result.ExitDate.Format("2006-01-02"))
		fmt.Printf("Total Stay Duration: %d days\n", result.TotalDays)
		fmt.Printf("Maximum Allowed Stay: %d days\n", result.MaxAllowedDays)
		fmt.Printf("Remaining Days: %d days\n", result.RemainingDays)

		if result.RemainingDays < 0 {
			fmt.Println("\nWARNING: You have exceeded your allowed stay duration!")
		} else if result.RemainingDays <= 7 {
			fmt.Println("\nWARNING: Your visa is about to expire soon!")
		}
		break
	}
}
