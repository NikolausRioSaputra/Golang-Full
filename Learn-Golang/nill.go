package main

import "fmt"

func say(name string) map[string]string {
	if name == "" {
		return nil
	}else{
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := say("")

	if data == nil {
		fmt.Println("data ini masih kosong")
	}else{
		fmt.Println(data["name"])
	}
}