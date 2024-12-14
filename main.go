package main

import (
	"encoding/json"
	"log"
	"net/http"
	"visa-calculator/visa"
)

type CalculateRequest struct {
	VisaType  string `json:"visaType"`
	EntryDate string `json:"entryDate"`
	ExitDate  string `json:"exitDate"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	calculator := visa.NewCalculator()
	result, err := calculator.CalculateStay(req.VisaType, req.EntryDate, req.ExitDate)
	
	w.Header().Set("Content-Type", "application/json")
	
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Create a response structure that matches what the frontend expects
	response := map[string]interface{}{
		"TotalDays":      result.TotalDays,
		"MaxAllowedDays": result.MaxAllowedDays,
		"RemainingDays":  result.RemainingDays,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	// Serve static files from the templates directory
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/", fs)

	// Handle calculate endpoint
	http.HandleFunc("/calculate", calculateHandler)

	// Start server
	log.Println("Server starting on http://0.0.0.0:3000")
	if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
		log.Fatal(err)
	}
}