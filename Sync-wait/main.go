package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // tells counter to increment by 1 when the function is done
	fmt.Printf("worker %d starting\n", i)
	fmt.Printf("worker %d ending\n", i)
}

func main() {
	fmt.Println("heyhey")

	var wg sync.WaitGroup // check list or task list

	for i := 1; i <= 5; i++ {
		wg.Add(1) // increment the WaitGroup counter
		go worker(i, &wg)
	}
	wg.Wait() // wait for all goroutines to finish
}
