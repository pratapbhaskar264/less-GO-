package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("JSON FILE")

	person1 := Person{Name: "Bhaskar", Age: 22}

	fmt.Println(person1)
	fmt.Println(person1.Age)

	// person1 to Json encoding (marshalling)

	jsonData, err := json.Marshal(person1)

	if err != nil {
		fmt.Println("Cannot be marshelled")
		return
	}

	fmt.Println(string(jsonData))

	// decoding unmarshelling

	var person11 Person

	err = json.Unmarshal(jsonData, &person11)

	if err != nil {
		fmt.Println("not able to demarshal")
		return
	}

	fmt.Println(person11)

}
