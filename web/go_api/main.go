package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// main struct
type Book struct {
	Title     string  `json:"title"`  //no space in json
	Author    *Author `json:"author"` //points to the func
	Pages     int     `json:"pages"`
	BookPrice int     `json:"book_price"`
	IsSaved   bool
}

// Author struct
type Author struct {
	FirstName string
	LastName  string
}

// fake data
var library []Book

// runs automatically before your main() functions
func init() {
	// *author so we need to pass a ref
	// field Author *Author `json:"author"`
	library = []Book{
		{
			Title:     "The Go Programming Blueprints",
			Author:    &Author{FirstName: "Mat", LastName: "Ryer"},
			Pages:     300,
			BookPrice: 45,
			IsSaved:   true,
		},
		{
			Title:     "Clean Code",
			Author:    &Author{FirstName: "Robert", LastName: "Martin"},
			Pages:     464,
			BookPrice: 50,
			IsSaved:   false,
		},
		{
			Title:     "The Pragmatic Programmer",
			Author:    &Author{FirstName: "Andrew", LastName: "Hunt"},
			Pages:     352,
			BookPrice: 42,
			IsSaved:   true,
		},
		{
			Title:     "Concurrency in Go",
			Author:    &Author{FirstName: "Katherine", LastName: "Cox-Buday"},
			Pages:     240,
			BookPrice: 35,
			IsSaved:   false,
		},
		{
			Title:     "Head First Go",
			Author:    &Author{FirstName: "Jay", LastName: "McGavren"},
			Pages:     500,
			BookPrice: 50,
			IsSaved:   true,
		},
		{
			Title: "",
		},
	}
}

// middleware, go way of ... ,how to use this
func (b *Book) IsEmpty() bool {
	return b.Title == "" && b.BookPrice == 0
}

func main() {

	r := mux.NewRouter()
	//ref of the func
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/book", serveBook).Methods("GET")
	r.HandleFunc("/book/{id}", get1Book).Methods("GET")

	fmt.Println("Serving!")
	log.Fatal(http.ListenAndServe(":8000", r))

}
