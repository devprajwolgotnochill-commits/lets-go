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
	w.Write([]byte("Hellno"))
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

// adds new book to the book_database
func createBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding a new book to DB!")

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Invalid book_data")
	}

	// {}

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	if book.IsEmpty() {
		json.NewEncoder(w).Encode("Invalid book_data")
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

// not effective
// only works when provided full json ,remove and re-add
// better way to do ,find the book by its index (i) and modify the fields of library[i]
// updates a book in the db
func updateBook(w http.ResponseWriter, r *http.Request) {
	//using for loop
	// // id == // we have id // remove the value
	// using slice.append
	fmt.Println("Updating !")

	w.Header().Set("Content-Type", "application/json")

	//updating book through book_id ,grabs id for request
	params := mux.Vars(r)
	id := params["id"]

	// looping, id ,remove ,add

	for i, book_data := range library {
		if book_data.BookID == id {

			// Remove the old book
			library = append(library[:i], library[i+1:]...)

			// a variable to hold the NEW data
			var updatedBook Book

			err1 := json.NewDecoder(r.Body).Decode(&updatedBook)

			if err1 != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// ID remains the same as the URL param
			updatedBook.BookID = id

			// Append the updated book
			library = append(library, updatedBook)

			// Send back the updated book (or the whole library)
			json.NewEncoder(w).Encode(library)

			return

		}

	}
	//finished need to learn JSON and Slice technic
}

// deletesBook from the db
func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleted !")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	for i, book_data := range library {
		if book_data.BookID == id {

			// Remove
			library = append(library[:i], library[i+1:]...)

			// Send back the updated data (or the whole library)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(library)

			return

		}

	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Book not found"})
}
