package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
	Age   int    `json:"age"`
}

// type User struct {
// 	Name  string
// 	Hobby string
// 	Age   int
// }

// func main() {
// 	var typeString = `{"name" : "Nikolaus rio saputra", "hobby":"Badminton", "age": 22 }`
// 	var JsonData = []byte(typeString)

// 	var user User //json ke struct

// 	err := json.Unmarshal(JsonData, &user)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println(user)
// 	fmt.Println(user.Name)
// 	fmt.Println(user.Age)
// 	fmt.Println(user.Hobby)
// }

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := []User{{"nikolaus", "badminton", 12}}
	// user := []User{{"name": "nikolaus", "hoby":"badminton", "age":12}}

	var bytes, err = json.Marshal(user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonObject = string(bytes)
	fmt.Println(jsonObject)

	// http.Handle("/foo", fooHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		user := &Person{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("got user:", user)
		w.WriteHeader(http.StatusCreated)
	})

	if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
		panic(err)
	}
}
