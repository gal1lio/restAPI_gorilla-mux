package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {
	r := mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "Golang pointers", Author: "Author 1", Year: "2010"},
		Book{ID: 2, Title: "Concurrency", Author: "Author 2", Year: "2012"},
		Book{ID: 3, Title: "Goroutines", Author: "Author 3", Year: "2016"},
		Book{ID: 4, Title: "Go four", Author: "Author 4", Year: "2019"},
		Book{ID: 5, Title: "Go five", Author: "Author 5", Year: "2020"},
	)

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", addBook).Methods("POST")
	r.HandleFunc("/books", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Post a book endpoint is called.")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update book endpoint is called.")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove a book endpoint is called.")
}
