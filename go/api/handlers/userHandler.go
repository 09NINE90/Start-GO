package handlers

import (
	"GoStart/go/models"
	"GoStart/go/services"
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"log"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	response := services.GetUserService(w)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, `{"status":"error","message":"failed to encode json"}`, http.StatusInternalServerError)
	}
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	person, err := models.PersonFromJSON(body)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	log.Printf("Received data: Name=%s, Email=%s, Age=%d",
		person.GetName(), person.GetEmail(), person.GetAge())

	models.NewPerson(person.GetName(), person.GetEmail(), person.GetAge())
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Hello %s, your data has been received", person.GetName()),
	})
}
