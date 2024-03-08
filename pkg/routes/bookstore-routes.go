package routes

import (
	"github.com/gorilla/mux"
	"github.com/arnabnath01/Go-bookshop/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router){
	router.HandleFunc("/books",controllers.Getbooks).Methods("GET")
	router.HandleFunc("/books/{id}",controllers.Getbook).Methods("GET")
	router.HandleFunc("/books/",controllers.Createbooks).Methods("POST")
	router.HandleFunc("/books/{id}",controllers.Updatebooks).Methods("PUT")
	router.HandleFunc("/books/{id}",controllers.Deletebooks).Methods("DELETE")
}


