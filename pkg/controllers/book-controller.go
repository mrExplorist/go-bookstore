package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mrExplorist/go-bookstore/pkg/models"
	"github.com/mrExplorist/go-bookstore/pkg/utils"
)



var NewBook models.Book


// CreateBook creates a book in the database
// r => request object
// w => response object
func GetAllBooks(w http.ResponseWriter , r *http.Request) {
	newBooks := models.GetAllBooks() 
	res,_ := json.Marshal(newBooks) // convert the books array into a JSON string
	w.Header().Set("Content-Type","pkglication/json") // set the response content type as JSON
	w.WriteHeader(http.StatusOK) // set the response status as 200 OK
	w.Write(res) // send the response as JSON string 
}


// Get a single book by its id 




func GetBookById(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r) // get the book id from request params, key is "id"
	bookId := vars["bookId"] 
	ID,err := strconv.ParseInt(bookId,0,0) // convert the string id to int
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID) // fetch the book details from the database by its id
	res,_ := json.Marshal(bookDetails) // convert the book details to JSON
	w.Header().Set("Content-Type","pkglication/json") // set the response content type as JSON
	w.WriteHeader(http.StatusOK) // set the response status as 200 OK
	w.Write(res) // send the response as JSON string

}

func CreateBook(w http.ResponseWriter , r *http.Request){
	CreateBook := &models.Book{}// initialize book variable as a pointer to a new empty Book struct
	utils.ParseBody(r,CreateBook) // parse the request body so that our database can read it

	b := CreateBook.CreateBook() // create book in the database

	res , _ := json.Marshal(b) // convert the book variable to JSON
	w.Header().Set("Content-Type","pkglication/json") // set the response content type as JSON
	w.WriteHeader(http.StatusOK) // set the response status as 200 OK
	w.Write(res) // send the response as JSON string


} 


// DeleteBook deletes a book from the database by its id

func DeleteBook(w http.ResponseWriter , r *http.Request){

	vars := mux.Vars(r) 

	bookId := vars["bookId"]

	ID,err := strconv.ParseInt(bookId,0,0) // convert the string id to int

	if err != nil {
		fmt.Println("Error while parsing")
	}

	// delete book by id
	book := models.DeleteBook(ID)

	// Now Response
	res,_ := json.Marshal(book) // convert the book variable to JSON

	w.Header().Set("Content-Type","pkglication/json") // set the response content type as JSON

	w.WriteHeader(http.StatusOK) // set the response status as 200 OK

	w.Write(res) // send the response as JSON string

}

// UpdateBook updates a book in the database by its id

func UpdateBook(w http.ResponseWriter , r *http.Request){

	updateBook := &models.Book{}// initialize book variable as a pointer to a new empty Book struct
	utils.ParseBody(r,updateBook) // parse the request body so that our datab ase can read it
	
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID,err := strconv.ParseInt(bookId,0,0) // convert the string id to int

	if err != nil {
		fmt.Println("Error while parsing")
	}
 
	// update book by id
	bookDetails, db := models.GetBookById(ID) // fetch the book details from the database by its id


	// update the book details with new values if it is not empty

	if updateBook.Name != "" {  // update the book details with new values if it is not empty 
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" { // update the book details with new values if it is not empty
   bookDetails.Author = updateBook.Author

	}

	if updateBook.Publication != "" { // update the book details with new values if it is not empty

		bookDetails.Publication = updateBook.Publication

	}

	// save the updated book details in the database
	db.Save(&bookDetails)



	// Now Response 
	res,_ := json.Marshal(bookDetails) // convert the book variable to JSON

	w.Header().Set("Content-Type","pkglication/json") // set the response content type as JSON

	w.WriteHeader(http.StatusOK) // set the response status as 200 OK
	w.Write(res) // send the response as JSON string

}