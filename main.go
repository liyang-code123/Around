package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("started-service")

	r := mux.NewRouter()
	// OPTION is the CORS: cross origin resource sharing.
	r.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8080", r))

}
