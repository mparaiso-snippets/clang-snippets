/*
  Go concureency example
  1484210530.pdf
  page 51
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("Start Goroutines")
	go printCounts("A")
	go printCounts("B")
	wg.Wait()
	fmt.Println("Terminating program")
}

func printCounts(label string) {
	defer wg.Done()
	// randomly wait
	for count := 1; count <= 10; count++ {
		sleep := rand.Int63n(500)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Count %d from %s\n", count, label)
	}
}
