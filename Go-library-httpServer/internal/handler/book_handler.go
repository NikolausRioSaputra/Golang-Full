package handler

import (
	"GoLibraryHttpserver/internal/domain"
	"GoLibraryHttpserver/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

type BookHendlerInterface interface {
	StoreNewBook(w http.ResponseWriter, r *http.Request)
	ListBooks(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}

type BookHandler struct {
	BookUsecase usecase.BookUsecaseInterface
}

func NewBookHandler(bookUsecase usecase.BookUsecaseInterface) BookHendlerInterface {
	return BookHandler{
		BookUsecase: bookUsecase,
	}
}

// ListBooks implements BookHendlerInterface.
func (h BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	

	books, err := h.BookUsecase.ListBooks()
	if err != nil {
		log.Printf("Invalid request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
	log.Printf("Books retrieved: %v", books)

}

func (h BookHandler) StoreNewBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Info().Int("http.status.code", http.StatusMethodNotAllowed).Msg("invalid method")
		return
	}

	var book domain.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		log.Info().Int("http.status.code", http.StatusBadRequest).Msg("invalid method")
		return
	}

	err = h.BookUsecase.AddBook(book)
	if err != nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		log.Info().Int("http.status.code", http.StatusInternalServerError).Msg("invalid method")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
	log.Printf("Book added: %v", book)

}

func (h BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Invalid request body", http.StatusInternalServerError)

		log.Info().
			Int("Http.status.code", http.StatusInternalServerError).
			Msg("invalid method")

		return
	}

	idStar := r.URL.Query().Get("id")
	if idStar == "" {
		log.Printf("Mising ID parameter")
		http.Error(w, "Mising ID parameter", http.StatusBadRequest)
	}

	bookID, err := strconv.Atoi(idStar)
	if err != nil {
		log.Printf("Invalid Id parameter : %v", err)
		http.Error(w, "Invalid id parameter", http.StatusBadGateway)
		return
	}

	err = h.BookUsecase.DeleteBook(bookID)
	if err != nil {
		log.Printf("Failed to delete book: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(bookID) 
	if err != nil {
		log.Printf("Book with id %d not found", bookID)
	}
	log.Printf("Book with id %d deleted", bookID)

}
