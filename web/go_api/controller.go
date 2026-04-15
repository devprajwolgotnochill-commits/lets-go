package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// middleware, go way of ... ,how to use this

func (b *Book) IsEmpty() bool {
	// Check if Title is empty AND if the Author pointer is nils
	if b.Author == nil {
		return b.Title == ""
	}

	// If the author exists, but names are blank
	return b.Title == "" && b.Author.FirstName == "" && b.Author.LastName == ""
}

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

// adds new book to the database
func createBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding a new book to DB!")

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Invalid data")
	}

	// {}

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	if book.IsEmpty() {
		json.NewEncoder(w).Encode("Invalid data")
		return
	}

	// rand.Seed(time.Now().UnixNano()) //no seeding required

	book.BookID = strconv.Itoa(rand.Intn(100))

	// library is a slice
	// var library []Book

	library = append(library, book)
	fmt.Printf("Library now has %d books!\n", len(library))

	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	// no db //using for loop // id == // we have id // remove the value // change the value
}
