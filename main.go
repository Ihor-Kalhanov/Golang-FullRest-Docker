package main

import (
	"encoding/json"
	"fmt"
	"github.com/Ihor-Kalhanov/Golang-FullRest-Docker/controllers"
	"github.com/Ihor-Kalhanov/Golang-FullRest-Docker/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

//// User is the data type for user object
//type User struct {
//	ID        string    `json:"id" sql:"id"`
//	Email     string    `json:"email" validate:"required" sql:"email"`
//	Password  string    `json:"password" validate:"required" sql:"password"`
//	Username  string    `json:"username" sql:"username"`
//	TokenHash string    `json:"tokenhash" sql:"tokenhash"`
//	CreatedAt time.Time `json:"createdat" sql:"createdat"`
//	UpdatedAt time.Time `json:"updatedat" sql:"updatedat"`
//}
//
//// schema for user table
//const schema = `
//		create table if not exists users (
//			id varchar(36) not null,
//			email varchar(225) not null unique,
//			username varchar(225),
//			password varchar(225) not null,
//			tokenhash varchar(15) not null,
//			createdat timestamp not null,
//			updatedat timestamp not null,
//			primary key (id)
//		);
//`

func main() {
	HandleRequests()
}

func HandleRequests() {
	router := mux.NewRouter()

	// Get all books
	router.HandleFunc("/books/", GetAllBooks).Methods("GET")

	router.HandleFunc("/book/create/", CreateBook).Methods("POST")

	router.HandleFunc("/register/", controllers.Register).Methods("POST")

	router.HandleFunc("/signin/", controllers.Signin).Methods("POST")

	router.HandleFunc("/home/", controllers.Home)

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
