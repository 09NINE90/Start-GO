package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	massage, err := fmt.Fprintf(w, "Welcome to the home page!")
	if err != nil {
		return
	}
	log.Println(massage)
}
