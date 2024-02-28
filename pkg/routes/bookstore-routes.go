package routes

import (
	"github.com/gorilla/mux"
	"go_bookstore/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router){
	router.HandleFunc("/books",controllers.Getbooks).Methods("GET")
	router.HandleFunc("/books/{id}",controllers.Getbook).Methods("GET")
	router.HandleFunc("/books/",controllers.Createbooks).Methods("POST")
	router.HandleFunc("/books/{id}",controllers.Getbooks).Methods("PUT")
	router.HandleFunc("/books/{id}",controllers.Getbooks).Methods("DELETE")
}


