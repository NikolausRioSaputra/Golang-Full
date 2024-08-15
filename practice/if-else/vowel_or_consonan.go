package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("give me a args")
		return
	}

	// s := args[1]
	// if s == "a" || s == "i" || s == "u" || s == "e" || s == "o" {
	// 	fmt.Printf("%q ini merupakan vowel\n", s)
	// } else if s == "y" || s == "w" {
	// 	fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	// } else {
	// 	fmt.Printf("%q is a consonant.\n", s)
	// 	return
	// }

	s := args[1]
	if strings.IndexAny(s, "aiueo") != -1 {
		fmt.Printf("%q what is this", s)
	}else if strings.IndexAny(s, "yw") != -1{
		fmt.Printf("%q this is", s)
	}else {
		fmt.Printf("%q is what bro i dont understand", s)
	}
}
