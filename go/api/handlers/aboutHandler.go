package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	massage, err := fmt.Fprintf(w, "This is the about page.")
	if err != nil {
		return
	}
	log.Println(massage)
}
