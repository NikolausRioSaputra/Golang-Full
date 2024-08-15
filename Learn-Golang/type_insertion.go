package main

import "fmt"

// func random() interface{} {
// 	return "ok"
// }

func random() any {
	return true
}

func main() {
	var data any = random()
	// fmt.Println(data)
	// fmt.Println(data.(string))

	// joko := random()
	// fmt.Println(joko.(int))

	switch value := data.(type) {
	case int:
		fmt.Println("Integer",value)
	case string:
		fmt.Println("string", value)
	default:
		fmt.Println("Uknown", value)
	}

}
