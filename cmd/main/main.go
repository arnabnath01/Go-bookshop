package main

import (
	
	"log"
	"net/http"
	"github.com/arnabnath01/Go-bookshop/pkg/routes"
	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r:=mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":8080",nil))

}