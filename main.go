package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"visa-calculator/visa"
)

type PageData struct {
	Result *VisaResult
	Error  string
}

type VisaResult struct {
	VisaType       string
	EntryDate      string
	ExitDate       string
	TotalDays      int
	MaxAllowedDays int
	RemainingDays  int
}

func main() {
	// Load templates
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Handle main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := &PageData{
			Result: nil,
			Error:  "",
		}
		if err := tmpl.Execute(w, data); err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// Handle calculation
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		data := &PageData{}

		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Get form values
		visaType := strings.ToLower(r.FormValue("visaType"))
		entryDate := r.FormValue("entryDate")
		exitDate := r.FormValue("exitDate")

		// Calculate stay
		calc := visa.NewCalculator()
		result, err := calc.CalculateStay(visaType, entryDate, exitDate)

		if err != nil {
			log.Printf("Calculation error: %v", err)
			data.Error = err.Error()
		} else {
			// Create view model
			data.Result = &VisaResult{
				VisaType:       strings.Title(visaType),
				EntryDate:      result.EntryDate.Format("2006-01-02"),
				ExitDate:       result.ExitDate.Format("2006-01-02"),
				TotalDays:      result.TotalDays,
				MaxAllowedDays: result.MaxAllowedDays,
				RemainingDays:  result.RemainingDays,
			}
		}

		if err := tmpl.Execute(w, data); err != nil {
			log.Printf("Template error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	// Start server
	log.Println("Server starting on http://0.0.0.0:3000")
	if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
		log.Fatal(err)
	}
}
