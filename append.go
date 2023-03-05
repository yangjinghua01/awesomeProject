package main

import "fmt"

// append()为切片扩容
func main() {
	s1 := []string{"beijing", "shanghai", "guangzhou", "shenzhen"}
	s1[3] = "广州"
	fmt.Println(s1)
	s1 = append(s1, "浙江")
	fmt.Println(s1)
	/**
	调用append函数必须用原来的切片变量接受返回值
	append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
	必须用变量接受append的返回值
	*/
	s1 = append(s1, "长沙")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...) //...是拆开的意思
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
