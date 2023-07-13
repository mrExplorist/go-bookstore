package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mrExplorist/go-bookstore/pkg/config"
)




var db *gorm.DB



// Book struct

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}


func init(){
	config.ConnectDB()
	db = config.GetDB() // returns a pointer to a gorm.DB object which represents a pool of database connections 
	db.AutoMigrate(&Book{}) // creates a books table in the database if it doesn't exist already 
}

//& ------------> Database -helper functions 

// CreateBook creates a book in the database
func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // => returns `true` as primary key is blank
	db.Create(&b)
	return b

}

// GetBooks fetches all the books from the database

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// GetBook fetches a single book from the database by its id

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

// Delete deletes a book from the database by its id

func DeleteBook(id int64) Book {
	var book Book
	db.Where("id = ?", id).Delete(&book)
	return book
}