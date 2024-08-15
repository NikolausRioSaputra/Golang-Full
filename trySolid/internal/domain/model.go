package domain

// User represents a library user.
type User struct {
    ID    int
    Name  string
    Email string
}

// Book represents a book in the library.
type Book struct {
    ID     int
    Title  string
    Author string
}

// Loan represents a book loan.
type Loan struct {
    ID     int
    UserID int
    BookID int
}

// Library contains embedded structs for users, books, and loans.
type Library struct {
    Users map[int]*User
    Books map[int]*Book
    Loans map[int]*Loan
}

func NewLibrary() *Library {
    return &Library{
        Users: make(map[int]*User),
        Books: make(map[int]*Book),
        Loans: make(map[int]*Loan),
    }
}
