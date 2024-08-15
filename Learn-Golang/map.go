package main

import "fmt"

func main(){

	person:= map[string]string{
		"nama" : "nikolausriosaputra",
		"kelas": "phincon group",
	}

	fmt.Println(person["nama"])
	fmt.Println(person["kelas"])
	fmt.Println(person)


	book:= make(map[string]string)

	book["nama"] = "niko"
	book["author"] = "nk"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")
	
	fmt.Println(book)

}