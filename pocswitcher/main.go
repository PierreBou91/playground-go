package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 16)
	var step int

	go func() {
		for {
			fmt.Scan(&step)
			ch <- step
		}
	}()

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
