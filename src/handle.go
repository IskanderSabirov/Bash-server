package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	writingAnswerError string = "Writing answer error, all data added to system"
)

func executeCommandsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("Correct decode of json to Request struct %v\n", request)

	answers := executeCommands(request.Commands)

	jsonData, err := json.Marshal(answers)
	if err != nil {
		http.Error(w, "Server error in encoding answers to json", http.StatusInternalServerError)
		return
	}

	log.Printf("Correct encoding answer list to json\n")

	w.Header().Set("Content-Type", "application/json")

	if _, err = w.Write(jsonData); err != nil {
		http.Error(w, writingAnswerError, http.StatusInternalServerError)
		return
	}

}
