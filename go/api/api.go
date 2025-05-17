package api

import (
	"GoStart/go/database"
	"GoStart/go/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func GetFormHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	users, err := database.GetAllUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"error","message":"%v"}`, err), http.StatusInternalServerError)
		return
	}

	// Формируем ответ
	response := map[string]interface{}{
		"status": "success",
		"users":  users,
	}

	// Кодируем и отправляем JSON
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, `{"status":"error","message":"failed to encode json"}`, http.StatusInternalServerError)
	}
}

func PostFormHandler(w http.ResponseWriter, r *http.Request) {
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
