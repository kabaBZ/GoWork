package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			// 如果c可写，则进入该case
			x, y = y, x+y
		case c <- y:
		case <-quit:
			// 如果quit可读，则进入该case
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// sub go
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	// main go
	fibonacii(c, quit)
}
