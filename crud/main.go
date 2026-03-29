package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todos struct {
	Id          int    `json:"id"`
	Todo        string `json:"todo"`
	IsCompleted bool   `json:"completed"`
	UserId      int    `json:"userId"`
}

func getRequest() {
	data, err := http.Get("https://dummyjson.com/todos/45")

	if err != nil {
		fmt.Println("error finding todosss")
		return
	}
	defer data.Body.Close()

	if data.StatusCode != http.StatusOK {
		fmt.Println("Not 200", data.StatusCode)
		return
	}

	// orData, err := io.ReadAll(data.Body)

	// if err != nil {
	// 	fmt.Println("error finding todosss")
	// 	return
	// }

	// fmt.Println(string(orData))

	var todo Todos

	err1 := json.NewDecoder(data.Body).Decode(&todo)

	if err1 != nil {
		fmt.Println("error : ", err1)
		return
	}

	fmt.Println(todo)
}

func main() {
	fmt.Println("Hello, World!")
	getRequest()
	todo1 := Todos{
		Id:          1,
		Todo:        "Leetcode",
		IsCompleted: false,
		UserId:      45,
	}

	fmt.Println(todo1)

	//    jsonData , err := json.Marshal(todo1)

	//    if err != nil {
	// 	fmt.Println("Error marshelling : " , err)
	// 	return
	//    }

	//    const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

}
