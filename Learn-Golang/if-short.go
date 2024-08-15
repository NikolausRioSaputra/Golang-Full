package main

import (
	"fmt"
)

func main() {

	nama := "nikolaus";

	// if length := len(nama); length > 5  {
	// 	fmt.Println("perkenalakan nama saya ?")
	// }else{
	// 	fmt.Println("maaf saya masi permula")
	// }

	switch nama{
		case "eko":
			fmt.Println("niko")
		case "nikolaus":
			fmt.Println("nah benar")
		default:
			fmt.Println("maaf kamu error") 
	}	

	// switch length:= len(nama); length > 5 {
	// case true:
	// 	fmt.Println("benar")
	// case false:
	// 	fmt.Println("salah")
	// }

	length := len(nama)
	switch {
	case length > 10:
		fmt.Println("nama kamu terlalu panjang")
	case length > 5:
		fmt.Println("nama kmu udah pas si")
	default:
		fmt.Println("nama kmu sudah pas")
	}
}