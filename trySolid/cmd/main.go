package main

import (
	"fmt"
	"try-solid/internal/domain"
	"try-solid/internal/handler"
	"try-solid/internal/repository"
	"try-solid/internal/usecase"
)

func main() {
    library := domain.NewLibrary()
    userRepo := repository.NewInMemoryRepository(library)
    bookRepo := repository.NewInMemoryRepository(library)
    loanRepo := repository.NewInMemoryRepository(library)
    libraryUC := usecase.NewLibraryInteractor(userRepo, bookRepo, loanRepo)
    libraryHandler := handler.NewLibraryHandler(libraryUC)

    // Menu pilihan
    for {
        fmt.Println("Pilih opsi:")
        fmt.Println("1. Masukkan Buku")
        fmt.Println("2. Pinjam Buku")
        fmt.Println("3. Pengembalian Buku")
        fmt.Println("4. List Pinjaman")
        fmt.Println("5. Keluar")
        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            var bookID int
            var title, author string
            fmt.Println("Masukkan ID Buku:")
            fmt.Scan(&bookID)
            fmt.Println("Masukkan Judul Buku:")
            fmt.Scan(&title)
            fmt.Println("Masukkan Penulis Buku:")
            fmt.Scan(&author)
            book := &domain.Book{ID: bookID, Title: title, Author: author}
            if err := libraryHandler.RegisterBook(book); err != nil {
                fmt.Println("Error registering book:", err)
            } else {
                fmt.Println("Buku berhasil didaftarkan")
            }
        case 2:
            var userID, bookID int
            fmt.Println("Masukkan ID User:")
            fmt.Scan(&userID)
            fmt.Println("Masukkan ID Buku:")
            fmt.Scan(&bookID)
            if err := libraryHandler.BorrowBook(userID, bookID); err != nil {
                fmt.Println("Error borrowing book:", err)
            } else {
                fmt.Println("Buku berhasil dipinjam")
            }
        case 3:
            var loanID int
            fmt.Println("Masukkan ID Peminjaman:")
            fmt.Scan(&loanID)
            if err := libraryHandler.ReturnBook(loanID); err != nil {
                fmt.Println("Error returning book:", err)
            } else {
                fmt.Println("Buku berhasil dikembalikan")
            }
        case 4:
            libraryHandler.ListLoans()
        case 5:
            fmt.Println("Keluar")
            return
        default:
            fmt.Println("Pilihan tidak valid")
        }
    }
}
