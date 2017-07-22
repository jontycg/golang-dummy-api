package main

import "github.com/gorilla/mux"
import "log"
import "net/http"
import "encoding/json"
import "fmt"

//Book a struct for Book to be stored in the database
type Book struct {
	ID        string     `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	Author    string     `json:"author,omitempty"`
	Bookstore *BookStore `json:"bookstore,omitempty"`
}

//BookStore a struct for a bookstore which is part of the Book struct
type BookStore struct {
	Shop string `json:"shop, omitempty"`
	City string `json:"city, omitempty"`
}

// This will be the "Database" that will be used for this dummy api
var books []Book

//GetBookEndpoint finds a singular book from it's ID
func GetBookEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//GetBooksEndpoint returns the entire database of books
func GetBooksEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(books)
}

//CreateBookEndpoint adds a new book to the database
func CreateBookEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

//DeleteBookEndpoint deletes a book from the database
func DeleteBookEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for i, book := range books {
		if book.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	router := mux.NewRouter()
	// Declaring the data for the "database" that the dummy api will be using
	books = append(books, Book{ID: "1", Title: "The Da Vinci Code", Author: "Dan Brown", Bookstore: &BookStore{Shop: "Borders", City: "Auckland"}})
	books = append(books, Book{ID: "2", Title: "Angels and Demons", Author: "Dan Brown", Bookstore: &BookStore{Shop: "Borders", City: "Auckland"}})
	router.HandleFunc("/books", GetBooksEndpoint).Methods("GET")
	router.HandleFunc("/book/{id}", GetBookEndpoint).Methods("GET")
	router.HandleFunc("/book/{id}", CreateBookEndpoint).Methods("POST")
	router.HandleFunc("/book/{id}", DeleteBookEndpoint).Methods("DELETE")
	fmt.Println("Server has started on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
