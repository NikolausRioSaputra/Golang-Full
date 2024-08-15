package main

import "fmt"

func main() {
	var price, quantity, discount, total, hasil int
	fmt.Println("masukan harga :")
	fmt.Scanln(&price)
	fmt.Println("masukan quantity:")
	fmt.Scanln(&quantity)
	fmt.Println("masukan discount :")
	fmt.Scanln(&discount)

	total = price * quantity
	if discount > 0 {
		hasil = total - (total * discount / 100)
	}

	fmt.Printf("total harga barang %d\n", hasil)
}
