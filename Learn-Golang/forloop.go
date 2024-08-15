package main
 
import "fmt"

func main(){


	// for i := 0; i < 10; i++ {
	// 	fmt.Println("print kan saya dengan jumala ", i)
	// }

	// names := []string{"niko", "laus", "saputra"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Print(names[i])
	// }


	// for index, name := range names {
	// 	fmt.Println("index", index, "=", name[1])
	// }

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("waw" , i)
		
	}

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("lajut ke " , i)
	}
}


