package main

import (
	"fmt"
)

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key = ", key+", value = ", value)
	}
}

func AddValue(cityMap map[string]string) {
	// map 是引用传递
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)
	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"
	// 删除
	delete(cityMap, "China")
	// 遍历
	printMap(cityMap)
	fmt.Println("-------")

	// 修改
	cityMap["USA"] = "DC"
	// 添加
	AddValue(cityMap)
	// 遍历
	printMap(cityMap)
}
