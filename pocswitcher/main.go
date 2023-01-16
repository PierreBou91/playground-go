package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	ch4 := make(chan int, 1)

	go func() {
		for {
			select {
			case <-ch1:
				fmt.Println("2")
				ch2 <- 1
			case <-ch2:
				fmt.Println("3")
				ch3 <- 1
			case <-ch3:
				fmt.Println("4")
				ch4 <- 1
			case <-ch4:
				fmt.Println("5")
			}
		}
	}()

	var step int
	for {
		fmt.Scan(&step)
		switch step {
		case 1:
			ch1 <- 1
		case 2:
			ch2 <- 1
		case 3:
			ch3 <- 1
		case 4:
			ch4 <- 1
		}
	}
}
