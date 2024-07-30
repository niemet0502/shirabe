# Shirabe 

## Architecture 

<img src="architecture.png" /> <br />

## Functionnal services 

## Backend for frontend

## Web

## Installation

## Built by 
### Packages 
```
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install google.golang.org/grpc
	go install google.golang.org/grpc/reflection

    go get -u github.com/lib/pq

```

### Generate code

```
    protoc -I protos/ --go_out=. --go-grpc_out=. protos/currency.proto
```
## Features 
- [x] Create a book 
- [x] Search for a book => pass a search params and apply it to (title and author)
- [x] Update book's reading progress 
- [x] Update book status => (reading, to read, finished)
- [x] Fetch currently reading books 
- [x] Fetch next reading 
- [x] Create a shelf
- [x] Remove a Shelf
- [x] Fetch Shelves list 
- [x] Fetch Shelf details
- [x] Add a book to a shelf 
- [x] Remove book from a 
- [x] Get shelf's books 
- [x] Sign in (create user)
- Sign up (login)
- Sign out
- Remove an account 
- Generate book's information details
