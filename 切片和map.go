package main

import "fmt"

func main() {
	//元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10)
	//没有对内部的map做初始化
	s1[0] = make(map[int]string)
	s1[0][10] = "sv"
	fmt.Println(s1)
	//值为切片类型的map
	var m1 = make(map[string][]int, 10)
	fmt.Println(m1)
}
