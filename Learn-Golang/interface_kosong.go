package main

import "fmt"

func Ups() any {
	return "niko"
}

func main() {
	var kosong = Ups()
	fmt.Println(kosong)
}