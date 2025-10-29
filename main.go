package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handleExample(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	w.Write([]byte("HELLO!"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := http.NewServeMux()
	router.HandleFunc("/", handleExample)

	port := os.Getenv("DEV_PORT")
	server := http.Server{
		Addr:    port,
		Handler: router,
	}

	fmt.Println("Server listening on port: {}", port)
	server.ListenAndServe()
}
