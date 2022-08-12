package main

import (
	"fmt"
)

//定义接口
type Phone interface {
	call()
	call2()
}

//一直都搞不懂这是干啥的
//原来是用来定义结构体内的数据类型的

type Phone1 struct {
	id            int
	name          string
	category_id   int
	category_name string
}

//第一个类的第一个回调函数
func (test Phone1) call() {
	fmt.Println("这是第一个类的第一个接口回调函数 结构体数据：", Phone1{id: 1, name: "浅笑"})
}

//第一个类的第二个回调函数
func (test Phone1) call2() {
	fmt.Println("这是一个类的第二个接口回调函数call2", Phone1{id: 1, name: "浅笑", category_id: 4, category_name: "分类名称"})
}

//第二个结构体的数据类型
type Phone2 struct {
	member_id       int
	member_balance  float32
	member_sex      bool
	member_nickname string
}

//第二个类的第一个回调函数
func (test2 Phone2) call() {
	fmt.Println("这是第二个类的第一个接口回调函数call", Phone2{member_id: 22, member_balance: 15.23, member_sex: false, member_nickname: "浅笑18"})
}

//第二个类的第二个回调函数
func (test2 Phone2) call2() {
	fmt.Println("这是第二个类的第二个接口回调函数call2", Phone2{member_id: 44, member_balance: 100, member_sex: true, member_nickname: "陈超"})
}

//开始运行
func main() {
	var phone Phone

	//先实例化第一个接口
	phone = new(Phone1)
	phone.call()
	phone.call2()

	//实例化第二个接口
	phone = new(Phone2)
	phone.call()
	phone.call2()
}
