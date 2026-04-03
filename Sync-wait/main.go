package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d starting\n", i)
	fmt.Printf("worker %d ending\n", i)
}

func main() {
	fmt.Println("heyhey")

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
