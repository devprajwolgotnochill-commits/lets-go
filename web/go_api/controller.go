package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// hello world served
func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
	w.Write([]byte("<h1>Hellno</h1>"))
}

func serveBook(w http.ResponseWriter, r *http.Request) {

	fmt.Println("serving Books")
	w.Header().Set("Content-Type", "application/json")

	// seeding with fake value to test
	json.NewEncoder(w).Encode(library)

}

func get1Book(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Serving a book")

	// set the header to json only accepts json
	w.Header().Set("Content-Type", "application/json")

	// htt....com/user?name=".."
	prams := mux.Vars(r)

	// turning the str parms in int ,don't need when used name to serve book
	index, _ := strconv.Atoi(prams["id"])

	// Access by index (with a safety check)
	if index >= 0 && index < len(library) {

		json.NewEncoder(w).Encode(library[index])
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Book not found",
		"info":  "Index out of range",
	})
}
