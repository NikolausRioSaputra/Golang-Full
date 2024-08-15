package main

import "fmt"

type Consumer struct {
	Nama, alamat string
	umur         int
}

func (consumer Consumer) sayHello(name string){
	fmt.Println("hallo ", name,  " saya adalah " , consumer.Nama, "dan saat ini saya tingal di", consumer.alamat, "saat ini saya berumur,", consumer.umur)
}

func main() {

	//cara ke 1
	var eko Consumer
	eko.Nama = "nikolaus"
	eko.umur = 22
	eko.alamat = "jl. grogol"

	fmt.Println(eko)
	fmt.Println(eko.alamat)
	fmt.Println(eko.umur)
	fmt.Println(eko.Nama)

	//cara ke 2
	joko := Consumer {
		alamat: "jalan kemangi",
		umur: 28,
		Nama: "nanai",
	}

	fmt.Println(joko)

	// cara ke 3 
	budi := Consumer {"budi", "jakarta pusat", 24}
	fmt.Println(budi)

	niko := Consumer{"Niko", "jakarta", 23}
	fmt.Print(niko)


	budi.sayHello("joko")
	niko.sayHello("joko")
}