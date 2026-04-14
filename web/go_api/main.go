package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	BookID    string  `json:"book_id"` // Added tag so it's lowercase in JSON
	Title     string  `json:"title"`
	Author    *Author `json:"author"` // This points to the Author struct below
	Pages     int     `json:"pages"`
	BookPrice int     `json:"book_price"`
	IsSaved   bool    `json:"is_saved"`
}

// book.Author == nil == panic == crash

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// fake data
var library []Book

// runs automatically before your main() functions
func init() {
	// *author so we need to pass a ref
	// field Author *Author `json:"author"`
	library = []Book{
		{
			BookID:    "12",
			Title:     "The Go Programming Blueprints",
			Author:    &Author{FirstName: "Mat", LastName: "Ryer"},
			Pages:     300,
			BookPrice: 45,
			IsSaved:   true,
		},
		{
			BookID:    "122",
			Title:     "Clean Code",
			Author:    &Author{FirstName: "Robert", LastName: "Martin"},
			Pages:     464,
			BookPrice: 50,
			IsSaved:   false,
		},
		{
			BookID:    "21",
			Title:     "The Pragmatic Programmer",
			Author:    &Author{FirstName: "Andrew", LastName: "Hunt"},
			Pages:     352,
			BookPrice: 42,
			IsSaved:   true,
		},
		{
			BookID:    "99",
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

func main() {

	r := mux.NewRouter()
	//ref of the func
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/book", serveBook).Methods("GET")
	r.HandleFunc("/book/{id}", get1Book).Methods("GET")
	r.HandleFunc("/createbook", createBook).Methods("POST")

	fmt.Println("Serving at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
