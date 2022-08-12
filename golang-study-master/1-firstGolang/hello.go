package main

// 导多个包
import (
	"fmt"
	"time"
)

// Golang 有无 ; 都可，建议不加
func main() { // 函数的 { 必须和函数名同行, 否则编译报错
	fmt.Println("Hello Go!")
	time.Sleep(1 * time.Second)
}
