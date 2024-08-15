package domain

type Address struct {
	City string
}

type Person struct {
	ID      int
	Name    string
	Address Address //menempelkan sebuah struct addres sebagai properti struct Person.
}