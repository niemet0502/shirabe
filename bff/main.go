package main

import (
	"log"
	"net/http"

	"github.com/niemet0502/shirabe/bff/handlers"
	"github.com/niemet0502/shirabe/bff/services"
	bookProto "github.com/niemet0502/shirabe/books/book"
	shelfProto "github.com/niemet0502/shirabe/shelves/shelve"

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

	// shelves
	// create client for each grpc
	shelvesConn, shelvesErr := grpc.NewClient("localhost:9003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if shelvesErr != nil {
		panic(shelvesErr)
	}

	defer shelvesConn.Close()

	shelvesClient := shelfProto.NewShelveServiceClient(shelvesConn)

	shelvesSvc := services.NewShelvesService(shelvesClient)

	shelvesHandler := handlers.NewShelvesHandler(shelvesSvc)

	r.HandleFunc("/shelves", shelvesHandler.GetShelves).Methods("GET")
	r.HandleFunc("/shelves", shelvesHandler.CreateShelf).Methods("POST")
	r.HandleFunc("/shelves/{id:[0-9]+}", shelvesHandler.RemoveShelf).Methods("DELETE")
	r.HandleFunc("/shelves/{id:[0-9]+}", shelvesHandler.UpdateShelf).Methods("PUT")
	r.HandleFunc("/shelves/{id:[0-9]+}", shelvesHandler.GetShelf).Methods("GET")

	// start the server
	err = http.ListenAndServe(PORT, r)

	if err != nil {
		log.Fatal("The server didn't start")
	}
}
