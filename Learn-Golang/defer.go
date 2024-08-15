package main

import "fmt"

func jalan() {
	fmt.Println("seteleh selesai jalan ")
}

func runAplication(){
	defer jalan()
	fmt.Println("di jalankan ")
}

func main(){
	runAplication()
}