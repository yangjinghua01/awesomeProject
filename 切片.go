package main

import "fmt"

// 切片
func main() {
	//切片定义
	var s []int
	var s2 []string
	fmt.Println(s, s2)
	fmt.Println(s == nil)
	fmt.Println(s2 == nil)
	s = []int{1, 2, 3, 45, 5, 67, 7, 8, 8}
	s2 = []string{"k", "i", "l", "l"}
	//初始化
	fmt.Println(s, s2)
	fmt.Println(s == nil)
	fmt.Println(s2 == nil)
	//长度
	fmt.Printf("len(s1):%d cap(s1)%d\n", len(s), cap(s))
	fmt.Printf("len(s2):%d cap(s2)%d\n", len(s2), cap(s2))
	//由数组得到切片
	a1 := [...]int{1, 2, 3, 45, 55}
	s3 := a1[0:4] //基于一个数组切割，左包含右不包含
	fmt.Println(s3)
	s4 := a1[:len(a1)]
	fmt.Println(s4)
	s5 := a1[:]
	fmt.Println(s5)
}
