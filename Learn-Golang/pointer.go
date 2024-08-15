package main

import "fmt"

type Address struct {
	nama, kota, alamat string
}

func main() {
	// address1 := Address{"niko", "jakarta", "jl.kota"}
	// address2 := address1

	// address2.nama = "rio"

	// fmt.Println(address1)
	// fmt.Println(address2)

	var address1 Address = Address{"niko", "jakarta", "jl.kota"}
	var address2 *Address = &address1
	

	address2.nama = "rio"

	fmt.Println(address1)
	fmt.Println(address2)

}