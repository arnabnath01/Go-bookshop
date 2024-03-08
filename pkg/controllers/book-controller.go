package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/arnabnath01/Go-bookshop/pkg/model"
	"github.com/arnabnath01/Go-bookshop/pkg/utils"
	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
)


var NewBook model.Book


func Getbooks(w http.ResponseWriter, r *http.Request) {

	newBooks:= model.GetAllBooks()
	res,_ := json.Marshal(newBooks)

	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Getbook(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	bookId:=vars["id"]

	ID,_ := strconv.ParseInt(bookId,0,0)

	// we have to use the blank identifier to ignore the db object
	//  returned by the GetBookById method as we are only interested in the book object
	
	
	bookdetail,err := json.Marshal(ID)
	if err!=nil {
		fmt.Println("Marshall error !")
	}

	w.Header().Set("Content_type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bookdetail)


}

func Createbooks(w http.ResponseWriter, r *http.Request) {
	cb := &model.Book{}
	utils.ParseBody(r,cb)
	b:=cb.CreateBook()
	res,_:=json.Marshal(b)

	w.Header().Set("Content_type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


func Updatebooks(w http.ResponseWriter, r *http.Request) {
	update:=&model.Book{}

	utils.ParseBody(r,update)

	vars:=mux.Vars(r)
	bookId:=vars["bookId"]

	ID,_ := strconv.ParseInt(bookId,0,0)

	

	bookdetails,db := model.GetBookById(int(ID)) 

	if update.Name !=""{
		bookdetails.Name=update.Name
	}
	if update.Author !=""{
		bookdetails.Author=update.Author
	}
	if update.Publication !=""{
		bookdetails.Publication=update.Publication
	}

	db.Save(&bookdetails)
	res,_:=json.Marshal(bookdetails)

	w.Header().Set("Content_type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func Deletebooks(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	bookId:=vars["id"]

	ID,_ := strconv.ParseInt(bookId,0,0)

	// we have to use the blank identifier to ignore the db object
	//  returned by the GetBookById method as we are only interested in the book object
	bookdetails:= model.DeleteBook(ID) 

	bookdetail,err := json.Marshal(bookdetails)
	if err!=nil {
		fmt.Println("Marshall error !")
	}

	w.Header().Set("Content_type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bookdetail)

}


