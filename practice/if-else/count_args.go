package main

import (
	"fmt"
	"os"
)

func main() {
	msg := os.Args
	l := len(msg) - 1

	if l == 0 {
		fmt.Printf("give me args")
	}else if l == 1 {
		fmt.Printf("There is one %q\n ", msg[1])
	}else if l == 2 {
		fmt.Printf("There is two: %s %s ", msg[1], msg[2])
	}else if l == 3 {
		fmt.Printf("There is three: %s\n", msg[1], msg[2], msg[3])
	}else {
		fmt.Printf("there are %d arguments", l)
	}
}