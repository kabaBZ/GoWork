package main

import "fmt"

// 声明一种新的数据类型 myint， 是int的一个别名
type myint int

type Book struct {
	title string
	price string
}

func changeBook(book Book) {
	// 传递一个book的副本
	book.price = "666"
}

func changeBook2(book *Book) {
	// 指针传递
	book.price = "777"
}

func main() {
	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var book Book
	book.title = "Golang"
	book.price = "111"
	fmt.Printf("%v\n", book)

	changeBook(book)
	fmt.Printf("%v\n", book)

	changeBook2(&book)
	fmt.Printf("%v\n", book)
}
