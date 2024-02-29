package model

import (
	"github.com/jinzhu/gorm"
	"go_bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`

}