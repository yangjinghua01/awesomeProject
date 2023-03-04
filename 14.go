package main

import "fmt"

// 数组
// 存放元素的容器
// 数组的长度是数组类型的一部分
func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Println("a1:%T a2:%T", a1, a2)
	//数组的初始化
	//如果不初始化：默认元素都是零值（布尔值：false，整型和浮点型都是0，字符串：""）
	fmt.Println(a1, a2)
	//1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//	初始化方式2
	a100 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a100)
	a10 := [...]int{0, 1, 7, 2, 3, 4, 5}
	//根据出事值自动推断数组的铲毒是多少
	fmt.Println(a10)
	//3.初始化方式3
	a3 := [5]int{2, 3}
	fmt.Println(a3)
	//指定索引的方式
	a5 := [5]int{0: 1, 4: 5}
	fmt.Println(a5)
	//遍历数组
	citys := [5]string{"北京", "上海", "广州", "深圳", "浙江"}
	for i := 0; i < len(citys); i++ {
		print("数组索引%d是%s", i, citys[i])
	}
}
