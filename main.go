package main

import (
	"encoding/json"
	"fmt"
	"github.com/Ihor-Kalhanov/Golang-FullRest-Docker/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func main() {
	HandleRequests()
}

func HandleRequests() {
	router := mux.NewRouter()

	// Get all books
	router.HandleFunc("/books/", GetAllBooks).Methods("GET")

	router.HandleFunc("/book/create/", CreateBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	db := data.SetupDB()

	var b Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStatement := `INSERT INTO books (name, author) VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, b.Name, b.Author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	fmt.Println("Inserting book ... ")

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	db := data.SetupDB()

	fmt.Println("Getting all Books ... ")
	rows, err := db.Query("SELECT * FROM books ")

	data.CheckError(err)

	var books []Book
	for rows.Next() {
		var book Book

		err := rows.Scan(&book.Name, &book.Author)

		data.CheckError(err)
		books = append(books, book)
	}

	//var response = JsonResponse{Type: "success", Data: books, Message: "Good"}
	bookBytes, _ := json.MarshalIndent(books, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(bookBytes)
	defer rows.Close()
	defer db.Close()

}
