package main

import "fmt"

//如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	//如果说类的属性首字母大写, 表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name  string
	Ad    int
	level int
}

/*
func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string) {
	//this 是调用该方法的对象的一个副本（拷贝）
	this.Name = newName
}
*/
func (h *Hero) Show() {
	fmt.Println("Name = ", h.Name)
	fmt.Println("Ad = ", h.Ad)
	fmt.Println("Level = ", h.level)
	fmt.Println("---------")
}

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) SetName(newName string) {
	h.Name = newName
}

func main() {
	// 创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100}
	hero.Show()

	hero.SetName("li4")
	hero.Show()
}
