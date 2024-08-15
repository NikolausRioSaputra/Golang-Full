package repository

import (
	"GoLibraryHttpserver/internal/domain"
	"errors"
)

type BookRepositoryInterface interface {
	BookSaver
	BookLister
	BookDeleter
}

// ini interface implementasi di repo untuk nyimpan buku
type BookSaver interface {
	SaveBook(book *domain.Book) error
}

type BookLister interface{
	ListBooks() ([]domain.Book, error)
}

type BookDeleter interface {
	DeleteBook(bookID int) error
}

type BookRepository struct {
	books map[int]domain.Book
}

func NewBookRepository() BookRepositoryInterface {
	return &BookRepository{
		books: map[int]domain.Book{},
	}
}

func (repo *BookRepository) SaveBook(book *domain.Book) error {
	if _, exists := repo.books[book.ID]; exists {
		return errors.New("book already exists")
	}

	repo.books[book.ID] = *book
	return nil
}

func (repo *BookRepository) ListBooks() ([]domain.Book, error) {
	books := []domain.Book{}
	for _,book := range repo.books {
		books = append(books, book)
	}

	return books, nil
}

func (repo *BookRepository) DeleteBook(bookID int) error {
	if _, exists := repo.books[bookID]; !exists {
		return errors.New("books not found")
	}

	delete(repo.books, bookID)
	return nil
}