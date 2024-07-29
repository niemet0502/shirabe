package main

import (
	"log"
	"net/http"

	"github.com/niemet0502/shirabe/bff/handlers"
	"github.com/niemet0502/shirabe/bff/services"
	bookProto "github.com/niemet0502/shirabe/books/book"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	PORT = ":9005"
)

func main() {

	r := mux.NewRouter()
	// create client for each grpc
	conn, err := grpc.NewClient("localhost:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	bookClient := bookProto.NewBookClient(conn)

	// create each service
	bookSvc := services.NewBookService(bookClient)

	// create handlers
	bookHandler := handlers.NewBookHandler(bookSvc)

	// create routes
	r.HandleFunc("/books/{id:[0-9]+}", bookHandler.GetBook).Methods("GET")
	r.HandleFunc("/books/search", bookHandler.SearchBooks).Methods("GET")
	r.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", bookHandler.UpdateBook).Methods("PUT")
	r.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")

	// start the server
	err = http.ListenAndServe(PORT, r)

	if err != nil {
		log.Fatal("The server didn't start")
	}
}
