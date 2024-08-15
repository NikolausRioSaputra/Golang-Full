package main

import "fmt"

func main() {
	var score int
	fmt.Println("masukan nilai :")
	fmt.Scanln(&score)

	var result string
	if score >= 90 {
		result = "Selamat anda mendapatkan nilai A"
	}else if score >= 80 && score < 90 {
		result = "Anda mendaptkan nilai B"
	}else if score >= 70 && score < 80 {
		result = "Anda mendapatkan nilai C"
	}else if score >= 60 && score < 70 {
		result = "Anda mendapatkan nilai D"
	}else {
		result = "Anda mendapatkan nilai E"
	}

	fmt.Println(result)
}