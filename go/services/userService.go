package services

import (
	"GoStart/go/database"
	"fmt"
	"net/http"
)

func GetUserService(w http.ResponseWriter) map[string]interface{} {
	users, err := database.GetAllUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"error","message":"%v"}`, err), http.StatusInternalServerError)
		return nil
	}

	return map[string]interface{}{
		"status": "success",
		"users":  users,
	}
}
