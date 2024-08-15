package main

import (
	"calculator-app/model"
	"calculator-app/repository"
	"fmt"
)

func main() {
	calculator := repository.Calculator{}
	var exit bool

	for !exit {
		var value1, value2 float32
		var choice int

		fmt.Println("Input nilai pertama:")
		fmt.Scanln(&value1)
		fmt.Println("Input nilai kedua:")
		fmt.Scanln(&value2)

		op := model.Operation{
			Value1: value1,
			Value2: value2,
		}

		fmt.Println("Pilih Operator:")
		fmt.Println("1. Tambah")
		fmt.Println("2. Kurangi")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagaian")
		fmt.Println("5. Keluar")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			result := calculator.Tambah(op)
			fmt.Printf("Result: %.2f\n", result)
		case 2:
			result := calculator.Kurangi(op)
			fmt.Printf("Result: %.2f\n", result)
		case 3:
			result := calculator.Perkalian(op)
			fmt.Printf("Result: %.2f\n", result)
		case 4:
			result, err := calculator.Pembagaian(op)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %.2f\n", result)
			}
		case 5:
			exit = true
			fmt.Println("dasar bisanya phpin doang, terimakasih...")
		default:
			fmt.Println("Pilihan anda kurang tepat")
		}

		if !exit {
			fmt.Println("apakah kamu masi mau mengunakan kalkulator saya ? (yes/no)")
			var response string
			fmt.Scanln(&response)
			if response != "yes" {
				exit = true
				fmt.Println("Terimakasih ...")
			}
		}
	}
}
