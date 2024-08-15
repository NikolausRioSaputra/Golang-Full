package main

import (
	"day-one-golang/repository"
	"fmt"
	"strings"
)

func main() {
	var id int
	var namaProduct string
	var action string

	fmt.Println("Masukan nama id: ")
	fmt.Scanln(&id)

	fmt.Println("Masukan nama product")
	fmt.Scanln(&namaProduct)

	productRepo := repository.NewProductRepository(id, namaProduct)

	for{
		fmt.Println("\nPilih tidakan (tambah/kurang/tampilkan/keluar)")
		fmt.Scanln(&action)

		action = strings.ToLower(action)

		switch action {
		case "tambah":
			var jumlahTambah int
			fmt.Println("Masukan jumlah yang ingin ditambah: ")
			fmt.Scanln(&jumlahTambah)
			productRepo.Tambah(jumlahTambah)

		case "kurang":
			var jumlahKurang int 
			fmt.Println("Masukan jumlah yang ingin dikurang: ")
			fmt.Scanln(&jumlahKurang)
			err := productRepo.Kurang(jumlahKurang)
			if err != nil {
				fmt.Println(err.Error())
			}
		
		case "tampilkan":
			productRepo.Tampilkan()
		
		case "keluar":
			fmt.Println("Terimakasih")
			return

		default:
			fmt.Println("Tindakan tidak dikenalai, silahkan coba lagi")
		}
	}
}