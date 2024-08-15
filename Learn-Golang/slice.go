package main

import (
	"fmt"
)

func main(){

	names := [...]string { "niko", "laus", "rio" ,"saputra", "yohanes", "apheng"}


	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:4]
	fmt.Println(slice2)


	slice3 := names[4:]


	nameSlice := append(slice3, "batuampar")
	nameSlice[1] = "jakarata"
	fmt.Println(nameSlice)
	fmt.Println(slice3)

	slice4 := names[:]

	testCopy := make([]string, len(slice4), cap(slice4))
	copy(testCopy, slice4)

	fmt.Println(testCopy)
	fmt.Println(slice4) 


	sliceTerbaru := make([]string, 2, 5)
	sliceTerbaru[0] = "Niko"
	sliceTerbaru[1] = "Niko"

	fmt.Println(sliceTerbaru)
	fmt.Println(len(sliceTerbaru))
	fmt.Println(cap(sliceTerbaru))

	sliceTerbaru2 := append(sliceTerbaru, "new niko")
	

	fmt.Println(sliceTerbaru2)
	fmt.Println(len(sliceTerbaru2))
	fmt.Println(cap(sliceTerbaru2))
	

	sliceTerbaru2[0] = "try"
	fmt.Println(sliceTerbaru2)
	fmt.Println(sliceTerbaru)

}