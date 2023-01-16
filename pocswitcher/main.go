package main

import (
	"fmt"
	"time"
)

func main() {
	// ch is abuffered channel to receive orders, depending on the time it takes to process the
	// order and the complexity (CPU), the buffer size can be adapted with this rule of thumbs:
	// "How many maximum processes should we expect to run at the same time without crashing
	// the system?"
	// 16 should be plenty and not too much for the CPU since it's mainly API calls
	ch := make(chan int, 16)
	var step int

	// goroutine to send orders either from the watcher or from the main process
	go func() {
		for {
			fmt.Scan(&step)
			ch <- step
		}
	}()

	// loop to receive orders and process them according to the step.
	// each time a gets a value, it will spawn a goroutine to process the order, this means that
	// every step will be launched in its own goroutine == MUCHO EFFICIENT
	for a := <-ch; ; a = <-ch {
		go func(a int) {
			switch a {
			case 1:
				fmt.Println("2")
				ch <- 2
			case 2:
				fmt.Println("3")
				time.Sleep(1 * time.Second) // to confirm multiple goroutines
				ch <- 3
			case 3:
				fmt.Println("4")
				ch <- 4
			case 4:
				fmt.Println("5")
			}
		}(a)
	}
}
