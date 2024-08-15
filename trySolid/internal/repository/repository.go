package repository

import (
	"errors"
	"try-solid/internal/domain"
)

// UserRepository adalah interface untuk akses data user.
type UserRepository interface {
    Save(user *domain.User) error
    FindByID(id int) (*domain.User, error)
}

// BookRepository adalah interface untuk akses data buku.
type BookRepository interface {
    Save(book *domain.Book) error
    FindByID(id int) (*domain.Book, error)
}

// LoanRepository adalah interface untuk akses data peminjaman buku.
type LoanRepository interface {
    Save(loan *domain.Loan) error
    FindByID(id int) (*domain.Loan, error)
    List() []*domain.Loan
}

// InMemoryRepository adalah implementasi in-memory dari semua repository.
type InMemoryRepository struct {
    library *domain.Library
}

func NewInMemoryRepository(library *domain.Library) *InMemoryRepository {
    return &InMemoryRepository{library: library}
}

// User repository methods
func (r *InMemoryRepository) Save(user *domain.User) error {
    r.library.Users[user.ID] = user
    return nil
}

func (r *InMemoryRepository) FindByID(id int) (*domain.User, error) {
    user, exists := r.library.Users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

// Book repository methods
func (r *InMemoryRepository) Save(book *domain.Book) error {
    r.library.Books[book.ID] = book
    return nil
}

func (r *InMemoryRepository) FindByID(id int) (*domain.Book, error) {
    book, exists := r.library.Books[id]
    if !exists {
        return nil, errors.New("book not found")
    }
    return book, nil
}

// Loan repository methods
func (r *InMemoryRepository) Save(loan *domain.Loan) error {
    r.library.Loans[loan.ID] = loan
    return nil
}

func (r *InMemoryRepository) FindByID(id int) (*domain.Loan, error) {
    loan, exists := r.library.Loans[id]
    if !exists {
        return nil, errors.New("loan not found")
    }
    return loan, nil
}

func (r *InMemoryRepository) List() []*domain.Loan {
    loans := []*domain.Loan{}
    for _, loan := range r.library.Loans {
        loans = append(loans, loan)
    }
    return loans
}
