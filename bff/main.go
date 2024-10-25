package main

import (
	"log"
	"net/http"

	"github.com/niemet0502/shirabe/bff/handlers"
	"github.com/niemet0502/shirabe/bff/services"

	bookProto "github.com/niemet0502/shirabe/books/book"
	readingProto "github.com/niemet0502/shirabe/books/readingprogress"
	shelfProto "github.com/niemet0502/shirabe/shelves/shelve"
	userProto "github.com/niemet0502/shirabe/users/user"

	goHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	PORT = ":9005"
)

func main() {

	r := mux.NewRouter()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-north-1"),
	})

	if err != nil {
		log.Printf("Failed to connect to aws")
		panic(err)
	}

	// Create an S3 service client
	s3svc := s3.New(sess)

	// books
	conn, err := grpc.NewClient("localhost:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	bookClient := bookProto.NewBookClient(conn)

	bookSvc := services.NewBookService(bookClient, s3svc)

	bookHandler := handlers.NewBookHandler(bookSvc)

	r.HandleFunc("/books/{id:[0-9]+}", bookHandler.GetBook).Methods("GET")
	r.HandleFunc("/books/search", bookHandler.SearchBooks).Methods("GET")
	r.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", bookHandler.UpdateBook).Methods("PUT")
	r.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")

	// shelves
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

	// shelves
	usersConn, usersErr := grpc.NewClient("localhost:9004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if usersErr != nil {
		panic(usersErr)
	}

	defer shelvesConn.Close()

	userClient := userProto.NewUserServiceClient(usersConn)

	userSvc := services.NewUserService(userClient)

	userHandler := handlers.NewUserHandler(userSvc)

	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUser).Methods("GET")

	// bookshelf

	bsClient := shelfProto.NewBookShelfServiceClient(shelvesConn)

	bsSvc := services.NewBoolShelfService(bsClient, shelvesClient, bookClient)

	bsHandler := handlers.NewBookShelfHandler(bsSvc)

	r.HandleFunc("/shelves/{shelfId:[0-9]+}/books", bsHandler.GetBooksByShelf).Methods("GET")
	r.HandleFunc("/shelves/{shelfId:[0-9]+}/books/{bookId:[0-9]+}", bsHandler.AddBookToShelf).Methods("POST")
	r.HandleFunc("/shelves/{shelfId:[0-9]+}/books/{bookId:[0-9]+}", bsHandler.RemoveBookFromShelf).Methods("DELETE")

	// reading progress

	pgClient := readingProto.NewReadingServiceClient(conn)

	pgSvc := services.NewReadingService(pgClient)

	pgHandler := handlers.NewReadingProgressHandler(pgSvc)

	r.HandleFunc("/readingprogress/{bookId:[0-9]+}", pgHandler.Get).Methods("GET")
	r.HandleFunc("/readingprogress/", pgHandler.Create).Methods("POST")

	// CORS Handlers
	ch := goHandlers.CORS(goHandlers.AllowedOrigins([]string{"http://localhost:5173"}))

	// start the server
	err = http.ListenAndServe(PORT, ch(r))

	if err != nil {
		log.Fatal("The server didn't start")
	}
}
