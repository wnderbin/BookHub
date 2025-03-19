package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"main/handlers"

	"github.com/gorilla/mux"
)

var (
	PORT string = ":8080"
)

func main() {
	if len(os.Args) == 2 {
		PORT = ":" + os.Args[1]
	}
	log.Printf("[PORT] %s\n", PORT)

	router := mux.NewRouter()
	router.HandleFunc("/main", handlers.MainPageHandler).Methods("GET")    // main page handler
	router.HandleFunc("/main/", handlers.MainPageHandler).Methods("GET")   // main page handler
	router.HandleFunc("/books/", handlers.BooksHandler).Methods("GET")     // books page handler
	router.HandleFunc("/books/{category}", handlers.CategoryBooksHandler)  // category books page handlers
	router.HandleFunc("/community/", handlers.UsersHandler).Methods("GET") // users page handle

	router.HandleFunc("/user/register/", handlers.RegisterNewUserHandler).Methods("GET") // registration page handler
	router.HandleFunc("/books/add/", handlers.AddNewBook).Methods("GET")                 // add book page handler
	router.HandleFunc("/books/update/", handlers.UpdateBookFormHandler).Methods("GET")   // update book page handler
	router.HandleFunc("/user/update/", handlers.UpdateUserFormHandler).Methods("GET")    // update user page handler
	router.HandleFunc("/books/delete/", handlers.DeleteBookHandler).Methods("GET")       // delete book handler
	router.HandleFunc("/user/delete/", handlers.DeleteUserHandler).Methods("GET")        // delete user handler

	router.HandleFunc("/user/register/postform", handlers.PostformRegisterHandler).Methods("POST")  // postform registration page handler
	router.HandleFunc("/books/add/postform", handlers.PostformAddBookHandler).Methods("POST")       // postform add book page handler
	router.HandleFunc("/books/update/postform", handlers.PostformUpdateBookHandler).Methods("POST") // postform update book handler
	router.HandleFunc("/user/update/postform", handlers.PostformUpdateUserHandler).Methods("POST")  // postform update user handler
	router.HandleFunc("/books/delete/postform", handlers.PostformDeleteBook).Methods("POST")        // postform delete book handler
	router.HandleFunc("/user/delete/postform", handlers.PostformDeleteUser).Methods("POST")         // postform delete user handler

	router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler) // NotFound handler

	http.Handle("/", router)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join("..", "..", "ui", "static")))))

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
