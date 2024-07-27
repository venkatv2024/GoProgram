package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server Running")
	router := mux.NewRouter()
	router.HandleFunc("/helloapi", helloapi).Methods("GET")

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, router))
}

func helloapi(w http.ResponseWriter, r *http.Request) {
	writeMessage(w, http.StatusOK, "Hello venkat")
}

func writeJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeMessage(w http.ResponseWriter, status int, message string) {
	writeJson(w, status, map[string]string{
		"message": message,
	})
}
