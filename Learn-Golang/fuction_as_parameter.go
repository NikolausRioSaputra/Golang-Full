package main

import "fmt"

func sayHelloWithFilter(name string , filter func (string) string ) {
	sayHello := filter(name)
	fmt.Println("hallo " , sayHello)
}

func spamWithFilter(name string) string {
	if name == "anjing"{
		return "/./."
	}else{
		return name
	}
}

func main() {
	filter := spamWithFilter
	sayHelloWithFilter("niko", filter)
	sayHelloWithFilter("anjing", filter)
}