package main

import (
	"fmt"
	"log"
	"movieapp/functions"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var key, host string

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key = os.Getenv("API_KEY")
	host = os.Getenv("API_HOST")

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		functions.HomePage(w, r, key, host)
	}).Methods("GET")

	fmt.Println("Server is running on port 3030...")
	http.ListenAndServe(":3030", r)
}
