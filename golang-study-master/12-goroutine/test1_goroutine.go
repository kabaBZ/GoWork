package main

import (
	"fmt"
	"time"
)

// 子routine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutie: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主routine
func main() {
	// 创建一个子进程 去执行newTask()流程
	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
