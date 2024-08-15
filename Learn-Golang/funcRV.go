package main

import "fmt"

// func say(name string) string {
// 	hallo := "hallo " + name
// 	return hallo;
// }

// func main(){
// 	result := say("niko")
// 	fmt.Println(result)

// 	fmt.Print(say("good atitude"))
// }

func callGroup() (string, string){
	return "niko" ,"saputra"
}

func main(){
	firstName, lastName := callGroup()
	fmt.Println(firstName, lastName)
}