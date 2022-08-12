package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
}

func (b *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (b *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	// b: pair<type: Book, value: book{} 地址>
	b := &Book{}

	// book ---> reader
	// r: pair<type: , value: >
	var r Reader
	// r: pair<type: Book, value: book{} 地址>
	r = b
	r.ReadBook()

	// reader ---> writer
	// w: pair<type: , value: >
	var w Writer
	// w: pair<type: Book, value: book{} 地址>
	w = r.(Writer) // 此处的断言为什么成功？因为 w, r 的type是一致的
	w.WriteBook()
}
