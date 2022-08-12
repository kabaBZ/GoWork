package main

import "fmt"

func main() {
	s := []int{1, 2, 3} //len = 3, cap = 3, [1,2,3]

	// //[0, 2)
	// s1 := s[0:2] // [1, 2]

	// fmt.Println(s1)

	// s1[0] = 100

	// fmt.Println(s)
	// fmt.Println(s1)

	//copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) //s2 = [0,0,0]

	// //将s中的值 依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2)

	/* 创建切片 */
	// numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// fmt.Println(numbers)

	// /* 打印原始切片 */
	// fmt.Println("number ==", numbers)

	// /* 打印子切片从索引1(包含)到索引4(不包含) */
	// fmt.Println("numbers[1:4] === ", numbers[1:4])

	// /* 默认下限为 0 */
	// fmt.Println("numbers[:3] ==", numbers[:3])

	// /* 默认上限为 len(s) */
	// fmt.Println("numbers[4:] ==", numbers[4:])

	// numbers1 := make([]int, 0, 5)
	// fmt.Println(numbers1)

	// /* 打印子切片从索引 0(包含) 到索引 2(不包含) */
	// numbers2 := numbers[:2]
	// fmt.Println(numbers2)

	// /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	// numbers3 := numbers[2:5]
	// fmt.Println(numbers3)
}
