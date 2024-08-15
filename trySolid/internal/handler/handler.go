package handler

import (
	"fmt"
	"try-solid/internal/domain"
	"try-solid/internal/usecase"
)

type LibraryHandler struct {
    libraryUsecase usecase.LibraryUsecase
}

func NewLibraryHandler(uc usecase.LibraryUsecase) *LibraryHandler {
    return &LibraryHandler{libraryUsecase: uc}
}

func (h *LibraryHandler) RegisterUser(user *domain.User) error {
    return h.libraryUsecase.RegisterUser(user)
}

func (h *LibraryHandler) RegisterBook(book *domain.Book) error {
    return h.libraryUsecase.RegisterBook(book)
}

func (h *LibraryHandler) BorrowBook(userID, bookID int) error {
    return h.libraryUsecase.BorrowBook(userID, bookID)
}

func (h *LibraryHandler) ReturnBook(loanID int) error {
    return h.libraryUsecase.ReturnBook(loanID)
}

func (h *LibraryHandler) ListLoans() {
    loans, err := h.libraryUsecase.ListLoans()
    if err != nil {
        fmt.Println("Error listing loans:", err)
        return
    }
    fmt.Println("Loan List:")
    for _, loan := range loans {
        fmt.Printf("Loan ID: %d, UserID: %d, BookID: %d\n", loan.ID, loan.UserID, loan.BookID)
    }
}
