package main

import "fmt"

func endApp() {
	fmt.Println("udah berakhir")
	pesan := recover()
	fmt.Println("terjadi error ini  " , pesan) 
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("ups error")
	}
}


func main(){
	runApp(true)
	fmt.Println("hallo niko ganteng")
}

