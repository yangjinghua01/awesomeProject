package main

import "fmt"

// 函数的定义
func sum(a int, b int) (ret int) {
	return a + b
}

// 有参数无返回值
func void(x int, y int) {
	fmt.Println("有参数无返回值%d", x+y)
}

// 无参数无返回值
func void1() {
	fmt.Println("无参数无返回值")
}

// 无参数但有返回值
func void2() (ret string) {
	fmt.Println("无参数但有返回值")
	return "hahhaha"
}

// 返回值可以命名也可以不命名
func f(x int, y int) (ret int) {
	ret = x + y
	return ret
}

// 多个返回值
func f5() (int, string) {
	return 1, "xxx"
}

// 参数的类型简写
func f6(x, y int, m, n string) (ret int) {
	return x * y
}

// 可变长参数
func f7(x int, y ...string) {
	fmt.Println(x)
	fmt.Println(y)
}
func main() {
	r := sum(1, 2)
	fmt.Println(r)
	_, s := f5()
	fmt.Println(s)
}
