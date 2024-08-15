package main

import (
	"fmt"
	"os"
	// "strings"
	// "unicode/utf8"
)

func main() {
	// task 1
	// color := "green"
	// color = "blue"

	// fmt.Println(color)

	// task 2
	// var color string
	// fmt.Print("masukin nama warna nya: ")
	// fmt.Scanln(&color)

	// hasil := "dark" + " " + color
	// fmt.Println(hasil)

	// task 3
	// var value1, value2 , hasil float32

	// fmt.Print("masukan nilai pertama ")
	// fmt.Scanln(&value1)
	// fmt.Print("masukan nilai kedua ")
	// fmt.Scanln(&value2)

	// hasil = value1 * value2
	// fmt.Println(hasil)

	// var (
	// 	perimiter, width, height, pilih int
	// )

	// for {
	// 	fmt.Println("menu")
	// 	fmt.Println("1. area")
	// 	fmt.Println("2. rectangle")
	// 	fmt.Println("3. exit")

	// 	fmt.Println("masukan lebar nya : ")
	// 	fmt.Scanln(&width)

	// 	fmt.Println("masukan panjang nya : ")
	// 	fmt.Scanln(&height)

	// 	fmt.Print("masukan pilihan: ")
	// 	fmt.Scanln(&pilih)

	// 	if pilih == 1 {
	// 		perimiter = width * height
	// 	} else if pilih == 2 {
	// 		perimiter = 2 * (width + height)
	// 	} else if pilih == 3 {
	// 		fmt.Println("terimakasih")
	// 		return
	// 	} else {
	// 		fmt.Println("maaf inputan gagal")
	// 	}

	// 	fmt.Printf("ini hasil Rectangle's Perimeter %v\n", perimiter)

	// }

	// sekaang string

	// raw string literal

	// name:= os.Args[1]
	// name2:= os.Args[2]

	// hi:= `hi ` + name + `how are you ` + name2
	// fmt.Println(hi)
	// we := os.Args[1]
	// convert := utf8.RuneCountInString(we)
	// fmt.Println(convert)

	// msg:= os.Args[1]
	// convert := utf8.RuneCountInString(msg)

	// s:= msg + strings.Repeat("!", convert)
	// fmt.Println(s)

	// msg:= os.Args[1]

	// fmt.Println(strings.ToLower(msg))

	// 	msg := `

	// 	The weather looks good.
	// I should go and play.

	// 	`

	// 	fmt.Println(strings.TrimSpace(msg))

	// name :="inanç           "
	// hapusKanan := strings.TrimRight(name, " ")
	// convert := utf8.RuneCountInString(hapusKanan)

	// fmt.Println(name)
	// fmt.Print(convert)

	// const (
	// 	Nov  = iota + 9
	// 	Oct
	// 	Sep
	// )

	// fmt.Println(Nov, Oct , Sep)

	// const (
	// 	_ = iota
	// 	Jan
	// 	Feb
	// 	Mar
	// )
	// fmt.Println(Jan, Feb, Mar)

	// const (
	// 	Spring = (iota + 1) * 3
	// 	Summer
	// 	Fall
	// 	Winter
	// )

	// fmt.Println(Winter, Spring, Summer, Fall)

	// age := 30
	// fmt.Printf("I'm %d years old.", age)

	// fmt.Printf("Type of %s is %[1]T\n", "hello")

	// var age int
	// fmt.Println("masukin nama age :")
	// fmt.Scanln(&age)

	// if age >= 60 {
	// 	fmt.Println("getting older")
	// } else if age >= 30 {
	// 	fmt.Println("getting wiser")
	// }else if age >= 20 {
	// 	fmt.Println("Adulthood")
	// }else if age >= 10 {
	// 	fmt.Println("Young blood")
	// }else{
	// 	fmt.Println("Booting up")
	// }

	// isSphere, radius := true, 200

	// var big bool

	// if !isSphere == false) {
	// 	fmt.Println("It's a big sphere.")
	// } else {
	// 	fmt.Println("I don't know.")
	// }

	//  go run main.go

//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments

	msg:= os.Args
	l := len(msg) - 1

	if l == 0 {
		fmt.Println("Give me args")
	}else if l == 1{
		fmt.Printf("There is one: %q", msg[1])
	}else if l == 2{
		fmt.Printf("There is two: %s %s", msg[1], msg[2])
	}else if l == 3 {
		fmt.Printf("There is there:%s %s %s ", msg[1], msg[2], msg[3])
	}else{
		fmt.Printf("There are %d arguments", l)
	}

}
