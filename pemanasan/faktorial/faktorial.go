package main

import "fmt"

func factorial(number uint64) uint64 {
	if number == 1 {
		return 1
	}

	factorialOfNumber := number * factorial(number-1)
	return factorialOfNumber
}
func main() {
	var number uint64
	fmt.Println("masukan no yang mau kamu: ")
	fmt.Scanln(&number)

	fmt.Println("faktorial", number, "adalah", factorial(number))
}
