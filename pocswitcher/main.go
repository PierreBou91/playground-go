package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	go func() {
		for {
			a := <-ch
			switch a {
			case 1:
				fmt.Println("2")
				ch <- 2
			case 2:
				fmt.Println("3")
				ch <- 3
			case 3:
				fmt.Println("4")
				ch <- 4
			case 4:
				fmt.Println("5")
			}
		}
	}()

	var step int
	for {
		fmt.Scan(&step)
		ch <- step
	}
}
