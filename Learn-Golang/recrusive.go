package main

import (
	"fmt"

	"golang.org/x/telemetry/counter"
)

func factorialLoop(value int) int {
	reslut := 1

	for i := value; i > 0; i-- {
		reslut *= i

	}

	return reslut
}

func factorialRecrusive(value int ) int {
	if value == 1  {
		return 1
	}else{
		nilai := value * factorialRecrusive(value - 1 )
		return nilai
	}
}


func main() {

	
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecrusive(10))
}

