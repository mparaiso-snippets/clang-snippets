/*
	Go channels
	1484210530.pdf
	page 55
*/
package main

import "sync"

var wg sync.WaitGroup

func main() {
	count := make(chan int)
	wg.Add(2)
	println("Start Goroutines")
	go printCounts("A", count)
	go printCounts("B", count)
	println("Channel Begin")
	count <- 1
	println("Waiting to finish")
	wg.Wait()
	println("Terminating program")
}

func printCounts(label string, count chan int) {
	defer wg.Done()
	for {
		val, ok := <-count
		if !ok {
			println("Channel was closed")
			return
		}
		println("Count: ", val, " received from ", label)
		if val == 10 {
			println("Channel closed from ", label)
			close(count)
			return
		}
		val++
		count <- val
	}
}
