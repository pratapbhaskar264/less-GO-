package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
func postRequest() {
	todo1 := Todos{
		Id:          1,
		Todo:        "Grind Leetcode",
		IsCompleted: false,
		UserId:      45,
	}
	jsonData, err := json.Marshal(todo1)

	if err != nil {
		fmt.Println("Error in marshelling the Data ", err)
		return
	}

	jsonStrData := string(jsonData) // string json data

	jsonReaderData := strings.NewReader(jsonStrData) //reader

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	data, err := http.Post(myUrl, "application/json", jsonReaderData)

	if err != nil {
		fmt.Println("error posting data")
		return
	}

	defer data.Body.Close()

	deMarData, _ := ioutil.ReadAll(data.Body)

	fmt.Println(string(deMarData))
}
func updateRequest() {
	todo1 := Todos{
		Id:          1,
		Todo:        "Grind Leetcode and watch Rohit ",
		IsCompleted: false,
		UserId:      45,
	}

	jsonData, err := json.Marshal(todo1)

	if err != nil {
		fmt.Println("error in marshelling data ", err)
	}

	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	jsonStr := string(jsonData)

	jsonReader := strings.NewReader(jsonStr)

	//create a put request
	data, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)

	if err != nil {
		fmt.Println("error in put req ", err)
		return
	}

	data.Header.Set("Content-Type", "application/json") //setting the header

	//sending the request
	client := http.Client{}
	res, err := client.Do(data)

	if err != nil {
		fmt.Println("error in sending the request ", err)
		return
	}

	defer res.Body.Close()

	//reading the response
	respData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(respData))
}
func main() {
	fmt.Println("Hello, World!")
	// getRequest()
	// postRequest()
	updateRequest()
}
