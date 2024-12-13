package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Texts holds the text data for AppStrings and SignupLogin
type Texts struct {
	AppStrings  AppStrings  `json:"AppStrings"`
	SignupLogin SignupLogin `json:"SignupLogin"`
}

// AppStrings holds the text data for the main app strings
type AppStrings struct {
	EnjoyListening string `json:"enjoyListening"`
	LoremText      string `json:"loremText"`
	GetStarted     string `json:"getStarted"`
}

// SignupLogin holds the text data for signup and login-related strings
type SignupLogin struct {
	EnjoyListening     string `json:"enjoyListening"`
	SpotifyDescription string `json:"spotifyDescription"`
	Register           string `json:"register"`
	SignIn             string `json:"signIn"`
}

// Global variable to hold the loaded texts data
var texts Texts

// Load text data from the JSON file once, during server startup
func init() {
	// Load the text data from the JSON file
	filepath := "texts.json"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Could not open texts file: %s", err)
	}
	defer file.Close()

	// Decode the JSON file into the texts struct
	if err := json.NewDecoder(file).Decode(&texts); err != nil {
		log.Fatalf("Could not decode texts JSON: %s", err)
	}
}

func main() {
	// Handle the /api/texts endpoint
	http.HandleFunc("/api/texts", getTextsHandler)

	// Start the server and listen on port 8080
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}

// Handler for serving the text data at /api/texts
func getTextsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Send the pre-loaded texts as JSON response
	if err := json.NewEncoder(w).Encode(texts); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
