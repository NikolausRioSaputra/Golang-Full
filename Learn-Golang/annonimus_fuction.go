package main

import "fmt"

type Blockuser func (string) bool;

func registerUser (name string , blockUser Blockuser ){
	if blockUser(name) {
		fmt.Println("your are blocked", name)
	}else{
		fmt.Println("welcome " , name)
	}
}

func main(){

	guaBlock := func (name string) bool{
		return name == "anjing"
	}

	registerUser("niko", guaBlock)

	// contoh 2 

	registerUser("anjing", func (name string ) bool{
		return name == "anjing";
	})
}