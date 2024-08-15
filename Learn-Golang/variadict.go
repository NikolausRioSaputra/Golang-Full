package main

import "fmt"


// func sum(numbers ...int) int {
// 	total := 0

// 	for _, number := range numbers{
// 		total += number
// 	}  

// 	return total
// }


// func main(){

// 	fmt.Println(sum(10,10,10,10))
// 	fmt.Println(sum(10,10,10,10,10,10))
// 	fmt.Println(sum(10,10,10,10,10,10,10))
// }


func doubleVal() (age int , firstname , address string){
	age = 22
	firstname = "niko"
	address = "jln.yang jauh"

	return age, firstname, address
}

func main(){
	a, b , c := doubleVal()
	fmt.Printf("umur mu segini ya  %d\n", a)
	fmt.Printf("HAI %s\n",b)
	fmt.Printf("alamat mu disini kan  %s\n", c)

}