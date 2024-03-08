package model

import (
	"github.com/jinzhu/gorm"
	"github.com/arnabnath01/Go-bookshop/pkg/config"
)


var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`

}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}


func (b *Book) CreateBook() *Book {
	db.NewRecord(b) 	// gorm method to create a new record
	db.Create(&b) 		//  gorm create a new record
	return b
}

func GetAllBooks() []Book {
	var Books []Book

	db.Find(&Books)
	return Books
}


func GetBookById(Id int) (*Book, *gorm.DB) {
	var getBook Book
	db:=db.Where("ID=?",Id).Find(&getBook) 		// gorm method to find a record
	return &getBook,db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?",Id).Delete(&book)
	return book
	
}

func UpdateBook(b *Book) *Book {
	db.Save(&b)
	return b
}