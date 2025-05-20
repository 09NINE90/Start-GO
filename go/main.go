package main

import (
	"GoStart/go/api"
	"GoStart/go/database"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	portDB := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || portDB == "" {
		fmt.Println("Не заданы обязательные переменные окружения")
		os.Exit(1)
	}

	err := database.ConnectDB(host, portDB, user, password, dbname)
	if err != nil {
		panic(err)
	}
	defer database.CloseDB()
	port := ":7070"
	log.Printf("Starting server at port %v", port)
	if err := http.ListenAndServe(port, api.Router()); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
