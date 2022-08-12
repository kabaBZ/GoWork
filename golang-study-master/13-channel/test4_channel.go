package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer close(c)
		for i := 0; i < 5; i++ {
			c <- i
		}
	}()

	// 可以使用range来迭代不断操作channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished..")
}
