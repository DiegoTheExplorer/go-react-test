package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//go:embed client/build/client
var client embed.FS

func handleExample(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	w.Write([]byte("HELLO!"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientFS, err := fs.Sub(client, "client/build/client")
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()
	router.Handle("/", http.FileServerFS(clientFS))
	router.HandleFunc("/HELLO", handleExample)

	port := os.Getenv("DEV_PORT")
	server := http.Server{
		Addr:    port,
		Handler: router,
	}

	fmt.Println("Server listening on port: {}", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
