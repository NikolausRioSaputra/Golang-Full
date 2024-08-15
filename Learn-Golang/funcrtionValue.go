package main

import "fmt"

func getData(nama string) string {
	test := "selamat bergabung" + nama
	return test
}

func main(){
	contoh := getData
	mana := getData

	fmt.Println(contoh(" niko"))
	fmt.Println(mana(" yian"))
}
