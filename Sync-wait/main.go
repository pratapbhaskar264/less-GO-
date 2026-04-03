package main

import (
	"fmt"
	"sync"
)

func worker(i int) {
	defer wq.Done()
	fmt.Printf("worker %d starting\n", i)
	fmt.Printf("worker %d ending\n", i)
}

func main() {
	fmt.Println("heyhey")

	var wq sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wq.Add(1)
		go worker(i)
	}
}
