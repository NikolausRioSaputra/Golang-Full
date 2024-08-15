package main

import (
	"GoLibraryHttpserver/internal/handler"
	"GoLibraryHttpserver/internal/repository"
	"GoLibraryHttpserver/internal/usecase"
	"fmt"
	"log"
	"net/http"
)

func main() {
	bookRepo := repository.NewBookRepository()
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookHandler := handler.NewBookHandler(bookUsecase)

	routes := http.NewServeMux()
	routes.HandleFunc("/books", bookHandler.StoreNewBook)
	routes.HandleFunc("/books/all", bookHandler.ListBooks)
	routes.HandleFunc("/books/delete", bookHandler.DeleteBook)

	server := http.Server{
		Addr:   ":5000",
		Handler: routes,
	}

	fmt.Printf("Server running on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil{
		log.Fatal("Server failed to start: ", err)
	}
}
