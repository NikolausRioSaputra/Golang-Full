package usecase

import (
	"try-solid/internal/domain"
	"try-solid/internal/repository"
)

// LibraryUsecase adalah interface untuk use case perpustakaan.
type LibraryUsecase interface {
    RegisterUser(user *domain.User) error
    RegisterBook(book *domain.Book) error
    BorrowBook(userID, bookID int) error
    ReturnBook(loanID int) error
    ListLoans() ([]*domain.Loan, error)
}

// LibraryInteractor adalah implementasi dari LibraryUsecase.
type LibraryInteractor struct {
    userRepository  repository.UserRepository
    bookRepository  repository.BookRepository
    loanRepository  repository.LoanRepository
}

func NewLibraryInteractor(userRepo repository.UserRepository, bookRepo repository.BookRepository, loanRepo repository.LoanRepository) *LibraryInteractor {
    return &LibraryInteractor{
        userRepository: userRepo,
        bookRepository: bookRepo,
        loanRepository: loanRepo,
    }
}

func (uc *LibraryInteractor) RegisterUser(user *domain.User) error {
    return uc.userRepository.Save(user)
}

func (uc *LibraryInteractor) RegisterBook(book *domain.Book) error {
    return uc.bookRepository.Save(book)
}

func (uc *LibraryInteractor) BorrowBook(userID, bookID int) error {
    _, err := uc.userRepository.FindByID(userID)
    if err != nil {
        return err
    }
    _, err = uc.bookRepository.FindByID(bookID)
    if err != nil {
        return err
    }
    loan := &domain.Loan{
        ID:     len(uc.loanRepository.(*repository.InMemoryRepository).library.Loans) + 1,
        UserID: userID,
        BookID: bookID,
    }
    return uc.loanRepository.Save(loan)
}

func (uc *LibraryInteractor) ReturnBook(loanID int) error {
    _, err := uc.loanRepository.FindByID(loanID)
    if err != nil {
        return err
    }
    delete(uc.loanRepository.(*repository.InMemoryRepository).library.Loans, loanID)
    return nil
}

func (uc *LibraryInteractor) ListLoans() ([]*domain.Loan, error) {
    return uc.loanRepository.List(), nil
}
