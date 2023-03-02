package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)
	//算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独语句，不能放在=的右边赋值 ==》 a = a +1
	b-- //单独的语句，不能放在=的右边赋值 ==》 b = b -1
	fmt.Println(a)
	fmt.Println(b)
	//	关系运算符
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a <= b)
	age := 18
	if age >= 18 && age <= 60 {
		fmt.Println("打工人")
	} else {
		println("食利者")
	}
}
