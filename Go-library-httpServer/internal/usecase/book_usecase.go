package usecase

import (
	"GoLibraryHttpserver/internal/domain"
	"GoLibraryHttpserver/internal/repository"
)

type BookUsecaseInterface interface {
	AddBook
	ListBooks
	DeleteBook
}

type AddBook interface {
	AddBook(book domain.Book) error
}

type ListBooks interface {
	ListBooks() ([]domain.Book, error)
}

type DeleteBook interface {
	DeleteBook(bookID int) error
}

type BookUsecaseImpl struct {
	bookRepo repository.BookRepositoryInterface
}

func NewBookUsecase(bookRepo repository.BookRepositoryInterface) BookUsecaseInterface {
	return BookUsecaseImpl{
		bookRepo: bookRepo,
	}
}

func (uc BookUsecaseImpl) AddBook(book domain.Book) error {
	err := uc.bookRepo.SaveBook(&book)
	if err != nil {
		return err
	}

	return nil
}

func (uc BookUsecaseImpl) ListBooks() ([]domain.Book, error) {
	value ,err := uc.bookRepo.ListBooks()
	if err != nil {
		return nil ,err
	}

	return value, nil

}

func (uc BookUsecaseImpl) DeleteBook(bookID int) error {
	err := uc.bookRepo.DeleteBook(bookID)
	if err !=  nil {
		return err
	}

	return nil
}
