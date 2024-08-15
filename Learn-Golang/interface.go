package main

import "fmt"

type HasName interface {
	GetNama()	string
}

func SayHello(value HasName) {
	fmt.Println("hello" , value.GetNama())
}

type Person struct {
	Nama string
}

func(person Person) GetNama() string {
	return person.Nama
}

func main() {
	person := Person{Nama: "Niko"}
	SayHello(person)
}