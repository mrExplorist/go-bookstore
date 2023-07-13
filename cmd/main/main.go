package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrExplorist/go-bookstore/pkg/routes"

	// import the mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	r := mux.NewRouter() // create a new router
	routes.RegisterBookstoreRoutes(r) // register the routes for the bookstore app 
	http.Handle("/", r) // register the router as the handler for the root endpoint
	log.Fatal(http.ListenAndServe(":9010", r)) // start the server on port 9010 and pass in the router 

	fmt.Println("Server started on port 9010")
}