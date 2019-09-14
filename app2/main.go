package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe(":8082", router))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("received a call from %s\n", r.UserAgent())
	fmt.Fprintf(w, "welcome from app2")
}
