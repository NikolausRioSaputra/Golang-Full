package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("masukan bilang: ")
		fmt.Scan(&i)
		if i <= 10{
			if i%2==0 {
				fmt.Println(i, "ini bilangan genap")
			}else{
				fmt.Println(i, "ini bilangan ganjil")
			}
		}else {
			fmt.Println("pertanyaan selesai , angaka sudah melebihi 10 , terimakasih")
		}

	}
}