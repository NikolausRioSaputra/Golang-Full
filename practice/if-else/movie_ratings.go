package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Requires Ages")
		return
	}

	age, err := strconv.Atoi(os.Args[1])

	if err != nil || args < 0 {
		fmt.Printf(`wrong age: %q` + "\n", os.Args[1])
	}else if age > 17 {
		fmt.Println("")
	}else if age >